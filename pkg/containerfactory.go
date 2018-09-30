package pkg

import (
	"fmt"
	"reflect"

	"github.com/port-exporter/pkg/containers/docker"
	"github.com/port-exporter/pkg/containers"
)

// implement utilities for instantiating the supported
// containerDriver

type driverConfigTypes struct {
	DriverType reflect.Type
}

var containerDriverRegistry = map[string]driverConfigTypes{
	DockerNameStr: {
		DriverType: reflect.TypeOf(docker.DockerContainerHandler{}),
	},
}



const (

	// DockerNameStr is a string constant for docker driver
	DockerNameStr = "docker"

)


// initHelper initializes the NetPlugin by mapping driver names to
// configuration, then it imports the configuration.
func initHelper(driverRegistry map[string]driverConfigTypes, driverName string) (interface{}, error) {
	if _, ok := driverRegistry[driverName]; ok {
		driverType := driverRegistry[driverName].DriverType

		driver := reflect.New(driverType).Interface()
		return driver, nil
	}

	return nil, fmt.Errorf("Failed to find a registered driver for: %s", driverName)
}


// NewContainerDriver instantiates a 'named' container-driver with specified configuration
func NewContainerDriver(name string) (containers.ContainerHandler, error) {
	if name == ""  {
		return nil, fmt.Errorf("invalid driver name or configuration passed.")
	}

	driver, err := initHelper(containerDriverRegistry, name)
	if err != nil {
		return nil, err
	}

	d := driver.(containers.ContainerHandler)

	return d, nil
}
