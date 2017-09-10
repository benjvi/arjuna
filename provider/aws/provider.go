package aws

import (
	"github.com/benjvi/arjuna/provider"
	"github.com/aws/aws-sdk-go/aws/session"
)

type AWSProvider struct {}

func NewProvider() provider.Provider {
	return AWSProvider{}
}

func (this AWSProvider) ResourceCollections() map[string]provider.ResourceCollection {
	return map[string]provider.ResourceCollection{
		"ec2_instance": EC2InstanceCollection{
			SessionFactory: this.SessionFactory(),
		},
	}
}

func (AWSProvider) Name() string {
	return "AWS"
}

func (AWSProvider) SessionFactory() func()interface{} {
	return func()interface{} {
		return session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
	}
}
