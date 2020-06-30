package main

import (
	"net/http"
	"os"
	"time"

	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ping "github.com/sparrc/go-ping"
)

///////////////////
// Definitions
///////////////////

// Cluster App class
func NewCluster(MyCluster []Barebone) Cluster {
	var retCluster Cluster
	retCluster.Machines = &MyCluster
	return retCluster
}

type Cluster struct {
	Machines *[]Barebone
	Wg       sync.WaitGroup
}

//RecordMetrics: foreach machine, record metrics
func (c *Cluster) RecordMetrics() {

	for _, m := range *c.Machines {
		m.RecordMetrics()
	}
}

//Register: foreach machine, register metrics in prometheus
func (c *Cluster) RegisterMetrics() {

	for _, m := range *c.Machines {
		m.RegisterMetrics()
	}
}

// Barebone refers to a physical machine
type Barebone struct {
	Hostname string
	Network1 *Network
}

func (b *Barebone) RegisterMetrics() {

	opts := prometheus.GaugeOpts{
		Name: "ping_delay_for_" + b.Hostname,
		Help: "The total RSS time for one ping instance",
	}
	//Register Our PingDelay metric
	b.Network1.PingDelay = promauto.NewGauge(opts)

}

func (b *Barebone) RecordMetrics() {

	//Collect stats from pinging our machine
	b.Network1.DoPing()
}

type Network struct {
	Name      string
	IPAddress string
	PingDelay prometheus.Gauge
}

func (p *Network) DoPing() {

	pinger, _ := ping.NewPinger(p.IPAddress)
	pinger.Count = 1
	pinger.Timeout = time.Second * 10
	pinger.SetPrivileged(true)
	pinger.Run()

	//PingStats in My Machine
	stats := pinger.Statistics()

	p.PingDelay.Set(float64(stats.AvgRtt.Seconds() * 1000))

	fmt.Printf("(DoPing) Pinging %s from hostname %s .. stats returned are %f and stats value\n", p.IPAddress, p.Name, stats.AvgRtt.Seconds())

}

///////////////////
// Variables
///////////////////
var (
	myMachine1 = Barebone{
		Hostname: "client_1",
		Network1: &Network{
			Name:      "Data",
			IPAddress: os.Getenv("MACHINE_1"),
		},
	}

	myMachine2 = Barebone{
		Hostname: "client_2",
		Network1: &Network{
			Name:      "Data",
			IPAddress: os.Getenv("MACHINE_2"),
		},
	}

	cluster = NewCluster([]Barebone{myMachine1, myMachine2})
)

///////////////////
// Main() function
///////////////////
func main() {
	cluster.RegisterMetrics()
	go func() {
		for {
			cluster.RecordMetrics()
			time.Sleep(5 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)

}
