package containers

import "time"

type CpuSpec struct {
	// Requested cpu shares. Default is 1024.
	Limit uint64 `json:"limit"`
	// Requested cpu hard limit. Default is unlimited (0).
	// Units: milli-cpus.
	MaxLimit uint64 `json:"max_limit"`
	// Cpu affinity mask.
	// TODO(rjnagal): Add a library to convert mask string to set of cpu bitmask.
	Mask string `json:"mask,omitempty"`
	// CPUQuota Default is disabled
	Quota uint64 `json:"quota,omitempty"`
	// Period is the CPU reference time in ns e.g the quota is compared aginst this.
	Period uint64 `json:"period,omitempty"`
}

type MemorySpec struct {
	// The amount of memory requested. Default is unlimited (-1).
	// Units: bytes.
	Limit uint64 `json:"limit,omitempty"`

	// The amount of guaranteed memory.  Default is 0.
	// Units: bytes.
	Reservation uint64 `json:"reservation,omitempty"`

	// The amount of swap space requested. Default is unlimited (-1).
	// Units: bytes.
	SwapLimit uint64 `json:"swap_limit,omitempty"`
}

// Spec for custom metric.
type MetricSpec struct {
	// The name of the metric.
	Name string `json:"name"`

	// Type of the metric.
	Type string `json:"type"`

	// Data Type for the stats.
	Format string `json:"format"`

	// Display Units for the stats.
	Units string `json:"units"`
}

type InterfaceStats struct {
	// The name of the interface.
	Name string `json:"name"`
	// The stat of the interface up or down.
	State string `json:"state"`
	// Cumulative count of bytes received.
	RxBytes uint64 `json:"rx_bytes"`
	// Cumulative count of packets received.
	RxPackets uint64 `json:"rx_packets"`
	// Cumulative count of receive errors encountered.
	RxErrors uint64 `json:"rx_errors"`
	// Cumulative count of packets dropped while receiving.
	RxDropped uint64 `json:"rx_dropped"`
	// Cumulative count of bytes transmitted.
	TxBytes uint64 `json:"tx_bytes"`
	// Cumulative count of packets transmitted.
	TxPackets uint64 `json:"tx_packets"`
	// Cumulative count of transmit errors encountered.
	TxErrors uint64 `json:"tx_errors"`
	// Cumulative count of packets dropped while transmitting.
	TxDropped uint64 `json:"tx_dropped"`
}



type NetworkStats struct {
	// Network stats by interface.
	Interfaces []InterfaceStats `json:"interfaces,omitempty"`
	// TCP connection stats (Established, Listen...)
	/*Tcp TcpStat `json:"tcp"`
	// TCP6 connection stats (Established, Listen...)
	Tcp6 TcpStat `json:"tcp6"`
	// UDP connection stats
	Udp v1.UdpStat `json:"udp"`
	// UDP6 connection stats
	Udp6 v1.UdpStat `json:"udp6"`*/
}


type Containerinfo struct{

	Id string `json:"id"`

	// Time at which the container was created.
	CreationTime time.Time `json:"creation_time,omitempty"`

	Name string
	// Other names by which the container is known within a certain namespace.
	// This is unique within that namespace.
	Aliases []string `json:"aliases,omitempty"`

	// Namespace is the container's pid path,such as ./proc/%v/ns/
	Namespace string `json:"namespace,omitempty"`

	// Metadata labels associated with this container.
	Labels map[string]string `json:"labels,omitempty"`
	// Metadata envs associated with this container. Only whitelisted envs are added.
	Envs map[string]string `json:"envs,omitempty"`

	HasCpu bool    `json:"has_cpu,omitempty"`
	Cpu    CpuSpec `json:"cpu,omitempty"`

	HasMemory bool       `json:"has_memory"`
	Memory    MemorySpec `json:"memory,omitempty"`

	HasCustomMetrics bool            `json:"has_custom_metrics"`
	CustomMetrics    []MetricSpec `json:"custom_metrics,omitempty"`

	// Following resources have no associated spec, but are being isolated.
	HasNetwork    bool `json:"has_network"`
	Network       NetworkStats `json:"network"`

	HasFilesystem bool `json:"has_filesystem,omitempty"`
	HasDiskIo     bool `json:"has_diskio,omitempty"`

	// Image name used for this container.
	Image string `json:"image,omitempty"`


}
