package filter

import (
	"github.com/benjvi/arjuna/provider"
	"github.com/jmespath/go-jmespath"
	"github.com/benjvi/arjuna/assertion"
	"github.com/cloudflare/cfssl/log"
	"fmt"
)


type FilterConfig struct {
	Type		string
	Expression 	string
	Assert		assertion.AssertionConfig
}

func New(config FilterConfig) (Filter,error) {
	filter := &JMESPathFilter{}
	err := filter.init(config)
	if err != nil {
		return nil, err
	}
	return filter, nil
}

type Filter interface {
	init(config FilterConfig)		error
	Run(resources []provider.Resource)	([]provider.Resource, error)
}

type JMESPathFilter struct {
	Expression	*jmespath.JMESPath
	Assertion	assertion.Assertion
	ExpressionString	string
}

func (this *JMESPathFilter) Run(resources []provider.Resource) ([]provider.Resource, error) {
	var filtered []provider.Resource

	for _, res := range resources {
		result,err := this.Expression.Search(res.State())
		if err != nil {
			return nil, err
		}

		// only support value and list types though, for simplicity
		var actual interface{}
		if result != nil {
			log.Info(fmt.Sprintf("JMESPath search result found: %v\n from expression: %s", result.(interface{}), this.ExpressionString))

			actual = (result).(interface{})
		} else {
			// expression not found equivalent to key with value null??
			log.Warning(fmt.Sprintf("Key not found in resource: %+v", res.State()))
		}

		assertionResult, err := this.Assertion.Run(actual)
		if err != nil  {
			log.Info(fmt.Sprintf("Filtered out resource %+v due to error evaluating assertion: %s\n", res.State(), err.Error()))
			continue
		}
		if assertionResult {
			filtered = append(filtered, res)
		} else {
			log.Info(fmt.Sprintf("Filtered out resource %+v due to failed assertion\n", res.State()))
		}
	}
	return filtered, nil
}

func (this *JMESPathFilter) init(config FilterConfig) error {
	log.Debug(fmt.Sprintf("Compiling JMESPATH filter expression: %s\n", config.Expression))
	expr, err  := jmespath.Compile(config.Expression)
	if err != nil {
		return err
	}
	this.ExpressionString = config.Expression
	this.Assertion, err = assertion.NewValueAssertion(config.Assert)
	if err != nil {
		return err
	}
	this.Expression = expr
	return nil
}