package containers


//use this tool to get all containers pid


type ContainerHandler interface{


	GetContainerInfos() []*Containerinfo
	UpdateContainerInfos()

}
