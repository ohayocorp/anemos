// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { ResourceClaimTemplateSpec } from "./ResourceClaimTemplateSpec"

/**
 * ResourceClaimTemplate is used to produce ResourceClaim objects.
 * 
 * This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
 * 
 */
export declare class ResourceClaimTemplate {
    constructor();
    constructor(spec: Omit<ResourceClaimTemplate, "apiVersion" | "kind">);

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
     * Standard object metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Describes the ResourceClaim that is to be generated.
     * 
     * This field is immutable. A ResourceClaim will get created by the control plane for a Pod when needed and then not get updated anymore.
     * 
     */
    spec: ResourceClaimTemplateSpec
}