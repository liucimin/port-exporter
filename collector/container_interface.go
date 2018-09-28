package collector

import (

	"github.com/prometheus/client_golang/prometheus"
	"docker-interface-exporter/pkg/containers"
	"docker-interface-exporter/pkg/tools"
	"strconv"
)
func init() {
	registerCollector("container_interface", NewContainerInterfaceCollector())
}

type containerMetric struct {
	name        string
	help        string
	valueType   prometheus.ValueType
	extraLabels []string
	getValues   func(s *containers.Containerinfo) metricValues
}

func (cm *containerMetric) desc(baseLabels []string) *prometheus.Desc {
	return prometheus.NewDesc(cm.name, cm.help, append(baseLabels, cm.extraLabels...), nil)
}


type  ContainerInterfaceCollector struct{

	containerHd	    containers.ContainerHandler
	errors	        prometheus.Gauge
	cms             []containerMetric

}


func getContaierLabels(c *containers.Containerinfo) map[string] string{

	containerLabels := make(map[string] string)



}

// NewARPCollector returns a new Collector exposing ARP stats.
func NewContainerInterfaceCollector() (Collector) {


	if ch, err := tools.NewContainerDriver("docker");err == nil{

		return &ContainerInterfaceCollector{
			containerHd: ch,
			cms: []containerMetric{
				{//now just container interface state
					name:      "container_interface_state",
					help:      "The interface states in all container",
					valueType: prometheus.GaugeValue,
					extraLabels: []string{"interface"},
					getValues: func(c *containers.Containerinfo) metricValues {

						if c.HasNetwork {
							values := make(metricValues, 0, len(c.Network.Interfaces))
							for _, eth := range c.Network.Interfaces{

								value, _ := strconv.ParseFloat(eth.State, 64)
								values = append(values, metricValue{
									value:  value,
									labels: []string{eth.Name},
								})

							}
							return values

						}else{

							return nil
						}




					},

				},
			},
		}

	}else{

		return nil
	}

}

func (self *ContainerInterfaceCollector)Collect(ch chan<- prometheus.Metric) error  {




}