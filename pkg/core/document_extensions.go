package core

import (
	"reflect"

	"github.com/ohayocorp/anemos/pkg/js"
)

var (
	MetadataLabels LabelNode = func(document *Document) *Mapping {
		return document.EnsureLabels()
	}
	WorkloadLabels LabelNode = func(document *Document) *Mapping {
		if document.IsWorkload() {
			return document.EnsureWorkloadLabels()
		}

		return nil
	}
	WorkloadSelector LabelNode = func(document *Document) *Mapping {
		if document.IsWorkload() && !document.IsPod() {
			return document.GetRoot().EnsureMappingChain("spec", "selector", "matchLabels")
		}

		return nil
	}
	ServiceSelector LabelNode = func(document *Document) *Mapping {
		if document.IsService() {
			return document.GetRoot().EnsureMappingChain("spec", "selector")
		}

		return nil
	}
	ServiceMonitorSelector LabelNode = func(document *Document) *Mapping {
		if document.IsOfKind("monitoring.coreos.com/v1", "ServiceMonitor") {
			return document.GetRoot().EnsureMappingChain("spec", "selector", "matchLabels")
		}

		return nil
	}
)

type LabelNode func(document *Document) *Mapping

// Returns value of ".apiVersion". Returns nil if the value is not found.
func (document *Document) GetApiVersion() *string {
	return document.GetRoot().GetValue("apiVersion")
}

// Returns value of ".kind". Returns nil if the value is not found.
func (document *Document) GetKind() *string {
	return document.GetRoot().GetValue("kind")
}

// Returns value of ".metadata.name". Returns nil if the value is not found.
func (document *Document) GetName() *string {
	return document.GetRoot().GetValueChain("metadata", "name")
}

// Returns value of ".metadata.namespace". Returns nil if the value is not found.
func (document *Document) GetNamespace() *string {
	return document.GetRoot().GetValueChain("metadata", "namespace")
}

// Returns the [Mapping] for ".metadata". Returns nil if the [Mapping] is not found.
func (document *Document) GetMetadata() *Mapping {
	return document.GetRoot().GetMapping("metadata")
}

// Returns the [Mapping] for ".metadata.labels". Returns nil if the [Mapping] is not found.
func (document *Document) GetLabels() *Mapping {
	return document.GetRoot().GetMappingChain("metadata", "labels")
}

// Returns the value of ".metadata.labels.$label". Returns nil if the value is not found.
func (document *Document) GetLabel(label string) *string {
	return document.GetRoot().GetValueChain("metadata", "labels", label)
}

// Returns the [Mapping] for ".metadata.annotations". Returns nil if the [Mapping] is not found.
func (document *Document) GetAnnotations() *Mapping {
	return document.GetRoot().GetMappingChain("metadata", "annotations")
}

// Returns the value of ".metadata.annotations.$annotation". Returns nil if the value is not found.
func (document *Document) GetAnnotation(annotation string) *string {
	return document.GetRoot().GetValueChain("metadata", "annotations", annotation)
}

// Returns the [Mapping] for ".spec". Returns nil if the [Mapping] is not found.
func (document *Document) GetSpec() *Mapping {
	return document.GetRoot().GetMapping("spec")
}

// Returns the [Mapping] for ".spec" if the document specifies a Pod or ".spec.template.spec" if the document
// specifies another type of workload. Returns nil if the [Mapping] is not found.
func (document *Document) GetWorkloadSpec() *Mapping {
	if document.IsPod() {
		return document.GetRoot().GetMapping("spec")
	}

	return document.GetRoot().GetMappingChain("spec", "template", "spec")
}

// Returns the [Mapping] for ".metadata" if the document specifies a Pod or ".spec.template.metadata" if the document
// specifies another type of workload. Returns nil if the [Mapping] is not found.
func (document *Document) GetWorkloadMetadata() *Mapping {
	if document.IsPod() {
		return document.GetRoot().GetMapping("metadata")
	}

	return document.GetRoot().GetMappingChain("spec", "template", "metadata")
}

// Returns the [Mapping] for ".metadata.labels" if the document specifies a Pod or ".spec.template.metadata.labels" if the document
// specifies another type of workload. Returns nil if the [Mapping] is not found.
func (document *Document) GetWorkloadLabels() *Mapping {
	if document.IsPod() {
		return document.GetRoot().GetMappingChain("metadata", "labels")
	}

	return document.GetRoot().GetMappingChain("spec", "template", "metadata", "labels")
}

