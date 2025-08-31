// Auto generated code; DO NOT EDIT.
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import { TopologySelectorTerm } from "./../../core/v1"
import {Document} from '@ohayocorp/anemos';

/**
 * StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.

 * StorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.
 */
export declare class StorageClass extends Document {
    constructor();
    constructor(spec: Pick<StorageClass, "allowVolumeExpansion" | "allowedTopologies" | "metadata" | "mountOptions" | "parameters" | "provisioner" | "reclaimPolicy" | "volumeBindingMode">);

	/**
     * AllowVolumeExpansion shows whether the storage class allow volume expand.
     */
    allowVolumeExpansion?: boolean

	/**
     * AllowedTopologies restrict the node topologies where volumes can be dynamically provisioned. Each volume plugin defines its own supported topology specifications. An empty TopologySelectorTerm list means there is no topology restriction. This field is only honored by servers that enable the VolumeScheduling feature.
     */
    allowedTopologies?: TopologySelectorTerm

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta

	/**
     * MountOptions controls the mountOptions for dynamically provisioned PersistentVolumes of this storage class. e.g. ["ro", "soft"]. Not validated - mount of the PVs will simply fail if one is invalid.
     */
    mountOptions?: Array<string>

	/**
     * Parameters holds the parameters for the provisioner that should create volumes of this storage class.
     */
    parameters?: Record<string, string>

	/**
     * Provisioner indicates the type of the provisioner.
     */
    provisioner: string

	/**
     * ReclaimPolicy controls the reclaimPolicy for dynamically provisioned PersistentVolumes of this storage class. Defaults to Delete.
     */
    reclaimPolicy?: string

	/**
     * VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is only honored by servers that enable the VolumeScheduling feature.
     */
    volumeBindingMode?: string
}