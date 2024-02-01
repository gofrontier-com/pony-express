package apply

import (
	"github.com/gofrontier-com/go-utils/output"
	gocache "github.com/patrickmn/go-cache"
)

var cache gocache.Cache

func init() {
	cache = *gocache.New(gocache.NoExpiration, gocache.NoExpiration)
}

func ApplyADF(configDir string, subscriptionId string, dryRun bool) error {

	output.PrintlnfInfo("Loading and validating Azure RM config from %s", configDir)
	return nil

}