// Returns the [Mapping] for ".metadata.annotations" if the document specifies a Pod or ".spec.template.metadata.annotations" if the document
// specifies another type of workload. Returns nil if the [Mapping] is not found.
func (document *Document) GetWorkloadAnnotations() *Mapping {
	if document.IsPod() {
		return document.GetRoot().GetMappingChain("metadata", "annotations")
	}

	return document.GetRoot().GetMappingChain("spec", "template", "metadata", "annotations")
}

// Returns the [Sequence] for ".spec.volumes" if the document specifies a Pod or ".spec.template.spec.volumes" if the document
// specifies another type of workload. Returns nil if the [Sequence] is not found.
func (document *Document) GetVolumes() *Sequence {
	spec := document.GetWorkloadSpec()
	if spec == nil {
		return nil
	}

	return spec.GetSequence("volumes")
}

// Returns the [Mapping] for ".spec.volumes[i]" if the document specifies a Pod or ".spec.template.spec.volumes[i]" if the document
// specifies another type of workload. Returns nil if the [Mapping] is not found.
func (document *Document) GetVolume(i int) *Mapping {
	volumes := document.GetVolumes()
	if volumes == nil {
		return nil
	}

	return volumes.GetMapping(i)
}

// Returns the first [Mapping] under ".spec.volumes" if the document specifies a Pod or ".spec.template.spec.volumes" if the document
// specifies another type of workload that the name equals to the given value. Returns nil if the [Mapping] is not found.
func (document *Document) GetVolumeWithName(name string) *Mapping {
	volume := document.GetVolumeFunc(func(m *Mapping) bool {
		n := m.GetValue("name")
		if n == nil {
			return false
		}

		return *n == name
	})

	return volume
}

// Returns the first [Mapping] under ".spec.volumes" if the document specifies a Pod or ".spec.template.spec.volumes" if the document
// specifies another type of workload that the filter function returns true for. Returns nil if the [Mapping] is not found.
func (document *Document) GetVolumeFunc(filter func(*Mapping) bool) *Mapping {
	volumes := document.GetVolumes()
	if volumes == nil {
		return nil
	}

	for i := 0; i < volumes.Length(); i++ {
		volume := volumes.GetMapping(i)
		ok := filter(volume)

		if ok {
			return volume
		}
	}

	return nil
}

// Returns the [Sequence] for ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers" if the document
// specifies another type of workload. Returns nil if the [Sequence] is not found.
func (document *Document) GetContainers() *Sequence {
	spec := document.GetWorkloadSpec()
	if spec == nil {
		return nil
	}

	return spec.GetSequence("containers")
}

// Returns the [Sequence] for ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers" if the document
// specifies another type of workload. Returns nil if the [Sequence] is not found.
func (document *Document) GetInitContainers() *Sequence {
	spec := document.GetWorkloadSpec()
	if spec == nil {
		return nil
	}

	return spec.GetSequence("initContainers")
}

// Returns the [Mapping] for ".spec.containers[i]" if the document specifies a Pod or ".spec.template.spec.containers[i]" if the document
// specifies another type of workload. Returns nil if the [Mapping] is not found.
func (document *Document) GetContainer(i int) *Mapping {
	containers := document.GetContainers()
	if containers == nil {
		return nil
	}

	return containers.GetMapping(i)
}

// Returns the first [Mapping] under ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers"
// if the document specifies another type of workload that the filter function returns true for.
// Returns nil if the [Mapping] is not found.
func (document *Document) GetContainerFunc(filter func(*Mapping) bool) *Mapping {
	containers := document.GetContainers()
	if containers == nil {
		return nil
	}

	for i := 0; i < containers.Length(); i++ {
		container := containers.GetMapping(i)
		ok := filter(container)
		if ok {
			return container
		}
	}

	return nil
}

// Returns the first [Mapping] under ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers"
// if the document specifies another type of workload that the name equals to the given parameter.
// Returns nil if the [Mapping] is not found.
func (document *Document) GetContainerWithName(name string) *Mapping {
	container := document.GetContainerFunc(func(m *Mapping) bool {
		n := m.GetValue("name")
		return n != nil && *n == name
	})

	return container
}

// Returns the [Mapping] for ".spec.initContainers[i]" if the document specifies a Pod or ".spec.template.spec.initContainers[i]"
// if the document specifies another type of workload. Returns nil if the [Mapping] is not found.
func (document *Document) GetInitContainer(i int) *Mapping {
	initContainers := document.GetInitContainers()
	if initContainers == nil {
		return nil
	}

	return initContainers.GetMapping(i)
}

