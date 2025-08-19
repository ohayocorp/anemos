// Auto generated code; DO NOT EDIT.

import { ContainerExtendedResourceRequest } from "./ContainerExtendedResourceRequest"

/**
 * PodExtendedResourceClaimStatus is stored in the PodStatus for the extended resource requests backed by DRA. It stores the generated name for the corresponding special ResourceClaim created by the scheduler.
 * 
 */
export declare class PodExtendedResourceClaimStatus {
    constructor();
    constructor(spec: PodExtendedResourceClaimStatus);

	/**
     * RequestMappings identifies the mapping of <container, extended resource backed by DRA> to  device request in the generated ResourceClaim.
     * 
     */
    requestMappings: Array<ContainerExtendedResourceRequest>

	/**
     * ResourceClaimName is the name of the ResourceClaim that was generated for the Pod in the namespace of the Pod.
     * 
     */
    resourceClaimName: string
}