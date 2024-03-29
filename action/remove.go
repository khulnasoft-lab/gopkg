package action

import (
	"github.com/Khulnasoft-lab/gopkg/cache"
	"github.com/Khulnasoft-lab/gopkg/cfg"
	"github.com/Khulnasoft-lab/gopkg/msg"
	gpath "github.com/Khulnasoft-lab/gopkg/path"
	"github.com/Khulnasoft-lab/gopkg/repo"
)

// Remove removes a dependncy from the configuration.
func Remove(packages []string, inst *repo.Installer) {
	cache.SystemLock()
	base := gpath.Basepath()
	EnsureGopath()
	EnsureVendorDir()
	conf := EnsureConfig()
	gopkgfile, err := gpath.Gopkg()
	if err != nil {
		msg.Die("Could not find Gopkg file: %s", err)
	}

	msg.Info("Preparing to remove %d packages.", len(packages))
	conf.Imports = rmDeps(packages, conf.Imports)
	conf.DevImports = rmDeps(packages, conf.DevImports)

	// Copy used to generate locks.
	confcopy := conf.Clone()

	//confcopy.Imports = inst.List(confcopy)

	if err := repo.SetReference(confcopy, inst.ResolveTest); err != nil {
		msg.Err("Failed to set references: %s", err)
	}

	err = inst.Export(confcopy)
	if err != nil {
		msg.Die("Unable to export dependencies to vendor directory: %s", err)
	}

	// Write gopkg.yaml
	if err := conf.WriteFile(gopkgfile); err != nil {
		msg.Die("Failed to write gopkg YAML file: %s", err)
	}

	// Write gopkg lock
	writeLock(conf, confcopy, base)
}

// rmDeps returns a list of dependencies that do not contain the given pkgs.
//
// It generates neither an error nor a warning for a pkg that does not exist
// in the list of deps.
func rmDeps(pkgs []string, deps []*cfg.Dependency) []*cfg.Dependency {
	res := []*cfg.Dependency{}
	for _, d := range deps {
		rem := false
		for _, p := range pkgs {
			if p == d.Name {
				rem = true
			}
		}
		if !rem {
			res = append(res, d)
		}
	}
	return res
}
