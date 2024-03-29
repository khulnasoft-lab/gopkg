// Package path contains path and environment utilities for Gopkg.
//
// This includes tools to find and manipulate Go path variables, as well as
// tools for copying from one path to another.
package path

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// DefaultGopkgFile is the default name for the gopkg.yaml file.
const DefaultGopkgFile = "gopkg.yaml"

// VendorDir is the name of the directory that holds vendored dependencies.
//
// As of Go 1.5, this is always vendor.
var VendorDir = "vendor"

// Tmp is the temporary directory Gopkg should use. Defaults to "" which
// signals using the system default.
var Tmp = ""

// Cache the location of the homedirectory.
var homeDir = ""

// GopkgFile is the name of the Gopkg file.
//
// Setting this is not concurrency safe. For consistency, it should really
// only be set once, at startup, or not at all.
var GopkgFile = DefaultGopkgFile

// LockFile is the default name for the lock file.
const LockFile = "gopkg.lock"

func init() {

	// As of Go 1.8 the GOPATH is no longer required to be set. Instead there
	// is a default value. If there is no GOPATH check for the default value.
	// Note, checking the GOPATH first to avoid invoking the go toolchain if
	// possible.
	if gopaths = os.Getenv("GOPATH"); len(gopaths) == 0 {
		goExecutable := os.Getenv("GOPKG_GO_EXECUTABLE")
		if len(goExecutable) <= 0 {
			goExecutable = "go"
		}
		out, err := exec.Command(goExecutable, "env", "GOPATH").Output()
		if err == nil {
			gopaths = strings.TrimSpace(string(out))
		}
	}
}

// Home returns the Gopkg home directory ($GOPKG_HOME or ~/.gopkg, typically).
//
// This normalizes to an absolute path, and passes through os.ExpandEnv.
func Home() string {
	if homeDir != "" {
		return homeDir
	}

	if h, err := homedir.Dir(); err == nil {
		homeDir = filepath.Join(h, ".gopkg")
	} else {
		cwd, err := os.Getwd()
		if err == nil {
			homeDir = filepath.Join(cwd, ".gopkg")
		} else {
			homeDir = ".gopkg"
		}
	}

	return homeDir
}

// SetHome sets the home directory for Gopkg.
func SetHome(h string) {
	homeDir = h
}

// Vendor calculates the path to the vendor directory.
//
// Based on working directory, VendorDir and GopkgFile, this attempts to
// guess the location of the vendor directory.
func Vendor() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Find the directory that contains gopkg.yaml
	yamldir, err := GopkgWD(cwd)
	if err != nil {
		return cwd, err
	}

	gopath := filepath.Join(yamldir, VendorDir)

	// Resolve symlinks
	info, err := os.Lstat(gopath)
	if err != nil {
		return gopath, nil
	}
	for i := 0; IsLink(info) && i < 255; i++ {
		p, err := os.Readlink(gopath)
		if err != nil {
			return gopath, nil
		}

		if filepath.IsAbs(p) {
			gopath = p
		} else {
			gopath = filepath.Join(filepath.Dir(gopath), p)
		}

		info, err = os.Lstat(gopath)
		if err != nil {
			return gopath, nil
		}
	}

	return gopath, nil
}

// Gopkg gets the path to the closest gopkg file.
func Gopkg() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Find the directory that contains gopkg.yaml
	yamldir, err := GopkgWD(cwd)
	if err != nil {
		return cwd, err
	}

	gf := filepath.Join(yamldir, GopkgFile)
	return gf, nil
}

// GopkgWD finds the working directory of the gopkg.yaml file, starting at dir.
//
// If the gopkg file is not found in the current directory, it recurses up
// a directory.
func GopkgWD(dir string) (string, error) {
	fullpath := filepath.Join(dir, GopkgFile)

	if _, err := os.Stat(fullpath); err == nil {
		return dir, nil
	}

	base := filepath.Dir(dir)
	if base == dir {
		return "", fmt.Errorf("Cannot resolve parent of %s", base)
	}

	return GopkgWD(base)
}

// Stores the gopaths so they do not get repeatedly looked up. This is especially
// true when the default value needs to be retrieved from `go env GOPATH`.
// TODO(mattfarina): Instead of a singleton an application context would be a
// better place to store things like this.
var gopaths string

// Gopath gets GOPATH from environment and return the most relevant path.
//
// A GOPATH can contain a colon-separated list of paths. This retrieves the
// GOPATH and returns only the FIRST ("most relevant") path.
//
// This should be used carefully. If, for example, you are looking for a package,
// you may be better off using Gopaths.
func Gopath() string {
	gopaths := Gopaths()
	if len(gopaths) == 0 {
		return ""
	}
	return gopaths[0]
}

// Gopaths retrieves the Gopath as a list when there is more than one path
// listed in the Gopath.
func Gopaths() []string {
	p := strings.Trim(gopaths, string(filepath.ListSeparator))
	return filepath.SplitList(p)
}

// Basepath returns the current working directory.
//
// If there is an error getting the working directory, this returns ".", which
// should function in cases where the directory is unlinked... Then again,
// maybe not.
func Basepath() string {
	base, err := os.Getwd()
	if err != nil {
		return "."
	}
	return base
}

// StripBasepath removes the base directory from a passed in path.
func StripBasepath(p string) string {
	bp := Basepath()
	return strings.TrimPrefix(p, bp+string(os.PathSeparator))
}

// IsLink returns true if the given FileInfo references a link.
func IsLink(fi os.FileInfo) bool {
	return fi.Mode()&os.ModeSymlink == os.ModeSymlink
}

// HasLock returns true if this can stat a lockfile at the givin location.
func HasLock(basepath string) bool {
	_, err := os.Stat(filepath.Join(basepath, LockFile))
	return err == nil
}

// IsDirectoryEmpty checks if a directory is empty.
func IsDirectoryEmpty(dir string) (bool, error) {
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdir(1)

	if err == io.EOF {
		return true, nil
	}

	return false, err
}

// CopyDir copies an entire source directory to the dest directory.
//
// This is akin to `cp -a src/* dest/`
//
// We copy the directory here rather than jumping out to a shell so we can
// support multiple operating systems.
func CopyDir(source string, dest string) error {

	// get properties of source dir
	si, err := os.Stat(source)
	if err != nil {
		return err
	}

	err = os.MkdirAll(dest, si.Mode())
	if err != nil {
		return err
	}

	d, err := os.Open(source)
	if err != nil {
		return err
	}
	defer d.Close()

	objects, err := d.Readdir(-1)

	for _, obj := range objects {

		sp := filepath.Join(source, "/", obj.Name())

		dp := filepath.Join(dest, "/", obj.Name())

		if obj.IsDir() {
			err = CopyDir(sp, dp)
			if err != nil {
				return err
			}
		} else {
			// perform copy
			err = CopyFile(sp, dp)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

// CopyFile copies a source file to a destination.
//
// It follows symbolic links and retains modes.
func CopyFile(source string, dest string) error {
	ln, err := os.Readlink(source)
	if err == nil {
		return os.Symlink(ln, dest)
	}
	s, err := os.Open(source)
	if err != nil {
		return err
	}

	defer s.Close()

	d, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return err
	}

	si, err := os.Stat(source)
	if err != nil {
		return err
	}
	err = os.Chmod(dest, si.Mode())

	return err
}
