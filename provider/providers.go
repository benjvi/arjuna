package provider

import (
	"errors"
)

type Provider interface {
	Name() string
	ResourceCollections() map[string]ResourceCollection
	SessionFactory() func()interface{}
}

type ResourceCollection interface {
	List() []Resource
	Actions() map[string]func()error
	// sCollectionActions() map[string]func()error
}

type Resource interface {
	State() interface{}
	Summary() string
}

func New(name string, providerFactory map[string]func()Provider) (Provider, error) {
	if constructorFunc, ok := providerFactory[name] ; ok {
		provider := constructorFunc()
		return provider, nil
	} else {
		var empty Provider
		return empty, errors.New("Unknown provider")
	}
}
