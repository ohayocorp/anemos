// Auto generated code; DO NOT EDIT.

import { APIResource } from "./APIResource"

/**
 * APIResourceList is a list of APIResource, it is used to expose the name of the resources supported in a specific group and version, and if the resource is namespaced.
 * 
 */
export declare class APIResourceList {
    constructor();
    constructor(spec: Omit<APIResourceList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * GroupVersion is the group and version this APIResourceList is for.
     * 
     */
    groupVersion: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Resources contains the name of the resources and if they are namespaced.
     * 
     */
    resources: Array<APIResource>
}