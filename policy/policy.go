package policy

import (
	"path/filepath"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"strings"
	"github.com/benjvi/arjuna/filter"
	"math/rand"
	"strconv"
	"github.com/benjvi/arjuna/assertion"
	"github.com/benjvi/arjuna/alert"
	"log"
	"errors"
)

func FromDir(dir string) ([]Policy, error) {
	// get all with file ext: .pcy.json in directoty

	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var policies []Policy
	for _, file := range files {
		if file.Mode().IsRegular() {
			name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

			if filepath.Ext(file.Name()) == ".json" &&
				filepath.Ext(name) == ".pcy" {
				fileContent, _ := ioutil.ReadFile(file.Name())

				var pcyConf PolicyConfig
				err := json.Unmarshal(fileContent,  &pcyConf)
				if err != nil {
					log.Fatalf("Unmarshalling policy file %s: %+v\n", file.Name(), err)
				}

				// TODO: ok as a hash??
				// TODO enforce uniqueness of passed values
				if pcyConf.Id == "" {
					pcyConf.Id = strconv.Itoa(rand.Int())
				}

				policy, err := New(pcyConf)
				if err != nil {
					return nil, err
				}
				policies = append(policies, policy)
			}
		}
	}
	return policies, nil
}

type PolicyConfig struct {
	Id           string
	Provider     string
	ResourceType string
	Filters      []filter.FilterConfig
	Alerts       []alert.AlertConfig
	Assert       assertion.AssertionConfig
}

func New(conf PolicyConfig) (empty Policy, err error) {
	pcy := Policy{
		Id:	conf.Id,
		Provider: conf.Provider,
		ResourceType: conf.ResourceType,
	}

	filters := make([]filter.Filter, 0)
	for _, filterConf := range conf.Filters {
		filter, err := filter.New(filterConf)
		if err != nil {
			return empty, errors.New("Error loading filter: "+err.Error())
		}
		filters = append(filters, filter)
	}
	pcy.Filters = filters

	pcy.Assert, err = assertion.NewSetAssertion(conf.Assert)
	if err != nil {
		return empty, err
	}

	alerts := make([]alert.Alert, 0)
	for _, alertConf := range conf.Alerts {
		alert, err := alert.FromConfig(alertConf)
		if err != nil {
			return empty, err
		}
		alerts = append(alerts, alert)
	}
	pcy.Alerts = alerts

	return pcy, nil
}

//TODO Provider and ResourceType should be interfaces types
type Policy struct {
	Id           string
	Provider     string
	ResourceType string
	Filters      []filter.Filter
	Alerts       []alert.Alert
	Assert       assertion.Assertion
}

