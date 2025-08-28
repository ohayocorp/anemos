// Auto generated code; DO NOT EDIT.
import { ResourceClaim } from "./ResourceClaim"

/**
 * ResourceRequirements describes the compute resource requirements.
 */
export declare class ResourceRequirements {
    constructor();
    constructor(spec: Pick<ResourceRequirements, "claims" | "limits" | "requests">);

	/**
     * Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.
    
     * This field depends on the DynamicResourceAllocation feature gate.
    
     * This field is immutable. It can only be set for containers.
     */
    claims?: Array<ResourceClaim>

	/**
     * Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     */
    limits?: number | string

	/**
     * Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     */
    requests?: number | string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}