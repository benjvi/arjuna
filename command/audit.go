package command

import (
	"fmt"
	"github.com/benjvi/arjuna/policy"
	"github.com/benjvi/arjuna/provider"

	"os"
	"log"
	"github.com/benjvi/arjuna/provider/aws"
)

func Audit(policyDir string) {

	policies, err := policy.FromDir(policyDir)
	if err != nil {
		log.Fatalf("Error loading policies: %+v", err)
	}

	// request resource types
	resourceColls := getResourceCollections(policies)

	for _, resourceC := range resourceColls {
		resources := resourceC.List()
		for _, resource := range resources {
			fmt.Print(resource.Summary())
		}
	}
}

func getResourceCollections(policies []policy.Policy) map[string]provider.ResourceCollection {
	resourceColls := make(map[string]provider.ResourceCollection)
	for _, p := range policies {
		prov, err := provider.New(p.Provider, providerFactory)
		if err != nil {
			log.Fatalf("Provider not found: %s", p.Provider)
		}
		resourceC, ok := prov.ResourceCollections()[p.ResourceType]
		if !ok {
			log.Fatalf("Resource not found for provider %s: %s", p.Provider, p.ResourceType)
			os.Exit(1)
		}
		// add to map
		resourceColls[p.Id] = resourceC
	}
	return resourceColls
}

var providerFactory = map[string]func()provider.Provider{
	"aws": aws.NewProvider,
}

