// Auto generated code; DO NOT EDIT.

import { NodeSelector } from "../../core/v1"
import { DeviceAllocationResult } from "./DeviceAllocationResult"

/**
 * AllocationResult contains attributes of an allocated resource.
 * 
 */
export declare class AllocationResult {
    constructor();
    constructor(spec: AllocationResult);

	/**
     * AllocationTimestamp stores the time when the resources were allocated. This field is not guaranteed to be set, in which case that time is unknown.
     * 
     * This is an alpha field and requires enabling the DRADeviceBindingConditions and DRAResourceClaimDeviceStatus feature gate.
     * 
     */
    allocationTimestamp?: string

	/**
     * Devices is the result of allocating devices.
     * 
     */
    devices?: DeviceAllocationResult

	/**
     * NodeSelector defines where the allocated resources are available. If unset, they are available everywhere.
     * 
     */
    nodeSelector?: NodeSelector
}