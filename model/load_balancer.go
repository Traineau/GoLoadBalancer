package model

import "fmt"

type loadBalancer struct {
	instances	[]*instance
}

func InitInstances(){
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
	//loadBalancer := loadBalancer{instances:myInstances}
	printInstance(myInstances)
}

func printInstance(instances []*instance){
	for i := 0; i < len(instances); i++ {
		fmt.Printf("%v", instances[i])
	}
}

func sendMessages(){

}

func hearthBeat(){

}
