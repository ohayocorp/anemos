// Auto generated code; DO NOT EDIT.
import { CSIVolumeSource } from "./CSIVolumeSource"
import { ConfigMapVolumeSource } from "./ConfigMapVolumeSource"
import { DownwardAPIVolumeSource } from "./DownwardAPIVolumeSource"
import { EmptyDirVolumeSource } from "./EmptyDirVolumeSource"
import { EphemeralVolumeSource } from "./EphemeralVolumeSource"
import { FCVolumeSource } from "./FCVolumeSource"
import { HostPathVolumeSource } from "./HostPathVolumeSource"
import { ISCSIVolumeSource } from "./ISCSIVolumeSource"
import { ImageVolumeSource } from "./ImageVolumeSource"
import { NFSVolumeSource } from "./NFSVolumeSource"
import { PersistentVolumeClaimVolumeSource } from "./PersistentVolumeClaimVolumeSource"
import { ProjectedVolumeSource } from "./ProjectedVolumeSource"
import { SecretVolumeSource } from "./SecretVolumeSource"

/**
 * Volume represents a named volume in a pod that may be accessed by any container in the pod.
 */
export declare class Volume {
    constructor();
    constructor(spec: Pick<Volume, "configMap" | "csi" | "downwardAPI" | "emptyDir" | "ephemeral" | "fc" | "hostPath" | "image" | "iscsi" | "name" | "nfs" | "persistentVolumeClaim" | "projected" | "secret">);

	/**
     * ConfigMap represents a configMap that should populate this volume
     */
    configMap?: ConfigMapVolumeSource

	/**
     * Csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers.
     */
    csi?: CSIVolumeSource

	/**
     * DownwardAPI represents downward API about the pod that should populate this volume
     */
    downwardAPI?: DownwardAPIVolumeSource

	/**
     * EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir
     */
    emptyDir?: EmptyDirVolumeSource

	/**
     * Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.
    
     * Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity
    
     *    tracking are needed,
    
     * c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through
    
     *    a PersistentVolumeClaim (see EphemeralVolumeSource for more
    
     *    information on the connection between this volume type
    
     *    and PersistentVolumeClaim).
    
     * Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.
    
     * Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.
    
     * A pod can use both types of ephemeral volumes and persistent volumes at the same time.
     */
    ephemeral?: EphemeralVolumeSource

	/**
     * Fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
     */
    fc?: FCVolumeSource

	/**
     * HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     */
    hostPath?: HostPathVolumeSource

	/**
     * Image represents an OCI object (a container image or artifact) pulled and mounted on the kubelet's host machine. The volume is resolved at pod startup depending on which PullPolicy value is provided:
    
     * - Always: the kubelet always attempts to pull the reference. Container creation will fail If the pull fails. - Never: the kubelet never pulls the reference and only uses a local image or artifact. Container creation will fail if the reference isn't present. - IfNotPresent: the kubelet pulls if the reference isn't already present on disk. Container creation will fail if the reference isn't present and the pull fails.
    
     * The volume gets re-resolved if the pod gets deleted and recreated, which means that new remote content will become available on pod recreation. A failure to resolve or pull the image during pod startup will block containers from starting and may add significant latency. Failures will be retried using normal volume backoff and will be reported on the pod reason and message. The types of objects that may be mounted by this volume are defined by the container runtime implementation on a host machine and at minimum must include all valid types supported by the container image field. The OCI object gets mounted in a single directory (spec.containers[*].volumeMounts.mountPath) by merging the manifest layers in the same way as for container images. The volume will be mounted read-only (ro) and non-executable files (noexec). Sub path mounts for containers are not supported (spec.containers[*].volumeMounts.subpath) before 1.33. The field spec.securityContext.fsGroupChangePolicy has no effect on this volume type.
     */
    image?: ImageVolumeSource

	/**
     * Iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes/#iscsi
     */
    iscsi?: ISCSIVolumeSource

	/**
     * Name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name: string

	/**
     * Nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     */
    nfs?: NFSVolumeSource

	/**
     * PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
     */
    persistentVolumeClaim?: PersistentVolumeClaimVolumeSource

	/**
     * Projected items for all in one resources secrets, configmaps, and downward API
     */
    projected?: ProjectedVolumeSource

	/**
     * Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret
     */
    secret?: SecretVolumeSource

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}