package k8s

import (
	zephyr_discovery_controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1/controller"
	k8s_core_controller "github.com/solo-io/service-mesh-hub/pkg/api/kubernetes/core/v1/controller"
)

//go:generate mockgen -source ./interfaces.go -destination ./mocks/mock_interfaces.go -package service_discovery_mocks

type MeshServiceFinder interface {
	StartDiscovery(serviceEventWatcher k8s_core_controller.ServiceEventWatcher, meshWorkloadEventWatcher zephyr_discovery_controller.MeshWorkloadEventWatcher) error
}
