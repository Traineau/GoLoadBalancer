package model

import (
	"fmt"
	"time"
)

type loadBalancer struct {
	instances	[]*instance
}

func InitLoadBalancer() loadBalancer{
	var myInstances []*instance
	for i := 0; i < 10 ; i++  {
		var app = "helloWorld"
		var address = "/helloWorld"
		if i % 2 == 0 {
			app = "helloFrance"
			address = "/helloFrance"
		}

		anInstance := instance{
			address: address,
			app: app,
			port: 8000 + i}

		myInstances = append(myInstances, &anInstance)
	}

	printInstances(myInstances)
	loadBalancer := loadBalancer{instances:myInstances}

	return loadBalancer
}

func printInstances(instances []*instance){
	for i := range instances {
		fmt.Printf("%+v \n", *instances[i])
	}
}

func LaunchInstances(loadBalancer loadBalancer){
	instances := loadBalancer.instances
	for i := range loadBalancer.instances {
		time.Sleep(100 * time.Millisecond)
		if instances[i].app == "helloWorld" {
			go launchHelloWorld(*instances[i])
		}else{
			go launchHelloFrance(*instances[i])
		}
	}
	time.Sleep(10000 * time.Hour)
}

func sendMessages(){

}

func hearthBeat(){

}
