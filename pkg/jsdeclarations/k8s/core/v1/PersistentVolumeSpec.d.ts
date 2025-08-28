// Auto generated code; DO NOT EDIT.
import { CSIPersistentVolumeSource } from "./CSIPersistentVolumeSource"
import { FCVolumeSource } from "./FCVolumeSource"
import { HostPathVolumeSource } from "./HostPathVolumeSource"
import { ISCSIPersistentVolumeSource } from "./ISCSIPersistentVolumeSource"
import { LocalVolumeSource } from "./LocalVolumeSource"
import { NFSVolumeSource } from "./NFSVolumeSource"
import { ObjectReference } from "./ObjectReference"
import { VolumeNodeAffinity } from "./VolumeNodeAffinity"

/**
 * PersistentVolumeSpec is the specification of a persistent volume.
 */
export declare class PersistentVolumeSpec {
    constructor();
    constructor(spec: Pick<PersistentVolumeSpec, "accessModes" | "capacity" | "csi" | "fc" | "hostPath" | "iscsi" | "local" | "mountOptions" | "nfs" | "nodeAffinity" | "persistentVolumeReclaimPolicy" | "storageClassName" | "volumeAttributesClassName" | "volumeMode">);

	/**
     * AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
     */
    accessModes?: Array<string>

	/**
     * Capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
     */
    capacity?: number | string

	/**
     * Csi represents storage that is handled by an external CSI driver.
     */
    csi?: CSIPersistentVolumeSource

	/**
     * Fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
     */
    fc?: FCVolumeSource

	/**
     * HostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
     */
    hostPath?: HostPathVolumeSource

	/**
     * Iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.
     */
    iscsi?: ISCSIPersistentVolumeSource

	/**
     * Local represents directly-attached storage with node affinity
     */
    local?: LocalVolumeSource

	/**
     * MountOptions is the list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
     */
    mountOptions?: Array<string>

	/**
     * Nfs represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
     */
    nfs?: NFSVolumeSource

	/**
     * NodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
     */
    nodeAffinity?: VolumeNodeAffinity

	/**
     * PersistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
     */
    persistentVolumeReclaimPolicy?: string

	/**
     * StorageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
     */
    storageClassName?: string

	/**
     * Name of VolumeAttributesClass to which this persistent volume belongs. Empty value is not allowed. When this field is not set, it indicates that this volume does not belong to any VolumeAttributesClass. This field is mutable and can be changed by the CSI driver after a volume has been updated successfully to a new class. For an unbound PersistentVolume, the volumeAttributesClassName will be matched with unbound PersistentVolumeClaims during the binding process.
     */
    volumeAttributesClassName?: string

	/**
     * VolumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.
     */
    volumeMode?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}