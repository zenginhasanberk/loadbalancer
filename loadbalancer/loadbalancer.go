package loadbalancer

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Strategy int

const (
	RoundRobin Strategy = iota
)

type Server struct {
	Alive bool
	RP    *httputil.ReverseProxy
	URL   string
}

type LoadBalancer struct {
	servers  []*Server
	index    int
	strategy http.HandlerFunc
}

func NewLoadBalancer(servers []string) *LoadBalancer {
	serverObjects := make([]*Server, 0, len(servers))

	for _, server := range servers {
		target, _ := url.Parse(server)
		proxy := httputil.NewSingleHostReverseProxy(target)
		serverObjects = append(serverObjects, &Server{Alive: true, RP: proxy, URL: server})
	}

	return &LoadBalancer{
		servers: serverObjects,
		index:   0,
	}
}

func (lb *LoadBalancer) SetStrategy(strat Strategy) {
	strategies := []http.HandlerFunc{
		// Round Robin
		func(w http.ResponseWriter, r *http.Request) {
			lb.index = (lb.index + 1) % len(lb.servers)
			server := lb.servers[lb.index]
			fmt.Println("New Index is ", lb.index, " now serving ", server)
			server.RP.ServeHTTP(w, r)
		},
	}
	lb.strategy = strategies[int(strat)]
}

func (lb *LoadBalancer) Start(addr string) {
	http.HandleFunc("/", lb.strategy)
	http.ListenAndServe(addr, nil)
}
