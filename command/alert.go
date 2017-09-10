package command

import (
	"github.com/benjvi/arjuna/policy"
	"github.com/cloudflare/cfssl/log"
	"github.com/benjvi/arjuna/alert"
	"fmt"
)

func Alert(policyDir string) {

	policies, err := policy.FromDir(policyDir)
	if err != nil {
		log.Fatalf("Error loading policies: %+v", err)
	}
	//fmt.Printf("Got policies: %v\n", policies)

	// request resource types
	resourceColls := getResourceCollections(policies)

	for _, policy := range policies {
		var errors []error
		resourceC := resourceColls[policy.Id]
		resources := resourceC.List()

		for _, filter := range policy.Filters {
			resources, err = filter.Run(resources)
			if err != nil {
				log.Fatalf("Running filter failed: %+v", err)
				errors = append(errors, err)
				break
			}
		}

		interfaces := make([]interface{}, len(resources))
		for i := range resources {
			interfaces[i] = resources[i]
		}
		result, err := policy.Assert.Run(interfaces)
		if err != nil || len(errors) > 0 {
			log.Info(fmt.Sprintf("Policy: %s - Assertion failed", policy.Id))
			alertInfo := alert.AlertInfo{
				PolicyId: policy.Id,
				Resources: resources,
				ResourceErrors: errors,
				AssertionError: err,
			}
			for _, alert := range policy.Alerts {
				alert.Fire(alertInfo)
			}
		} else {
			if result {
				log.Info(fmt.Sprintf("Policy: %s - Assertion succeeded", policy.Id))
			} else {
				log.Info(fmt.Sprintf("Policy: %s - Assertion failed", policy.Id))
				alertInfo := alert.AlertInfo{
					PolicyId: policy.Id,
					Resources: resources,
					ResourceErrors: errors,
					AssertionError: err,
				}
				for _, alert := range policy.Alerts {
					alert.Fire(alertInfo)
				}
			}
		}

	}
}



