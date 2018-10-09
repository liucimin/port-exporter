package ovs

type StatisticSpec struct{
	//collisions=0, rx_bytes=39621276, rx_crc_err=0, rx_dropped=0, rx_errors=0,
	// rx_frame_err=0, rx_over_err=0, rx_packets=638603, tx_bytes=10228557150,
	// tx_dropped=0, tx_errors=0, tx_packets=768449
	Collisions uint64

	RxBytes    uint64
	RxCrcErr   uint64
	RxDropped  uint64
	RxErrors   uint64
	RxFrameErr uint64
	RxOverErr  uint64
	RxPackets  uint64

	TxBytes    uint64
	TxDropped  uint64
	TxErrors   uint64
	TxPackets  uint64



}
type OvsPortInfo struct{

	Uuid	string
	Name    string

	Statistics StatisticSpec

	EndpointId string
	//ExternalIds map[interface{}]interface{}

}