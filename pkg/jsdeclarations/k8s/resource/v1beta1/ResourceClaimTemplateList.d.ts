// Auto generated code; DO NOT EDIT.

import { ListMeta } from "../../apimachinery/meta/v1"
import { ResourceClaimTemplate } from "./ResourceClaimTemplate"

/**
 * ResourceClaimTemplateList is a collection of claim templates.
 * 
 */
export declare class ResourceClaimTemplateList {
    constructor();
    constructor(spec: Omit<ResourceClaimTemplateList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Items is the list of resource claim templates.
     * 
     */
    items: Array<ResourceClaimTemplate>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard list metadata
     * 
     */
    metadata?: ListMeta
}