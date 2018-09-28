package docker

import (

	"github.com/golang/glog"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"context"
	"fmt"
	"docker-interface-exporter/pkg/tools"
	"docker-interface-exporter/pkg/containers"

)

const (
	KubernetesPodNameLabel       = "io.kubernetes.pod.name"
	KubernetesPodNamespaceLabel  = "io.kubernetes.pod.namespace"
	KubernetesPodUIDLabel        = "io.kubernetes.pod.uid"
	KubernetesContainerNameLabel = "io.kubernetes.container.name"
	KubernetesContainerTypeLabel = "io.kubernetes.container.type"
)


type DockerContainerHandler struct{


	cli *client.Client


}



func NewContainerHandler() containers.ContainerHandler{
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return &DockerContainerHandler{

		cli,

	}
}



func (self *DockerContainerHandler)GetContainerInfos() []*containers.Containerinfo {

	var cInfos []*containers.Containerinfo

	//get all the container in the mechine
	containers, err := self.cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		glog.Infof("%s %s\n",container.ID, container.Image)

		var cInfo containers.Containerinfo
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
			cInfo = containers.Containerinfo{
				Id :	container.ID,
				HasNetwork :	true,
				Namespace : getNetworkNamespace(&containerJson),


			}

			//save the interface stat

			tools.GetInterfaceStat(cInfo.Namespace)

		}else{

			cInfo =  containers.Containerinfo{
				Id : container.ID,
				HasNetwork : false,
			}

		}
		var podName, containerName string

		if v, ok := containerJson.Config.Labels[KubernetesPodNameLabel]; ok {
			podName = v
		}

		containerName = containerJson.Name

		cInfo.Name = containerName
		cInfo.Aliases = []string{podName}

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


