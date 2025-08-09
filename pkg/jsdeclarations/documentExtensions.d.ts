import { Mapping } from "./mapping"
import { Sequence } from "./sequence"
import { Document } from "./document"

declare module "./document" {
    export interface Document {
        /** Ensures a {@link Mapping} for ".metadata". */
        ensureMetadata(): Mapping;

        /** Ensures a {@link Mapping} for ".metadata.labels". */
        ensureLabels(): Mapping;

        /** Ensures a {@link Mapping} for ".metadata.annotations". */
        ensureAnnotations(): Mapping;

        /**
         * Ensures a {@link Mapping} for ".spec" if the document specifies a Pod or ".spec.template.spec" if the document
         * specifies another type of workload.
         */
        ensureWorkloadSpec(): Mapping;

        /**
         * Ensures a {@link Sequence} for ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers" if the document
         * specifies another type of workload.
         */
        ensureContainers(): Sequence;

        /**
         * Ensures a {@link Sequence} for ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers" if the document
         * specifies another type of workload.
         */
        ensureInitContainers(): Sequence;

        /** 
         * Ensures a {@link Sequence} for ".spec.volumes" if the document specifies a Pod or ".spec.template.spec.volumes" if the document
         * specifies another type of workload.
         */
        ensureVolumes(): Sequence;

        /**
         * Ensures a {@link Mapping} for ".metadata" if the document specifies a Pod or ".spec.template.metadata" if the document
         * specifies another type of workload.
         */
        ensureWorkloadMetadata(): Mapping;

        /**
         * Ensures a {@link Mapping} for ".metadata.labels" if the document specifies a Pod or ".spec.template.metadata.labels" if the document
         * specifies another type of workload.
         */
        ensureWorkloadLabels(): Mapping;

        /**
         * Ensures a {@link Mapping} for ".metadata.annotations" if the document specifies a Pod or ".spec.template.metadata.annotations" if the document
         * specifies another type of workload.
         */
        ensureWorkloadAnnotations(): Mapping;

        /** Returns the value of ".apiVersion". Returns null if the value is not found. */
        getApiVersion(): string | null;

        /** Returns value of ".kind". Returns null if the value is not found. */
        getKind(): string | null;

        /** Returns the {@link Mapping} for ".metadata". Returns null if the {@link Mapping} is not found. */
        getMetadata(): Mapping | null;

        /** Returns value of ".metadata.name". Returns null if the value is not found. */
        getName(): string | null;

        /** Returns value of ".metadata.namespace". Returns null if the value is not found. */
        getNamespace(): string | null;

        /** Returns the value of ".metadata.labels.$label". Returns null if the value is not found. */
        getLabel(label: string): string | null;

        /** Returns the {@link Mapping} for ".metadata.labels". Returns null if the {@link Mapping} is not found. */
        getLabels(): Mapping | null;

        /** Returns the value of ".metadata.annotations.$annotation". Returns null if the value is not found. */
        getAnnotation(annotation: string): string | null;

        /** Returns the {@link Mapping} for ".metadata.annotations". Returns null if the {@link Mapping} is not found. */
        getAnnotations(): Mapping | null;

        /** Returns the {@link Mapping} for ".spec". Returns null if the value is not found. */
        getSpec(): Mapping | null;

        /**
         * Returns the {@link Mapping} for ".spec" if the document specifies a Pod or ".spec.template.spec" if the document
         * specifies another type of workload. Returns null if the {@link Mapping} is not found.
         */
        getWorkloadSpec(): Mapping | null;

        /** 
         * Returns the {@link ContainerMapping} for ".spec.containers[i]" if the document specifies a Pod or ".spec.template.spec.containers[i]" if the document
         * specifies another type of workload. Returns null if the {@link Mapping} is not found.
         */
        getContainer(index: number): Mapping | null;

        /**
         * Returns the first {@link Mapping} under ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers" if the document
         * specifies another type of workload that the name equals to the given parameter. Returns null if the {@link Mapping} is not found.
         */
        getContainer(name: string): Mapping | null;

        /**
         * Returns the first {@link Mapping} under ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers" if the document
         * specifies another type of workload that the filter function returns true for. Returns null if the {@link Mapping} is not found.
         */
        getContainer(filter: (container: Mapping) => boolean): Mapping | null;

        /**
         * Returns the {@link Sequence} for ".spec.containers" if the document specifies a Pod or ".spec.template.spec.containers" if the document
         * specifies another type of workload. Returns null if the {@link Sequence} is not found.
         */
        getContainers(): Sequence | null;

        /** 
         * Returns the {@link ContainerMapping} for ".spec.initContainers[i]" if the document specifies a Pod or ".spec.template.spec.initContainers[i]" if the document
         * specifies another type of workload. Returns null if the {@link Mapping} is not found.
         */
        getInitContainer(index: number): Mapping | null;

        /**
         * Returns the first {@link Mapping} under ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers" if the document
         * specifies another type of workload that the name equals to the given parameter. Returns null if the {@link Mapping} is not found.
         */
        getInitContainer(name: string): Mapping | null;

        /**
         * Returns the first {@link Mapping} under ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers" if the document
         * specifies another type of workload that the filter function returns true for. Returns null if the {@link Mapping} is not found.
         */
        getInitContainer(filter: (container: Mapping) => boolean): Mapping | null;

        /**
         * Returns the {@link Sequence} for ".spec.initContainers" if the document specifies a Pod or ".spec.template.spec.initContainers" if the document
         * specifies another type of workload. Returns null if the {@link Sequence} is not found.
         */
        getInitContainers(): Sequence | null;

