package alert

import (
	"github.com/benjvi/arjuna/provider"
	"fmt"
	"encoding/json"
)

type AlertConfig struct {
	Channel		string
	Config		json.RawMessage
}

type AlertInfo struct {
	PolicyId string
	Resources []provider.Resource
	ResourceErrors []error
	AssertionError error
}

type Alert interface {
	Init(config AlertConfig) error
	Fire(info AlertInfo)		error
}

func FromConfig(conf AlertConfig) (Alert, error) {
	alert := &ConsoleAlert{}
	alert.Init(conf)
	return alert, nil
}

type ConsoleAlert struct {}

func (this *ConsoleAlert) Init(config AlertConfig) error {
	return nil
}

func (this *ConsoleAlert) Fire(info AlertInfo) error {
	if info.AssertionError != nil {
		fmt.Printf("Policy: %s - Assertion failed with error: %v\n", info.PolicyId, info.AssertionError)
	}
	fmt.Printf("Policy:  %s - Resources after filtering:\n", info.PolicyId)
	for _, resource := range info.Resources {
		fmt.Println(resource.Summary())
	}
	return nil
}
