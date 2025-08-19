// Auto generated code; DO NOT EDIT.

import { DeviceClaim } from "./DeviceClaim"

/**
 * ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
 * 
 */
export declare class ResourceClaimSpec {
    constructor();
    constructor(spec: ResourceClaimSpec);

	/**
     * Devices defines how to request devices.
     * 
     */
    devices?: DeviceClaim
}