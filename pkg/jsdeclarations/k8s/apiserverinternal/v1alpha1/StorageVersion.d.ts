// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { StorageVersionSpec } from "./StorageVersionSpec"
import { StorageVersionStatus } from "./StorageVersionStatus"

/**
 * Storage version of a specific resource.
 * 
 */
export declare class StorageVersion {
    constructor();
    constructor(spec: Omit<StorageVersion, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * The name is <group>.<resource>.
     * 
     */
    metadata?: ObjectMeta

	/**
     * Spec is an empty spec. It is here to comply with Kubernetes API style.
     * 
     */
    spec: StorageVersionSpec

	/**
     * API server instances report the version they can decode and the version they encode objects to when persisting objects in the backend.
     * 
     */
    status: StorageVersionStatus
}