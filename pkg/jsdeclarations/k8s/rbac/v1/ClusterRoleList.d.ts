// Auto generated code; DO NOT EDIT.

import { ListMeta } from "../../apimachinery/meta/v1"
import { ClusterRole } from "./ClusterRole"

/**
 * ClusterRoleList is a collection of ClusterRoles
 * 
 */
export declare class ClusterRoleList {
    constructor();
    constructor(spec: Omit<ClusterRoleList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Items is a list of ClusterRoles
     * 
     */
    items: Array<ClusterRole>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard object's metadata.
     * 
     */
    metadata?: ListMeta
}