package client

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/ohayocorp/anemos/pkg/core"
	"github.com/ohayocorp/anemos/pkg/util"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/version"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/kubectl/pkg/cmd/apply"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

const (
	ManagedByLabel = "app.kubernetes.io/managed-by"
)

type KubernetesClient struct {
	Factory       cmdutil.Factory
	Mapper        meta.RESTMapper
	CoreClient    kubernetes.Interface
	DynamicClient dynamic.Interface
}

type ClusterInfo struct {
	Version      string
	Distribution core.KubernetesDistribution
}

func NewKubernetesClient() (*KubernetesClient, error) {
	getter := genericclioptions.NewConfigFlags(true)
	factory := cmdutil.NewFactory(getter)

	mapper, err := factory.ToRESTMapper()
	if err != nil {
		return nil, fmt.Errorf("failed to get REST mapper: %w", err)
	}

	dynamicClient, err := factory.DynamicClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create dynamic client: %w", err)
	}

	coreClient, err := factory.KubernetesClientSet()
	if err != nil {
		return nil, fmt.Errorf("failed to create core Kubernetes client: %w", err)
	}

	return &KubernetesClient{
		Factory:       factory,
		Mapper:        mapper,
		CoreClient:    coreClient,
		DynamicClient: dynamicClient,
	}, nil
}

func (client *KubernetesClient) GetClusterInfo() (*ClusterInfo, error) {
	config, err := client.Factory.ToRESTConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get REST config: %w", err)
	}

	version, err := client.CoreClient.Discovery().ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to get server version: %w", err)
	}

	distribution := client.detectDistribution(config, version)
	if distribution == nil {
		distribution = core.Pointer(core.KubernetesDistributionKubeadm)
	}

	return &ClusterInfo{
		Version:      version.String(),
		Distribution: *distribution,
	}, nil
}

func (client *KubernetesClient) detectDistribution(config *rest.Config, version *version.Info) *core.KubernetesDistribution {
	// Check for OpenShift
	_, resourcesErr := client.CoreClient.Discovery().ServerResourcesForGroupVersion("route.openshift.io/v1")
	if resourcesErr == nil {
		return core.Pointer(core.KubernetesDistributionOpenShift)
	}

	// Check for GKE
	if strings.Contains(version.GitVersion, "gke") {
		return core.Pointer(core.KubernetesDistributionGKE)
	}

	// Check for EKS
	if strings.Contains(version.GitVersion, "eks") {
		return core.Pointer(core.KubernetesDistributionEKS)
	}

	// Check for k3s
	if strings.Contains(version.GitVersion, "k3s") {
		return core.Pointer(core.KubernetesDistributionK3S)
	}

	if strings.Contains(config.Host, ":8443") {
		// This is a common port for Minikube.
		return core.Pointer(core.KubernetesDistributionMinikube)
	}

	if strings.Contains(config.Host, ":16443") {
		// This is a common port for MicroK8s.
		return core.Pointer(core.KubernetesDistributionMicroK8s)
	}

	return nil
}

func (client *KubernetesClient) getApplySetParentRef(name string, namespace string) (*apply.ApplySetParentRef, error) {
	fqdnName := getApplySetParentName(name)

	configNamespace, _, _ := client.Factory.ToRawKubeConfigLoader().Namespace()
	if namespace == "" {
		namespace = configNamespace
	}

	parentRef, err := apply.ParseApplySetParentRef(fqdnName, client.Mapper)
	if err != nil {
		return nil, fmt.Errorf("failed to parse apply set parent ref: %w", err)
	}

	if parentRef.IsNamespaced() {
		parentRef.Namespace = namespace
	}

	return parentRef, nil
}

func (client *KubernetesClient) getApplySetRestClient(applySetParentRef *apply.ApplySetParentRef) (resource.RESTClient, error) {
	return client.getRestClient(applySetParentRef.RESTMapping)
}

func (client *KubernetesClient) getRestClient(mapping *meta.RESTMapping) (resource.RESTClient, error) {
	restClient, err := client.Factory.UnstructuredClientForMapping(mapping)
	if err != nil {
		return nil, fmt.Errorf("failed to get unstructured client: %w", err)
	}
	if restClient == nil {
		return nil, fmt.Errorf("failed to build unstructured client for ApplySet")
	}

	return restClient, nil
}

func getApplySetParentName(applySetParentName string) string {
	return fmt.Sprintf("%s.anemos.sh", applySetParentName)
}

func cleanupApplySetParentName(applySetParentName string) string {
	if strings.HasSuffix(applySetParentName, ".anemos.sh") {
		return strings.TrimSuffix(applySetParentName, ".anemos.sh")
	}
	return applySetParentName
}

func getTooling() apply.ApplySetTooling {
	return apply.ApplySetTooling{
		Name:    "anemos",
		Version: fmt.Sprintf("v%s", util.AppVersion),
	}
}

func getUID(object runtime.Object) *types.UID {
	o, err := meta.Accessor(object)
	if err != nil {
		return nil
	}

	uid := o.GetUID()

	return &uid
}

func omitManagedFields(object runtime.Object) runtime.Object {
	accessor, err := meta.Accessor(object)
	if err != nil {
		return object
	}

	accessor.SetManagedFields(nil)

	return object
}

func addExtraLabels(info *resource.Info, extraLabels map[string]string) {
	accessor, err := meta.Accessor(info.Object)
	if err != nil {
		return
	}

	labels := accessor.GetLabels()
	if labels == nil {
		labels = make(map[string]string)
	}

	for k, v := range extraLabels {
		labels[k] = v
	}

	accessor.SetLabels(labels)
}

func setGeneration(live, merged runtime.Object) {
	mergedAccessor, err := meta.Accessor(merged)
	if err != nil {
	}

	liveAccessor, err := meta.Accessor(live)
	if err != nil {
		return
	}

	liveAccessor.SetGeneration(mergedAccessor.GetGeneration())
}

func confirmChanges() (bool, error) {
	slog.Info("Apply these changes? [y/N]: ", util.SlogNoLineBreakAttr())

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return false, fmt.Errorf("failed to read confirmation: %w", err)
		}
		return false, fmt.Errorf("no input received")
	}

	response := strings.ToLower(strings.TrimSpace(scanner.Text()))
	return response == "y" || response == "yes", nil
}
