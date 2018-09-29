package collector

import (

	"github.com/prometheus/client_golang/prometheus"
	"docker-interface-exporter/pkg/containers"
	"docker-interface-exporter/pkg/tools"
	"docker-interface-exporter/pkg/cache"
	"regexp"
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

	containerCache  *cache.Cache
	containerHd	    containers.ContainerHandler
	errors	        prometheus.Gauge
	cms             []containerMetric

}

var invalidLabelCharRE = regexp.MustCompile(`[^a-zA-Z0-9_]`)

// sanitizeLabelName replaces anything that doesn't match
// client_label.LabelNameRE with an underscore.
func sanitizeLabelName(name string) string {
	return invalidLabelCharRE.ReplaceAllString(name, "_")
}


func getContaierLabels(c *containers.Containerinfo) map[string] string{

	//containerLabels := make(map[string] string)
	return c.Labels


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

								value := float64(eth.State)
								values = append(values, metricValue{
									value:  value,
									labels: []string{eth.Name},
								})

							}
							return values

						}else{

							return make(metricValues, 0, 0)
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
	//first collect
/*	if self.containerCache == nil{

		self.containerCache = cache.NewCache()
		//init the cache
		containerInfos := self.containerHd.GetContainerInfos()

		self.containerCache.SetFromList(containerInfos)
	}else{


	}*/

	containerInfos := self.containerHd.GetContainerInfos()
	for _, containerInfo := range containerInfos{

		rawLabels := getContaierLabels(containerInfo)
		values := make([]string, 0, len(rawLabels))
		labels := make([]string, 0, len(rawLabels))

		for l, value := range rawLabels {
			labels = append(labels, sanitizeLabelName(l))
			values = append(values, value)
		}

		for _, cm := range self.cms {
			desc := cm.desc(labels)
			for _, metricValue := range cm.getValues(containerInfo) {
				ch <- prometheus.MustNewConstMetric(desc, cm.valueType, float64(metricValue.value), append(values, metricValue.labels...)...)
			}
		}
	}

	return nil

}


