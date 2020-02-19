package mesh_workload

import (
	"context"
	"fmt"
	"strings"

	pb_types "github.com/gogo/protobuf/types"
	core_types "github.com/solo-io/mesh-projects/pkg/api/core.zephyr.solo.io/v1alpha1/types"
	discoveryv1alpha1 "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	discovery_types "github.com/solo-io/mesh-projects/pkg/api/discovery.zephyr.solo.io/v1alpha1/types"
	"github.com/solo-io/mesh-projects/pkg/env"
	"github.com/solo-io/mesh-projects/services/common/constants"
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	DiscoveryLabels = map[string]string{
		constants.DISCOVERED_BY: constants.MESH_WORKLOAD_DISCOVERY,
		constants.MESH_TYPE:     constants.ISTIO,
	}
)

type IstioMeshWorkloadScanner MeshWorkloadScanner

// visible for testing
func NewIstioMeshWorkloadScanner(ownerFetcher OwnerFetcher) IstioMeshWorkloadScanner {
	return &istioMeshWorkloadScanner{
		deploymentFetcher: ownerFetcher,
	}
}

type istioMeshWorkloadScanner struct {
	deploymentFetcher OwnerFetcher
}

func (i *istioMeshWorkloadScanner) ScanPod(ctx context.Context, pod *core_v1.Pod) (*discoveryv1alpha1.MeshWorkload, error) {
	if !i.isIstioPod(pod) {
		return nil, nil
	}
	deployment, err := i.deploymentFetcher.GetDeployment(ctx, pod)
	if err != nil {
		return nil, err
	}
	return &discoveryv1alpha1.MeshWorkload{
		ObjectMeta: v1.ObjectMeta{
			Name:      i.buildMeshWorkloadName(deployment.Name, deployment.Namespace, pod.ClusterName),
			Namespace: env.DefaultWriteNamespace,
			Labels:    DiscoveryLabels,
		},
		Spec: discovery_types.MeshWorkloadSpec{
			KubeControllerRef: &core_types.ResourceRef{
				Kind:      &pb_types.StringValue{Value: deployment.Kind},
				Name:      deployment.Name,
				Namespace: deployment.Namespace,
				Cluster:   &pb_types.StringValue{Value: pod.ClusterName},
			},
		},
	}, nil
}

// iterate through pod's containers and check for one with name containing "istio" and "proxy"
func (i *istioMeshWorkloadScanner) isIstioPod(pod *core_v1.Pod) bool {
	for _, container := range pod.Spec.Containers {
		if strings.Contains(container.Image, "istio") && strings.Contains(container.Image, "proxy") {
			return true
		}
	}
	return false
}

func (i *istioMeshWorkloadScanner) buildMeshWorkloadName(deploymentName string, namespace string, clusterName string) string {
	// TODO: https://github.com/solo-io/mesh-projects/issues/141
	return fmt.Sprintf("%s-%s-%s-%s", "istio", deploymentName, namespace, clusterName)
}