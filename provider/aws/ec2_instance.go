package aws

import (
	"github.com/benjvi/arjuna/provider"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"encoding/json"
)

type EC2InstanceCollection struct {
	Name	string
	SessionFactory func()interface{}
}

func (this EC2InstanceCollection) List() []provider.Resource {
	sess  := this.SessionFactory().(*session.Session)

	// Create new EC2 client
	ec2Svc := ec2.New(sess)

	result, err := ec2Svc.DescribeInstances(nil)
	if err != nil {
		log.Fatalf("EC2 DEscribeInstances failed with err: %s", err)
	}

	var resources []provider.Resource
	for _, resv := range result.Reservations {
		for _, i := range resv.Instances{
			r := EC2Instance{}
			r.state = i
			resources = append(resources, r)
		}
	}

	return resources
}

func (EC2InstanceCollection) Actions() map[string]func()error {
	return map[string]func()error {

	}
}

type EC2Instance struct {
	state	*ec2.Instance
}

func (this EC2Instance) State() interface{} {
	return this.state
}

func (this EC2Instance) Summary() string {
	s := make(map[string]interface{})
	s["id"] = *this.state.InstanceId
	bytes, _ := json.Marshal(s)
	return string(bytes)
}