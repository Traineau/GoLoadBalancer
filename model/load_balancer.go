package model

import "fmt"

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
			port: string(8000 + i)}

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
		if instances[i].app == "helloWorld" {
			launchHelloWorld(*instances[i])
		}else{
			launchHelloFrance(*instances[i])
		}
	}
}

func sendMessages(){

}

func hearthBeat(){

}
