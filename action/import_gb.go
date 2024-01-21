package action

import (
	"github.com/Khulnasoft-lab/gopkg/cfg"
	"github.com/Khulnasoft-lab/gopkg/gb"
	"github.com/Khulnasoft-lab/gopkg/msg"
	gpath "github.com/Khulnasoft-lab/gopkg/path"
)

// ImportGB imports GB dependencies into the present gopkg config.
func ImportGB(dest string) {
	base := "."
	config := EnsureConfig()
	if !gb.Has(base) {
		msg.Die("There is no GB manifest to import.")
	}
	deps, err := gb.Parse(base)
	if err != nil {
		msg.Die("Failed to extract GB manifest: %s", err)
	}
	appendImports(deps, config)
	writeConfigToFileOrStdout(config, dest)
}

func appendImports(deps []*cfg.Dependency, config *cfg.Config) {
	if len(deps) == 0 {
		msg.Info("No dependencies added.")
		return
	}

	//Append deps to existing dependencies.
	if err := config.AddImport(deps...); err != nil {
		msg.Die("Failed to add imports: %s", err)
	}
}

// writeConfigToFileOrStdout is a convenience function for import utils.
func writeConfigToFileOrStdout(config *cfg.Config, dest string) {
	if dest != "" {
		if err := config.WriteFile(dest); err != nil {
			msg.Die("Failed to write %s: %s", gpath.GopkgFile, err)
		}
	} else {
		o, err := config.Marshal()
		if err != nil {
			msg.Die("Error encoding config: %s", err)
		}
		msg.Default.Stdout.Write(o)
	}
}