// Returns the first [Mapping] under ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers"
// if the document specifies another type of workload that the name equals to the given parameter.
// Returns nil if the [Mapping] is not found.
func (document *Document) GetInitContainerWithName(name string) *Mapping {
	initContainer := document.GetInitContainerFunc(func(m *Mapping) bool {
		n := m.GetValue("name")
		return n != nil && *n == name
	})

	return initContainer
}

// Returns the first [Mapping] under ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers"
// if the document specifies another type of workload that the filter function returns true for.
// Returns nil if the [Mapping] is not found.
func (document *Document) GetInitContainerFunc(filter func(*Mapping) bool) *Mapping {
	initContainers := document.GetInitContainers()
	if initContainers == nil {
		return nil
	}

	for i := 0; i < initContainers.Length(); i++ {
		initContainer := initContainers.GetMapping(i)
		ok := filter(initContainer)
		if ok {
			return initContainer
		}
	}

	return nil
}

// Ensures [Mapping] for ".metadata".
func (document *Document) EnsureMetadata() *Mapping {
	return document.GetRoot().EnsureMapping("metadata")
}

// Ensures [Mapping] for ".metadata.annotations".
func (document *Document) EnsureAnnotations() *Mapping {
	return document.EnsureMetadata().EnsureMapping("annotations")
}

// Ensures [Mapping] for ".metadata.labels".
func (document *Document) EnsureLabels() *Mapping {
	return document.EnsureMetadata().EnsureMapping("labels")
}

// Ensures [Mapping] for ".spec" if the document specifies a Pod or ".spec.template.spec" if the document
// specifies another type of workload.
func (document *Document) EnsureWorkloadSpec() *Mapping {
	if document.IsPod() {
		return document.GetRoot().EnsureMapping("spec")
	}

	return document.GetRoot().EnsureMappingChain("spec", "template", "spec")
}

// Ensures [Mapping] for ".metadata" if the document specifies a Pod or ".spec.template.metadata" if the document
// specifies another type of workload. Inserts the mapping at the beginning for better readability.
func (document *Document) EnsureWorkloadMetadata() *Mapping {
	mapping := document.GetRoot()

	if !document.IsPod() {
		mapping = document.GetRoot().EnsureMappingChain("spec", "template")
	}

	if mapping.ContainsKey("metadata") {
		return mapping.GetMapping("metadata")
	}

	metadata := NewEmptyMapping()
	mapping.InsertMapping(0, "metadata", metadata)

	return metadata
}

// Ensures [Mapping] for ".metadata.annotations" if the document specifies a Pod or ".spec.template.metadata.annotations" if the document
// specifies another type of workload.
func (document *Document) EnsureWorkloadAnnotations() *Mapping {
	return document.EnsureWorkloadMetadata().EnsureMapping("annotations")
}

// Ensures [Mapping] for ".metadata.labels" if the document specifies a Pod or ".spec.template.metadata.labels" if the document
// specifies another type of workload.
func (document *Document) EnsureWorkloadLabels() *Mapping {
	return document.EnsureWorkloadMetadata().EnsureMapping("labels")
}

// Ensures [Sequence] for ".spec.volumes" if the document specifies a Pod or ".spec.template.spec.volumes" if the document
// specifies another type of workload.
func (document *Document) EnsureVolumes() *Sequence {
	return document.EnsureWorkloadSpec().EnsureSequence("volumes")
}

// Ensures [Sequence] for ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers" if the document
// specifies another type of workload.
func (document *Document) EnsureContainers() *Sequence {
	return document.EnsureWorkloadSpec().EnsureSequence("containers")
}

// Ensures [Sequence] for ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers" if the document
// specifies another type of workload.
func (document *Document) EnsureInitContainers() *Sequence {
	return document.EnsureWorkloadSpec().EnsureSequence("initContainers")
}

// Returns true if the document has the given apiVersion and kind.
func (document *Document) IsOfKind(apiVersion string, kind string) bool {
	documentApiVersion := document.GetApiVersion()
	if documentApiVersion == nil {
		return false
	}

	documentKind := document.GetKind()
	if documentKind == nil {
		return false
	}

	return *documentApiVersion == apiVersion && *documentKind == kind
}

// Returns true if the document is a ClusterRole.
func (document *Document) IsClusterRole() bool {
	return document.IsOfKind("rbac.authorization.k8s.io/v1", "ClusterRole")
}

// Returns true if the document is a ClusterRoleBinding.
func (document *Document) IsClusterRoleBinding() bool {
	return document.IsOfKind("rbac.authorization.k8s.io/v1", "ClusterRoleBinding")
}

// Returns true if the document is a ConfigMap.
func (document *Document) IsConfigMap() bool {
	return document.IsOfKind("v1", "ConfigMap")
}

