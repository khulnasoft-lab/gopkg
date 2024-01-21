package action

import (
	"io/ioutil"
	"testing"

	"github.com/Khulnasoft-lab/gopkg/cfg"
	"github.com/Khulnasoft-lab/gopkg/msg"
)

func TestAddPkgsToConfig(t *testing.T) {
	// Route output to discard so it's not displayed with the test output.
	o := msg.Default.Stderr
	msg.Default.Stderr = ioutil.Discard

	conf := new(cfg.Config)
	dep := new(cfg.Dependency)
	dep.Name = "github.com/Khulnasoft-lab/gococ"
	dep.Subpackages = append(dep.Subpackages, "convert")
	conf.Imports = append(conf.Imports, dep)

	names := []string{
		"github.com/Khulnasoft-lab/gococ/fmt",
		"github.com/Khulnasoft-lab/goctl-semver",
	}

	addPkgsToConfig(conf, names, false, true, false)

	if !conf.HasDependency("github.com/Khulnasoft-lab/goctl-semver") {
		t.Error("addPkgsToConfig failed to add github.com/Khulnasoft-lab/goctl-semver")
	}

	d := conf.Imports.Get("github.com/Khulnasoft-lab/gococ")
	found := false
	for _, s := range d.Subpackages {
		if s == "fmt" {
			found = true
		}
	}
	if !found {
		t.Error("addPkgsToConfig failed to add subpackage to existing import")
	}

	// Restore messaging to original location
	msg.Default.Stderr = o
}
