// Auto generated code; DO NOT EDIT.

import { ListMeta } from "../../apimachinery/meta/v1"
import { PersistentVolume } from "./PersistentVolume"

/**
 * PersistentVolumeList is a list of PersistentVolume items.
 * 
 */
export declare class PersistentVolumeList {
    constructor();
    constructor(spec: Omit<PersistentVolumeList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Items is a list of persistent volumes. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes
     * 
     */
    items: Array<PersistentVolume>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    metadata?: ListMeta
}