package main

import M "GoLoadBalancer/model"

func main(){
	loadBalancer := M.InitLoadBalancer()
	M.LaunchInstances(loadBalancer)
}