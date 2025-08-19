// Auto generated code; DO NOT EDIT.

import { ListMeta } from "../../apimachinery/meta/v1"
import { ClusterTrustBundle } from "./ClusterTrustBundle"

/**
 * ClusterTrustBundleList is a collection of ClusterTrustBundle objects
 * 
 */
export declare class ClusterTrustBundleList {
    constructor();
    constructor(spec: Omit<ClusterTrustBundleList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Items is a collection of ClusterTrustBundle objects
     * 
     */
    items: Array<ClusterTrustBundle>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Metadata contains the list metadata.
     * 
     */
    metadata?: ListMeta
}