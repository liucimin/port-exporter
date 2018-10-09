package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

// metricValue describes a single metric value for a given set of label values
// within a parent containerMetric.
type metricValue struct {
	value  float64
	labelValues []string
}

type metricValues []metricValue

var (
	collectorMap = make(map[string]Collector)
)

func registerCollector(collectorName string, collectorInstance Collector) {

	collectorMap[collectorName] = collectorInstance
}

// Collector is the interface a collector has to implement.
type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Collect(ch chan<- prometheus.Metric) error
}

// PortCollector implements the prometheus.Collector interface.
type PortCollector struct {
	Collectors map[string]Collector
}
// NewCollector creates a new Collector
func NewPortCollector(cs ...string) (*PortCollector, error) {
	collectors := make(map[string]Collector)
	for _, c := range cs {
		collector, ok := collectorMap[c]
		if !ok {
			return nil, fmt.Errorf("missing collector: %s", c)
		}

		collectors[c] = collector

	}

	return &PortCollector{Collectors: collectors}, nil

}


// Describe implements the prometheus.Collector interface.
func (n PortCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc(
		"container_interface_exporter",
		"container_interface_exporter: A collector of containers.",
		[]string{"collector"},
		nil,
	)
	//ch <- scrapeSuccessDesc
}

// Collect implements the prometheus.Collector interface.
func (n PortCollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(n.Collectors))
	for name, c := range n.Collectors {
		go func(name string, c Collector) {
			c.Collect(ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
}