// Returns true if the document is a CronJob.
func (document *Document) IsCronJob() bool {
	return document.IsOfKind("batch/v1", "CronJob")
}

// Returns true if the document is a CustomResourceDefinition.
func (document *Document) IsCustomResourceDefinition() bool {
	return document.IsOfKind("apiextensions.k8s.io/v1", "CustomResourceDefinition")
}

// Returns true if the document is a DaemonSet.
func (document *Document) IsDaemonSet() bool {
	return document.IsOfKind("apps/v1", "DaemonSet")
}

// Returns true if the document is a Deployment.
func (document *Document) IsDeployment() bool {
	return document.IsOfKind("apps/v1", "Deployment")
}

// Returns true if the document is a HorizontalPodAutoscaler.
func (document *Document) IsHorizontalPodAutoscaler() bool {
	return document.IsOfKind("autoscaling/v2", "HorizontalPodAutoscaler")
}

// Returns true if the document is an Ingress.
func (document *Document) IsIngress() bool {
	return document.IsOfKind("networking.k8s.io/v1", "Ingress")
}

// Returns true if the document is a Job.
func (document *Document) IsJob() bool {
	return document.IsOfKind("batch/v1", "Job")
}

// Returns true if the document is a Namespace.
func (document *Document) IsNamespace() bool {
	return document.IsOfKind("v1", "Namespace")
}

// Returns true if the document is a PersistentVolume.
func (document *Document) IsPersistentVolume() bool {
	return document.IsOfKind("v1", "PersistentVolume")
}

// Returns true if the document is a PersistentVolumeClaim.
func (document *Document) IsPersistentVolumeClaim() bool {
	return document.IsOfKind("v1", "PersistentVolumeClaim")
}

// Returns true if the document is a Pod.
func (document *Document) IsPod() bool {
	return document.IsOfKind("v1", "Pod")
}

// Returns true if the document is a ReplicaSet.
func (document *Document) IsReplicaSet() bool {
	return document.IsOfKind("apps/v1", "ReplicaSet")
}

// Returns true if the document is a Role.
func (document *Document) IsRole() bool {
	return document.IsOfKind("rbac.authorization.k8s.io/v1", "Role")
}

// Returns true if the document is a RoleBinding.
func (document *Document) IsRoleBinding() bool {
	return document.IsOfKind("rbac.authorization.k8s.io/v1", "RoleBinding")
}

// Returns true if the document is a Secret.
func (document *Document) IsSecret() bool {
	return document.IsOfKind("v1", "Secret")
}

// Returns true if the document is a Service.
func (document *Document) IsService() bool {
	return document.IsOfKind("v1", "Service")
}

// Returns true if the document is a ServiceAccount.
func (document *Document) IsServiceAccount() bool {
	return document.IsOfKind("v1", "ServiceAccount")
}

// Returns true if the document is a StatefulSet.
func (document *Document) IsStatefulSet() bool {
	return document.IsOfKind("apps/v1", "StatefulSet")
}

// Returns true if the document is one of these: CronJob, DaemonSet, Deployment, Job, Pod, ReplicaSet, StatefulSet.
func (document *Document) IsWorkload() bool {
	return false ||
		document.IsCronJob() ||
		document.IsDaemonSet() ||
		document.IsDeployment() ||
		document.IsJob() ||
		document.IsPod() ||
		document.IsReplicaSet() ||
		document.IsStatefulSet()
}

// Adds the given label and value to ".metadata.labels".
func (document *Document) SetLabel(label string, value string) {
	document.EnsureLabels().SetValue(label, value)
}

// Adds the given labels and values to ".metadata.labels". Keys are sorted alphabetically.
func (document *Document) SetLabels(labels map[string]string) {
	for _, key := range SortedKeys(labels) {
		value := labels[key]
		document.SetLabel(key, value)
	}
}

// Adds the given labels and values to the specified nodes. Keys are sorted alphabetically.
func (document *Document) SetLabelsWithNodes(labels map[string]string, labelNodes []LabelNode) {
	for _, key := range SortedKeys(labels) {
		value := labels[key]

		for _, labelNode := range labelNodes {
			node := labelNode(document)
			if node != nil {
				node.SetValue(key, value)
			}
		}
	}
}

// Sets the given value to ".metadata.annotations.$key".
func (document *Document) SetAnnotation(key string, value string) {
	document.EnsureAnnotations().SetValue(key, value)
}

// Adds the given labels and values to ".metadata.annotations". Keys are sorted alphabetically.
func (document *Document) SetAnnotations(annotations map[string]string) {
	for _, key := range SortedKeys(annotations) {
		value := annotations[key]
		document.SetAnnotation(key, value)
	}
}

