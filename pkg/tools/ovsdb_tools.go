package tools

import (
	"github.com/socketplane/libovsdb"
	"github.com/port-exporter/pkg/cache"
	"github.com/port-exporter/pkg/ovs"
	"reflect"
	"github.com/golang/glog"
)



type OvsdbHandler struct{

	// OVS client
	ovsClient *libovsdb.OvsdbClient

	// OVSDB cache
	ovsdbCache  *cache.Cache  //map[string]map[string]libovsdb.Row

}




// Create a new OVS driver
func NewOvsdbHandler() *OvsdbHandler {
	ovsHandler := new(OvsdbHandler)

	// connect to OVS
	ovs, err := libovsdb.Connect("localhost", 6640)
	if err != nil {
		glog.Fatal("Failed to connect to ovsdb")
	}

	// Setup state
	ovsHandler.ovsClient = ovs

	ovsHandler.ovsdbCache = cache.NewCache()


	// Register for notifications
	ovs.Register(ovsHandler)

	// Populate initial state into cache
	initial, _ := ovs.MonitorAll("Open_vSwitch", "")
	ovsHandler.populateCache(*initial)


	// Return the new OVS driver
	return ovsHandler
}


// Populate local cache of ovs state
func (self *OvsdbHandler) populateCache(updates libovsdb.TableUpdates) {

	for table, tableUpdate := range updates.Updates {

		glog.V(4).Infof("table %s",table)

		if _, ok := self.ovsdbCache.Get(table); !ok {
			self.ovsdbCache.Set(table, make(map[libovsdb.UUID]libovsdb.Row))
		}
		for uuid, row := range tableUpdate.Rows {
			empty := libovsdb.Row{}
			if !reflect.DeepEqual(row.New, empty) {

				if tableCache, ok := self.ovsdbCache.Get(table); ok {
					tableCache.(map[libovsdb.UUID]libovsdb.Row)[libovsdb.UUID{GoUuid: uuid}] = row.New
				}

			} else {
				if tableCache, ok := self.ovsdbCache.Get(table); ok {

					delete(tableCache.(map[libovsdb.UUID]libovsdb.Row), libovsdb.UUID{GoUuid: uuid})
				}

			}
		}
	}
}


func (self *OvsdbHandler) GetInterfaces() []*ovs.OvsPortInfo{

	ovsPortInfoList := []*ovs.OvsPortInfo{}
	//first get Interface table
	tmpInterfaceTable, ok := self.ovsdbCache.Get("Interface")
	if ok{

		//then type change
		if interfaceTable, ok := tmpInterfaceTable.(map[libovsdb.UUID]libovsdb.Row); ok{

			//for each row
			for intfUUID, row := range interfaceTable{

				//get statistics in each row
				if tmpStatistics, ok := row.Fields["statistics"].(libovsdb.OvsMap); ok{

					ovsPortInfo := ovs.OvsPortInfo{
						Uuid:       intfUUID.GoUuid,
						Name:       row.Fields["name"].(string),
						Statistics: ovs.StatisticSpec{},
					}
					for key, value := range tmpStatistics.GoMap{
						key, _ := key.(string)
						switch key {
						case "collisions":
							//
							ovsPortInfo.Statistics.Collisions = value.(float64)
						case "rx_bytes":
							//
							ovsPortInfo.Statistics.RxBytes = value.(float64)

						case "rx_crc_err":
							//
							ovsPortInfo.Statistics.RxCrcErr = value.(float64)
						case "rx_dropped":
							//
							ovsPortInfo.Statistics.RxDropped = value.(float64)
						case "rx_errors":
							//
							ovsPortInfo.Statistics.RxErrors = value.(float64)
						case "rx_frame_err":
							//
							ovsPortInfo.Statistics.RxFrameErr = value.(float64)
						case "rx_over_err":
							//
							ovsPortInfo.Statistics.RxOverErr = value.(float64)
						case "rx_packets":
							//
							ovsPortInfo.Statistics.RxPackets = value.(float64)
						case "tx_bytes":
							//
							ovsPortInfo.Statistics.TxBytes = value.(float64)
						case "tx_dropped":
							//
							ovsPortInfo.Statistics.TxDropped = value.(float64)
						case "tx_errors":
							//
							ovsPortInfo.Statistics.TxErrors = value.(float64)
						case "tx_packets":
							//
							ovsPortInfo.Statistics.TxPackets = value.(float64)
						}
					}

					tmpExternalIds, _ := row.Fields["external_ids"].(libovsdb.OvsMap)
					for key, value := range tmpExternalIds.GoMap{
						key, _ := key.(string)
						switch key {
						case "endpoint-id":
							ovsPortInfo.EndpointId = value.(string)
						}
					}
					tmpState, _ := row.Fields["link_state"].(string)
					if tmpState == "up"{

						ovsPortInfo.State = 1.0
					}

					ovsPortInfoList = append(ovsPortInfoList, &ovsPortInfo)

				}
			}
		}

	}
	return ovsPortInfoList
}
// ************************ Notification handler for OVS DB changes ****************
func (self *OvsdbHandler) Update(context interface{}, tableUpdates libovsdb.TableUpdates) {
	// fmt.Printf("Received OVS update: %+v\n\n", tableUpdates)
	self.populateCache(tableUpdates)
}
func (self *OvsdbHandler) Disconnected(ovsClient *libovsdb.OvsdbClient) {
	glog.Errorf("OVS BD client disconnected")
}
func (self *OvsdbHandler) Locked([]interface{}) {
}
func (self *OvsdbHandler) Stolen([]interface{}) {
}
func (self *OvsdbHandler) Echo([]interface{}) {
}
