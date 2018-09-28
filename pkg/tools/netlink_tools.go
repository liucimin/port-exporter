package tools



import (

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netns"
)



func GetInterfaceStat(ns string) (map[string] netlink.LinkOperState, error){

	//find the interface all in the netns
	interfaceStat := make(map[string] netlink.LinkOperState)
	nh, err :=  netns.GetFromPath(ns)
	if err != nil{

		return interfaceStat, err
	}

	handle, err := netlink.NewHandleAt(nh)
	if err != nil{

		return interfaceStat, err
	}

	//first get all the linklist in the netns
	linkList, err := handle.LinkList()
	if err != nil{

		return interfaceStat, err
	}

	//foreach to get the name and state
	for _, link := range linkList{

		linkAttrs := link.Attrs()
		interfaceStat[linkAttrs.Name] = linkAttrs.OperState

	}
	return interfaceStat, nil
}