        /** 
         * Returns the {@link Mapping} for ".spec.volumes[i]" if the document specifies a Pod or ".spec.template.spec.volumes[i]"
         * if the document specifies another type of workload. Returns null if the {@link Mapping} is not found.
         */
        getVolume(index: number): Mapping | null;

        /**
         * Returns the first {@link Mapping} under ".spec.volumes" if the document specifies a Pod or ".spec.template.spec.volumes"
         * if the document specifies another type of workload that the filter function returns true for.
         * Returns null if the {@link Mapping} is not found.
         */
        getVolume(filter: (volume: Mapping) => boolean): Mapping | null;

        /**
         * Returns the {@link Sequence} for ".spec.volumes" if the document specifies a Pod or ".spec.template.spec.volumes"
         * if the document specifies another type of workload. Returns null if the {@link Sequence} is not found.
         */
        getVolumes(): Sequence | null;

        /**
         * Returns the {@link Mapping} for ".metadata" if the document specifies a Pod or ".spec.template.metadata" if the document
         * specifies another type of workload. Returns null if the {@link Mapping} is not found.
         */
        getWorkloadMetadata(): Mapping | null;

        /**
         * Returns the {@link Mapping} for ".metadata.labels" if the document specifies a Pod or ".spec.template.metadata.labels"
         * if the document specifies another type of workload. Returns null if the {@link Mapping} is not found.
         */
        getWorkloadLabels(): Mapping | null;

        /**
         * Returns the {@link Mapping} for ".metadata.annotations" if the document specifies a Pod or ".spec.template.metadata.annotations"
         * if the document specifies another type of workload. Returns null if the {@link Mapping} is not found.
         */
        getWorkloadAnnotations(): Mapping | null;

        /** Returns true if the document has the given apiVersion and kind. */
        isOfKind(apiVersion: string, kind: string): boolean;

        /** Returns true if the document is a ClusterRole. */
        isClusterRole(): boolean;

        /** Returns true if the document is a ClusterRoleBinding. */
        isClusterRoleBinding(): boolean;

        /** Returns true if the document is a ConfigMap. */
        isConfigMap(): boolean;

        /** Returns true if the document is a CronJob. */
        isCronJob(): boolean;

        /** Returns true if the document is a CustomResourceDefinition. */
        isCustomResourceDefinition(): boolean;

        /** Returns true if the document is a DaemonSet. */
        isDaemonSet(): boolean;

        /** Returns true if the document is a Deployment. */
        isDeployment(): boolean;

        /** Returns true if the document is a HorizontalPodAutoscaler. */
        isHorizontalPodAutoscaler(): boolean;

        /** Returns true if the document is an Ingress. */
        isIngress(): boolean;

        /** Returns true if the document is a Job. */
        isJob(): boolean;

        /** Returns true if the document is a Namespace. */
        isNamespace(): boolean;

        /** Returns true if the document is a PersistentVolume. */
        isPersistentVolume(): boolean;

        /** Returns true if the document is a PersistentVolumeClaim. */
        isPersistentVolumeClaim(): boolean;

        /** Returns true if the document is a Pod. */
        isPod(): boolean;

        /** Returns true if the document is a ReplicaSet. */
        isReplicaSet(): boolean;

        /** Returns true if the document is a Role. */
        isRole(): boolean;

        /** Returns true if the document is a RoleBinding. */
        isRoleBinding(): boolean;

        /** Returns true if the document is a Secret. */
        isSecret(): boolean;

        /** Returns true if the document is a Service. */
        isService(): boolean;

        /** Returns true if the document is a ServiceAccount. */
        isServiceAccount(): boolean;

        /** Returns true if the document is a StatefulSet. */
        isStatefulSet(): boolean;

        /** Returns true if the document is one of these: CronJob, DaemonSet, Deployment, Job, Pod, ReplicaSet, StatefulSet. */
        isWorkload(): boolean;

        /** Sets the given value to ".metadata.name". */
        setName(value: string): void;

        /** Sets the given value to ".metadata.namespace". */
        setNamespace(value: string): void;

        /** Sets the given key value pair to ".metadata.labels". */
        setLabel(key: string, value: string): void;

        /** Sets the given key value pairs to ".metadata.labels". Keys are sorted alphabetically. */
        setLabels(labels: Record<string, string>): void;

        /** Sets the given key value pairs to the specified nodes on the document. Keys are sorted alphabetically. */
        setLabels(labels: Record<string, string>, labelNodes: LabelNode[]): void;

        /** Sets the given key value pair to ".metadata.annotations". */
        setAnnotation(key: string, value: string): void;

        /** Sets the given key value pairs to ".metadata.annotations". Keys are sorted alphabetically. */
        setAnnotations(annotations: Record<string, string>): void;
    }
}

/**
 * LabelNode is a function that takes a Document and ensures a Mapping for labels in the document.
 * It returns null if the labels are not available for the document. This is used to retrieve the
 * labels from different parts of a Kubernetes document.
 */
export declare type LabelNode = (document: Document) => Mapping | null;

export declare class LabelNodes {
    /** .metadata.labels */
    static readonly metadataLabels: LabelNode;

    /** .metadata.labels for Pods, .spec.template.metadata.labels for other workloads. */
    static readonly workloadLabels: LabelNode;

    /** .spec.selector.matchLabels for workloads other than Pods. */
    static readonly workloadSelector: LabelNode;

    /** .spec.selector for Services. */
    static readonly serviceSelector: LabelNode;

    /** .spec.selector for monitoring.coreos.com/v1 ServiceMonitors. */
    static readonly serviceMonitorSelector: LabelNode;
}
