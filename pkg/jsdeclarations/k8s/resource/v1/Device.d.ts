// Auto generated code; DO NOT EDIT.

import { NodeSelector } from "../../core/v1"
import { DeviceCounterConsumption } from "./DeviceCounterConsumption"
import { DeviceTaint } from "./DeviceTaint"

/**
 * Device represents one individual hardware instance that can be selected based on its attributes. Besides the name, exactly one field must be set.
 * 
 */
export declare class Device {
    constructor();
    constructor(spec: Device);

	/**
     * AllNodes indicates that all nodes have access to the device.
     * 
     * Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
     * 
     */
    allNodes?: boolean

	/**
     * AllowMultipleAllocations marks whether the device is allowed to be allocated to multiple DeviceRequests.
     * 
     * If AllowMultipleAllocations is set to true, the device can be allocated more than once, and all of its capacity is consumable, regardless of whether the requestPolicy is defined or not.
     * 
     */
    allowMultipleAllocations?: boolean

	/**
     * Attributes defines the set of attributes for this device. The name of each attribute must be unique in that set.
     * 
     * The maximum number of attributes and capacities combined is 32.
     * 
     */
    attributes?: any

	/**
     * BindingConditions defines the conditions for proceeding with binding. All of these conditions must be set in the per-device status conditions with a value of True to proceed with binding the pod to the node while scheduling the pod.
     * 
     * The maximum number of binding conditions is 4.
     * 
     * The conditions must be a valid condition type string.
     * 
     * This is an alpha field and requires enabling the DRADeviceBindingConditions and DRAResourceClaimDeviceStatus feature gates.
     * 
     */
    bindingConditions?: Array<string>

	/**
     * BindingFailureConditions defines the conditions for binding failure. They may be set in the per-device status conditions. If any is set to "True", a binding failure occurred.
     * 
     * The maximum number of binding failure conditions is 4.
     * 
     * The conditions must be a valid condition type string.
     * 
     * This is an alpha field and requires enabling the DRADeviceBindingConditions and DRAResourceClaimDeviceStatus feature gates.
     * 
     */
    bindingFailureConditions?: Array<string>

	/**
     * BindsToNode indicates if the usage of an allocation involving this device has to be limited to exactly the node that was chosen when allocating the claim. If set to true, the scheduler will set the ResourceClaim.Status.Allocation.NodeSelector to match the node where the allocation was made.
     * 
     * This is an alpha field and requires enabling the DRADeviceBindingConditions and DRAResourceClaimDeviceStatus feature gates.
     * 
     */
    bindsToNode?: boolean

	/**
     * Capacity defines the set of capacities for this device. The name of each capacity must be unique in that set.
     * 
     * The maximum number of attributes and capacities combined is 32.
     * 
     */
    capacity?: any

	/**
     * ConsumesCounters defines a list of references to sharedCounters and the set of counters that the device will consume from those counter sets.
     * 
     * There can only be a single entry per counterSet.
     * 
     * The total number of device counter consumption entries must be <= 32. In addition, the total number in the entire ResourceSlice must be <= 1024 (for example, 64 devices with 16 counters each).
     * 
     */
    consumesCounters?: Array<DeviceCounterConsumption>

	/**
     * Name is unique identifier among all devices managed by the driver in the pool. It must be a DNS label.
     * 
     */
    name: string

	/**
     * NodeName identifies the node where the device is available.
     * 
     * Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
     * 
     */
    nodeName?: string

	/**
     * NodeSelector defines the nodes where the device is available.
     * 
     * Must use exactly one term.
     * 
     * Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
     * 
     */
    nodeSelector?: NodeSelector

	/**
     * If specified, these are the driver-defined taints.
     * 
     * The maximum number of taints is 4.
     * 
     * This is an alpha field and requires enabling the DRADeviceTaints feature gate.
     * 
     */
    taints?: Array<DeviceTaint>
}