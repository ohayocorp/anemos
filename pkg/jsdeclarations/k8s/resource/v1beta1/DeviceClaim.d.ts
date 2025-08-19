// Auto generated code; DO NOT EDIT.

import { DeviceClaimConfiguration } from "./DeviceClaimConfiguration"
import { DeviceConstraint } from "./DeviceConstraint"
import { DeviceRequest } from "./DeviceRequest"

/**
 * DeviceClaim defines how to request devices with a ResourceClaim.
 * 
 */
export declare class DeviceClaim {
    constructor();
    constructor(spec: DeviceClaim);

	/**
     * This field holds configuration for multiple potential drivers which could satisfy requests in this claim. It is ignored while allocating the claim.
     * 
     */
    config?: Array<DeviceClaimConfiguration>

	/**
     * These constraints must be satisfied by the set of devices that get allocated for the claim.
     * 
     */
    constraints?: Array<DeviceConstraint>

	/**
     * Requests represent individual requests for distinct devices which must all be satisfied. If empty, nothing needs to be allocated.
     * 
     */
    requests?: Array<DeviceRequest>
}