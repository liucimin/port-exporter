package collector

import (

	"github.com/prometheus/client_golang/prometheus"
	"github.com/golang/glog"
	"github.com/port-exporter/pkg/tools"
	"github.com/port-exporter/pkg/cache"
	"github.com/port-exporter/pkg/ovs"
)
func init() {
	registerCollector("ovs_db", NewOvsdbInterfaceCollector())
}

type ovsdbMetric struct {
	name        string
	help        string
	valueType   prometheus.ValueType
	extraLabels []string
	getValues   func(s *ovs.OvsPortInfo) metricValues
}

func (om *ovsdbMetric) desc(baseLabels []string) *prometheus.Desc {
	return prometheus.NewDesc(om.name, om.help, append(baseLabels, om.extraLabels...), nil)
}


type  OvsdbInterfaceCollector struct{

	ovsdbCache  *cache.Cache
	ovsdbHd	    *tools.OvsdbHandler
	errors	    prometheus.Gauge
	oms         []ovsdbMetric

}



// NewContainerInterfaceCollector returns a new Collector exposing Container stats.
func NewOvsdbInterfaceCollector() (Collector) {

	ch := tools.NewOvsdbHandler()


	return &OvsdbInterfaceCollector{
		ovsdbHd: ch,
		oms: []ovsdbMetric{
			{
				name:      "ovs_port_rx_byte_state",
				help:      "The ovs port rx byte states",
				valueType: prometheus.GaugeValue,
				extraLabels: []string{"name", "endpointId"},
				getValues: func(c *ovs.OvsPortInfo) metricValues {

						values := make(metricValues, 0, 1)
						values = append(values, metricValue{
							value:  float64(c.Statistics.RxBytes),
							labelValues: []string{c.Name, c.EndpointId},
						})
						return values
				},

			},
			{
				name:      "ovs_port_rx_packet_state",
				help:      "The ovs port rx packet states",
				valueType: prometheus.GaugeValue,
				extraLabels: []string{"name", "endpointId"},
				getValues: func(c *ovs.OvsPortInfo) metricValues {

					values := make(metricValues, 0, 1)
					values = append(values, metricValue{
						value:  float64(c.Statistics.RxPackets),
						labelValues: []string{c.Name, c.EndpointId},
					})
					return values
				},

			},
			{
				name:      "ovs_port_tx_byte_state",
				help:      "The ovs port tx byte states",
				valueType: prometheus.GaugeValue,
				extraLabels: []string{"name", "endpointId"},
				getValues: func(c *ovs.OvsPortInfo) metricValues {

					values := make(metricValues, 0, 1)
					values = append(values, metricValue{
						value:  float64(c.Statistics.TxBytes),
						labelValues: []string{c.Name, c.EndpointId},
					})
					return values
				},

			},
			{
				name:      "ovs_port_tx_packet_state",
				help:      "The ovs port tx packet states",
				valueType: prometheus.GaugeValue,
				extraLabels: []string{"name", "endpointId"},
				getValues: func(c *ovs.OvsPortInfo) metricValues {

					values := make(metricValues, 0, 1)
					values = append(values, metricValue{
						value:  float64(c.Statistics.TxPackets),
						labelValues: []string{c.Name, c.EndpointId},
					})
					return values
				},

			},

		},
	}

}

func (self *OvsdbInterfaceCollector)Collect(ch chan<- prometheus.Metric) error  {

	OvsPortInfos := self.ovsdbHd.GetInterfaces()
	glog.Infof("OvsPortInfos %v",OvsPortInfos)
	for _, ovsPortInfo := range OvsPortInfos{

		for _, om := range self.oms {
			desc := om.desc([]string{})
			for _, metricValue := range om.getValues(ovsPortInfo) {
				ch <- prometheus.MustNewConstMetric(desc, om.valueType, float64(metricValue.value), metricValue.labelValues...)
			}
		}
	}

	return nil

}


