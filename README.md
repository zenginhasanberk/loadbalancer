# Simple Load Balancer in Go

A simple load balancer implementation using the Go standard library. The load balances takes in a list of available `servers` that will be used when balancing the load. Then, using a specified strategy, 
it aims to reduce the burden of a single server handling requests. The default strategy uses round robin.

