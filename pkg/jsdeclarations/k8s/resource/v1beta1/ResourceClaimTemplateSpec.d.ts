// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { ResourceClaimSpec } from "./ResourceClaimSpec"

/**
 * ResourceClaimTemplateSpec contains the metadata and fields for a ResourceClaim.
 * 
 */
export declare class ResourceClaimTemplateSpec {
    constructor();
    constructor(spec: ResourceClaimTemplateSpec);

	/**
     * ObjectMeta may contain labels and annotations that will be copied into the ResourceClaim when creating it. No other fields are allowed and will be rejected during validation.
     * 
     */
    metadata?: ObjectMeta

	/**
     * Spec for the ResourceClaim. The entire content is copied unchanged into the ResourceClaim that gets created from this template. The same fields as in a ResourceClaim are also valid here.
     * 
     */
    spec: ResourceClaimSpec
}