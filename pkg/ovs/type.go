package ovs

type StatisticSpec struct{
	//collisions=0, rx_bytes=39621276, rx_crc_err=0, rx_dropped=0, rx_errors=0,
	// rx_frame_err=0, rx_over_err=0, rx_packets=638603, tx_bytes=10228557150,
	// tx_dropped=0, tx_errors=0, tx_packets=768449
	Collisions float64

	RxBytes    float64
	RxCrcErr   float64
	RxDropped  float64
	RxErrors   float64
	RxFrameErr float64
	RxOverErr  float64
	RxPackets  float64

	TxBytes    float64
	TxDropped  float64
	TxErrors   float64
	TxPackets  float64



}
type OvsPortInfo struct{

	Uuid	string
	Name    string
	State   float64
	Statistics StatisticSpec

	EndpointId string
	//ExternalIds map[interface{}]interface{}

}