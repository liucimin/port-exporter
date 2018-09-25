package docker

import (

	"github.com/golang/glog"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"context"
	"fmt"
	"github.com/docker-interface-exporter/pkg/tools"

)

type DockerContainerHandler struct{


	cli *client.Client


}



func NewContainerHandler() containerHandler{
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return &DockerContainerHandler{

		cli,

	}
}



func (self *DockerContainerHandler)GetContainerInfos() []*Containerinfo {

	var cInfos []*Containerinfo

	//get all the container in the mechine
	containers, err := self.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		glog.Infof("%s %s\n",container.ID, container.Image)

		var cInfo Containerinfo
		//inspect all the container to get container info
		containerJson, err := self.cli.ContainerInspect(context.Background(), container.ID)
		if err != nil {
			glog.Errorf("failed to inspect container %q: %v", container.ID, err)
			continue

		}


		needGetNet := isNeedNetwork(&containerJson)
		//just container net save the pid
		if needGetNet {

			//save the pid path
			cInfo = Containerinfo{
				Id :	container.ID,
				HasNetwork :	true,
				Namespace : getNetworkNamespace(&containerJson),


			}

			//save the interface stat

			tools.getInterfaceStat(cInfo.Namespace)

		}else{

			cInfo =  Containerinfo{
				Id : container.ID,
				HasNetwork : false,
			}

		}


		cInfos = append(cInfos, &cInfo)


	}


}

func (self *DockerContainerHandler)UpdateContainerInfos() {





}




func getNetworkNamespace(c *types.ContainerJSON) string{
	if c.State.Pid == 0 {
		// Docker reports pid 0 for an exited container.
		return ""
	}
	return fmt.Sprintf("/proc/%v", c.State.Pid)
}


func isNeedNetwork(c *types.ContainerJSON) bool{

	return !c.HostConfig.NetworkMode.IsContainer() && !c.HostConfig.NetworkMode.IsHost()
}



