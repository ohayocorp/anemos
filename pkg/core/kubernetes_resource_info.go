package core

import (
	"fmt"
	"log/slog"
	"reflect"

	"github.com/Masterminds/semver/v3"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/ohayocorp/anemos/pkg/js"
)

// KubernetesResource defines a Kubernetes API resource. This includes both built-in Kubernetes objects
// and external CRDs.
type KubernetesResource struct {
	ApiVersion   string
	Kind         string
	IsNamespaced bool
}

// KubernetesResourceInfo contains all the API resources defined in the target cluster and enables listing them
// and querying their existence.
type KubernetesResourceInfo struct {
	resources      mapset.Set[*KubernetesResource]
	namespaceCache map[string]bool
}

func NewKubernetesResource(apiVersion string, kind string, isNamespaced bool) *KubernetesResource {
	return &KubernetesResource{
		ApiVersion:   apiVersion,
		Kind:         kind,
		IsNamespaced: isNamespaced,
	}
}

// Creates a new [KubernetesResourceInfo] instance.
func NewKubernetesResourceInfo(version *semver.Version) *KubernetesResourceInfo {
	info := KubernetesResourceInfo{
		resources:      mapset.NewSet[*KubernetesResource](),
		namespaceCache: make(map[string]bool),
	}

	info.populateBuiltInResources(version)

	return &info
}

// Returns true if the given API resource is namespaced. E.g. returns true for v1/Pod,
// false for rbac.authorization.k8s.io/v1/ClusterRole.
func (info *KubernetesResourceInfo) IsNamespaced(apiVersion string, kind string) bool {
	return info.namespaceCache[fmt.Sprintf("%s%s", apiVersion, kind)]
}

func (info *KubernetesResourceInfo) AllResources() []*KubernetesResource {
	return info.resources.ToSlice()
}

// Adds the given API resource to the available resources list.
func (info *KubernetesResourceInfo) AddResource(apiVersion string, kind string, isNamespaced bool) {
	resource := &KubernetesResource{
		ApiVersion:   apiVersion,
		Kind:         kind,
		IsNamespaced: isNamespaced,
	}

	info.AddKubernetesResource(resource)
}

// Adds the given API resource to the available resources list.
func (info *KubernetesResourceInfo) AddKubernetesResource(resource *KubernetesResource) {
	info.resources.Add(resource)
	info.namespaceCache[fmt.Sprintf("%s%s", resource.ApiVersion, resource.Kind)] = resource.IsNamespaced
}

// Returns true if the given API resource exists in the target cluster.
func (info *KubernetesResourceInfo) Contains(apiVersion string, kind string) bool {
	for resource := range info.resources.Iter() {
		if resource.ApiVersion == apiVersion && resource.Kind == kind {
			return true
		}
	}

	return false
}

// Returns true if the given kind exists in the target cluster. This ignores the apiVersion field.
func (info *KubernetesResourceInfo) ContainsKind(kind string) bool {
	for resource := range info.resources.Iter() {
		if resource.Kind == kind {
			return true
		}
	}

	return false
}

func (info *KubernetesResourceInfo) populateBuiltInResources(version *semver.Version) {
	// Run following command to create the lines for a version:
	// kubectl api-resources --no-headers | awk '{printf "info.AddResource(\"%s\", \"%s\", %s)\n", $(NF-2), $(NF), $(NF-1)}'

	kubernetesVersion := fmt.Sprintf("%d.%d", version.Major(), version.Minor())
	switch kubernetesVersion {
	case "1.26":
		info.addKubernetes1_26()
	case "1.27":
		info.addKubernetes1_27()
	case "1.28":
		info.addKubernetes1_28()
	case "1.29":
		info.addKubernetes1_29()
	case "1.30":
		info.addKubernetes1_30()
	case "1.31":
		info.addKubernetes1_31()
	case "1.32":
		info.addKubernetes1_32()
	case "1.33":
		info.addKubernetes1_33()
	default:
		if version.Major() == 1 && version.Minor() > 33 {
			slog.Warn(
				"Using Kubernetes version ${version} greater than 1.33, Anemos may not support all resources. Using 1.33 resources as a base.",
				slog.String("version", version.String()))

			info.addKubernetes1_33()
			return
		}

		js.Throw(fmt.Errorf("kubernetes version %s is not supported", kubernetesVersion))
	}
}

func (info *KubernetesResourceInfo) addKubernetes1_26() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1beta3", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1beta3", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}
func (info *KubernetesResourceInfo) addKubernetes1_27() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "SelfSubjectReview", false)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1beta3", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1beta3", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}

func (info *KubernetesResourceInfo) addKubernetes1_28() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "SelfSubjectReview", false)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1beta3", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1beta3", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}

func (info *KubernetesResourceInfo) addKubernetes1_29() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "SelfSubjectReview", false)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}

func (info *KubernetesResourceInfo) addKubernetes1_30() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicy", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicyBinding", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "SelfSubjectReview", false)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}

