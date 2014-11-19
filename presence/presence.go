// Copyright © 2014-2015, Civis Analytics

package presence

import (
	"fmt"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/elb"
)

func JoinELB(name string) error {
	iIDs := []string{aws.InstanceId()}

	fmt.Printf("Instance[%s] Joining Load-Balancer Pool[%s]\n", iIDs[0], name)
	auth, err := aws.EnvAuth()
	if err != nil {
		return err
	}

	e := elb.New(auth, aws.USEast)

	_, err = e.RegisterInstancesWithLoadBalancer(iIDs, name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func LeaveELB(name string) error {
	iIDs := []string{aws.InstanceId()}

	fmt.Printf("Instance[%s] Leaving Load-Balancer Pool[%s]\n", iIDs[0], name)
	auth, err := aws.EnvAuth()
	if err != nil {
		return err
	}

	e := elb.New(auth, aws.USEast)

	e.DeregisterInstancesFromLoadBalancer(iIDs, name)
	return nil
}
