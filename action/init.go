package action

import (
	gpath "github.com/Khulnasoft-lab/gopkg/path"
)

// Init initializes the action subsystem for handling one or more subesequent actions.
func Init(yaml, home string) {
	gpath.GopkgFile = yaml
	gpath.SetHome(home)
}
