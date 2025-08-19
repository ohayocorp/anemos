// Auto generated code; DO NOT EDIT.

import { NodeSelector } from "../../core/v1"
import { CounterSet } from "./CounterSet"
import { Device } from "./Device"
import { ResourcePool } from "./ResourcePool"

/**
 * ResourceSliceSpec contains the information published by the driver in one ResourceSlice.
 * 
 */
export declare class ResourceSliceSpec {
    constructor();
    constructor(spec: ResourceSliceSpec);

	/**
     * AllNodes indicates that all nodes have access to the resources in the pool.
     * 
     * Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set.
     * 
     */
    allNodes?: boolean

	/**
     * Devices lists some or all of the devices in this pool.
     * 
     * Must not have more than 128 entries.
     * 
     */
    devices?: Array<Device>

	/**
     * Driver identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name.
     * 
     * Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver. This field is immutable.
     * 
     */
    driver: string

	/**
     * NodeName identifies the node which provides the resources in this pool. A field selector can be used to list only ResourceSlice objects belonging to a certain node.
     * 
     * This field can be used to limit access from nodes to ResourceSlices with the same node name. It also indicates to autoscalers that adding new nodes of the same type as some old node might also make new resources available.
     * 
     * Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set. This field is immutable.
     * 
     */
    nodeName?: string

	/**
     * NodeSelector defines which nodes have access to the resources in the pool, when that pool is not limited to a single node.
     * 
     * Must use exactly one term.
     * 
     * Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set.
     * 
     */
    nodeSelector?: NodeSelector

	/**
     * PerDeviceNodeSelection defines whether the access from nodes to resources in the pool is set on the ResourceSlice level or on each device. If it is set to true, every device defined the ResourceSlice must specify this individually.
     * 
     * Exactly one of NodeName, NodeSelector, AllNodes, and PerDeviceNodeSelection must be set.
     * 
     */
    perDeviceNodeSelection?: boolean

	/**
     * Pool describes the pool that this ResourceSlice belongs to.
     * 
     */
    pool: ResourcePool

	/**
     * SharedCounters defines a list of counter sets, each of which has a name and a list of counters available.
     * 
     * The names of the SharedCounters must be unique in the ResourceSlice.
     * 
     * The maximum number of SharedCounters is 32.
     * 
     */
    sharedCounters?: Array<CounterSet>
}