package main

import (
	"fmt"
	"loadbalancer/loadbalancer"
)

func main() {
	servers := []string{
		"http://localhost:2500",
		"http://localhost:2501",
		"http://localhost:2502",
	}
	lb := loadbalancer.NewLoadBalancer(servers)
	lb.SetStrategy(loadbalancer.RoundRobin)
	fmt.Println("Load balancer started on :8080")
	lb.Start(":8080")
}