func (info *KubernetesResourceInfo) addKubernetes1_31() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicy", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicyBinding", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "SelfSubjectReview", false)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}

func (info *KubernetesResourceInfo) addKubernetes1_32() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicy", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicyBinding", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "SelfSubjectReview", false)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}

func (info *KubernetesResourceInfo) addKubernetes1_33() {
	info.AddResource("v1", "Binding", true)
	info.AddResource("v1", "ComponentStatus", false)
	info.AddResource("v1", "ConfigMap", true)
	info.AddResource("v1", "Endpoints", true)
	info.AddResource("v1", "Event", true)
	info.AddResource("v1", "LimitRange", true)
	info.AddResource("v1", "Namespace", false)
	info.AddResource("v1", "Node", false)
	info.AddResource("v1", "PersistentVolumeClaim", true)
	info.AddResource("v1", "PersistentVolume", false)
	info.AddResource("v1", "Pod", true)
	info.AddResource("v1", "PodTemplate", true)
	info.AddResource("v1", "ReplicationController", true)
	info.AddResource("v1", "ResourceQuota", true)
	info.AddResource("v1", "Secret", true)
	info.AddResource("v1", "ServiceAccount", true)
	info.AddResource("v1", "Service", true)
	info.AddResource("admissionregistration.k8s.io/v1", "MutatingWebhookConfiguration", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicy", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingAdmissionPolicyBinding", false)
	info.AddResource("admissionregistration.k8s.io/v1", "ValidatingWebhookConfiguration", false)
	info.AddResource("apiextensions.k8s.io/v1", "CustomResourceDefinition", false)
	info.AddResource("apiregistration.k8s.io/v1", "APIService", false)
	info.AddResource("apps/v1", "ControllerRevision", true)
	info.AddResource("apps/v1", "DaemonSet", true)
	info.AddResource("apps/v1", "Deployment", true)
	info.AddResource("apps/v1", "ReplicaSet", true)
	info.AddResource("apps/v1", "StatefulSet", true)
	info.AddResource("authentication.k8s.io/v1", "SelfSubjectReview", false)
	info.AddResource("authentication.k8s.io/v1", "TokenReview", false)
	info.AddResource("authorization.k8s.io/v1", "LocalSubjectAccessReview", true)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectAccessReview", false)
	info.AddResource("authorization.k8s.io/v1", "SelfSubjectRulesReview", false)
	info.AddResource("authorization.k8s.io/v1", "SubjectAccessReview", false)
	info.AddResource("autoscaling/v2", "HorizontalPodAutoscaler", true)
	info.AddResource("batch/v1", "CronJob", true)
	info.AddResource("batch/v1", "Job", true)
	info.AddResource("certificates.k8s.io/v1", "CertificateSigningRequest", false)
	info.AddResource("coordination.k8s.io/v1", "Lease", true)
	info.AddResource("discovery.k8s.io/v1", "EndpointSlice", true)
	info.AddResource("events.k8s.io/v1", "Event", true)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "FlowSchema", false)
	info.AddResource("flowcontrol.apiserver.k8s.io/v1", "PriorityLevelConfiguration", false)
	info.AddResource("networking.k8s.io/v1", "IngressClass", false)
	info.AddResource("networking.k8s.io/v1", "Ingress", true)
	info.AddResource("networking.k8s.io/v1", "IPAddress", false)
	info.AddResource("networking.k8s.io/v1", "NetworkPolicy", true)
	info.AddResource("networking.k8s.io/v1", "ServiceCIDR", false)
	info.AddResource("node.k8s.io/v1", "RuntimeClass", false)
	info.AddResource("policy/v1", "PodDisruptionBudget", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRoleBinding", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "ClusterRole", false)
	info.AddResource("rbac.authorization.k8s.io/v1", "RoleBinding", true)
	info.AddResource("rbac.authorization.k8s.io/v1", "Role", true)
	info.AddResource("scheduling.k8s.io/v1", "PriorityClass", false)
	info.AddResource("storage.k8s.io/v1", "CSIDriver", false)
	info.AddResource("storage.k8s.io/v1", "CSINode", false)
	info.AddResource("storage.k8s.io/v1", "CSIStorageCapacity", true)
	info.AddResource("storage.k8s.io/v1", "StorageClass", false)
	info.AddResource("storage.k8s.io/v1", "VolumeAttachment", false)
}

func registerKubernetesResourceInfo(jsRuntime *js.JsRuntime) {
	jsRuntime.Type(reflect.TypeFor[KubernetesResource]()).Fields(
		js.Field("ApiVersion"),
		js.Field("Kind"),
		js.Field("IsNamespaced"),
	).Constructors(
		js.Constructor(reflect.ValueOf(NewKubernetesResource)),
	)

	jsRuntime.Type(reflect.TypeFor[KubernetesResourceInfo]()).Fields().Methods(
		js.Method("AddKubernetesResource").JsName("addResource"),
		js.Method("Contains"),
		js.Method("ContainsKind"),
		js.Method("IsNamespaced"),
	)
}
