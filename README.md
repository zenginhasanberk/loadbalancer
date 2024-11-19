# Simple Load Balancer in Go

A simple load balancer implementation using the Go standard library. The load balances takes in a list of available `servers` that will be used when balancing the load. Then, using a specified strategy, 
it aims toreduce the burden of a single server handling requests.

