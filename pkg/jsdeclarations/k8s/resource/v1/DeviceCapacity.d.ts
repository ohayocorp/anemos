// Auto generated code; DO NOT EDIT.

import { CapacityRequestPolicy } from "./CapacityRequestPolicy"

/**
 * DeviceCapacity describes a quantity associated with a device.
 * 
 */
export declare class DeviceCapacity {
    constructor();
    constructor(spec: DeviceCapacity);

	/**
     * RequestPolicy defines how this DeviceCapacity must be consumed when the device is allowed to be shared by multiple allocations.
     * 
     * The Device must have allowMultipleAllocations set to true in order to set a requestPolicy.
     * 
     * If unset, capacity requests are unconstrained: requests can consume any amount of capacity, as long as the total consumed across all allocations does not exceed the device's defined capacity. If request is also unset, default is the full capacity value.
     * 
     */
    requestPolicy?: CapacityRequestPolicy

	/**
     * Value defines how much of a certain capacity that device has.
     * 
     * This field reflects the fixed total capacity and does not change. The consumed amount is tracked separately by scheduler and does not affect this value.
     * 
     */
    value: any
}