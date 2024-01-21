package action

import (
	"os"

	"github.com/Khulnasoft-lab/gopkg/cache"
	"github.com/Khulnasoft-lab/gopkg/msg"
)

// CacheClear clears the Gopkg cache
func CacheClear() {
	l := cache.Location()

	err := os.RemoveAll(l)
	if err != nil {
		msg.Die("Unable to clear the cache: %s", err)
	}

	cache.SetupReset()
	cache.Setup()

	msg.Info("Gopkg cache has been cleared.")
}