// Sets the given value to ".metadata.name".
func (document *Document) SetName(value string) {
	document.EnsureMetadata().SetValue("name", value)
}

// Sets the given value to ".metadata.namespace".
func (document *Document) SetNamespace(namespace string) {
	metadata := document.EnsureMetadata()

	namespaceIndex := metadata.IndexOfKey("namespace")
	if namespaceIndex >= 0 {
		metadata.SetValue("namespace", namespace)
		return
	}

	// Place namespace under name if possible
	nameIndex := metadata.IndexOfKey("name")
	if nameIndex < 0 {
		// Just append to the end
		metadata.SetValue("namespace", namespace)
		return
	}

	metadata.InsertValue(nameIndex+1, "namespace", namespace)
}

func registerYamlDocumentExtensions(jsRuntime *js.JsRuntime) {
	jsRuntime.Function(reflect.ValueOf(MetadataLabels)).JsNamespace("LabelNodes").JsName("metadataLabels")
	jsRuntime.Function(reflect.ValueOf(WorkloadLabels)).JsNamespace("LabelNodes").JsName("workloadLabels")
	jsRuntime.Function(reflect.ValueOf(WorkloadSelector)).JsNamespace("LabelNodes").JsName("workloadSelector")
	jsRuntime.Function(reflect.ValueOf(ServiceSelector)).JsNamespace("LabelNodes").JsName("serviceSelector")
	jsRuntime.Function(reflect.ValueOf(ServiceMonitorSelector)).JsNamespace("LabelNodes").JsName("serviceMonitorSelector")

	jsRuntime.Type(reflect.TypeFor[Document]()).Methods(
		js.Method("EnsureAnnotations"),
		js.Method("EnsureContainers"),
		js.Method("EnsureInitContainers"),
		js.Method("EnsureLabels"),
		js.Method("EnsureMetadata"),
		js.Method("EnsureVolumes"),
		js.Method("EnsureWorkloadAnnotations"),
		js.Method("EnsureWorkloadLabels"),
		js.Method("EnsureWorkloadMetadata"),
		js.Method("EnsureWorkloadSpec"),
		js.Method("GetAnnotation"),
		js.Method("GetAnnotations"),
		js.Method("GetApiVersion"),
		js.Method("GetContainer"),
		js.Method("GetContainerWithName").JsName("getContainer"),
		js.Method("GetContainerFunc").JsName("getContainer"),
		js.Method("GetContainers"),
		js.Method("GetInitContainer"),
		js.Method("GetInitContainerWithName").JsName("getInitContainer"),
		js.Method("GetInitContainerFunc").JsName("getInitContainer"),
		js.Method("GetInitContainers"),
		js.Method("GetKind"),
		js.Method("GetLabel"),
		js.Method("GetLabels"),
		js.Method("GetMetadata"),
		js.Method("GetName"),
		js.Method("GetNamespace"),
		js.Method("GetSpec"),
		js.Method("GetVolume"),
		js.Method("GetVolumeFunc").JsName("getVolume"),
		js.Method("GetVolumes"),
		js.Method("GetWorkloadAnnotations"),
		js.Method("GetWorkloadLabels"),
		js.Method("GetWorkloadMetadata"),
		js.Method("GetWorkloadSpec"),
		js.Method("IsClusterRole"),
		js.Method("IsClusterRoleBinding"),
		js.Method("IsConfigMap"),
		js.Method("IsCronJob"),
		js.Method("IsCustomResourceDefinition"),
		js.Method("IsDaemonSet"),
		js.Method("IsDeployment"),
		js.Method("IsHorizontalPodAutoscaler"),
		js.Method("IsIngress"),
		js.Method("IsJob"),
		js.Method("IsNamespace"),
		js.Method("IsOfKind"),
		js.Method("IsPersistentVolume"),
		js.Method("IsPersistentVolumeClaim"),
		js.Method("IsPod"),
		js.Method("IsReplicaSet"),
		js.Method("IsRole"),
		js.Method("IsRoleBinding"),
		js.Method("IsSecret"),
		js.Method("IsService"),
		js.Method("IsServiceAccount"),
		js.Method("IsStatefulSet"),
		js.Method("IsWorkload"),
		js.Method("SetAnnotation"),
		js.Method("SetAnnotations"),
		js.Method("SetLabel"),
		js.Method("SetLabels"),
		js.Method("SetLabelsWithNodes").JsName("setLabels"),
		js.Method("SetName"),
		js.Method("SetNamespace"),
	)
}
