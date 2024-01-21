package action

import (
	"github.com/Khulnasoft-lab/gopkg/msg"
)

// Name prints the name of the package, according to the gopkg.yaml file.
func Name() {
	conf := EnsureConfig()
	msg.Puts(conf.Name)
}
