// Auto generated code; DO NOT EDIT.

import { AttachedVolume } from "./AttachedVolume"
import { ContainerImage } from "./ContainerImage"
import { NodeAddress } from "./NodeAddress"
import { NodeCondition } from "./NodeCondition"
import { NodeConfigStatus } from "./NodeConfigStatus"
import { NodeDaemonEndpoints } from "./NodeDaemonEndpoints"
import { NodeFeatures } from "./NodeFeatures"
import { NodeRuntimeHandler } from "./NodeRuntimeHandler"
import { NodeSystemInfo } from "./NodeSystemInfo"

/**
 * NodeStatus is information about the current status of a node.
 * 
 */
export declare class NodeStatus {
    constructor();
    constructor(spec: NodeStatus);

	/**
     * List of addresses reachable to the node. Queried from cloud provider, if available. More info: https://kubernetes.io/docs/reference/node/node-status/#addresses Note: This field is declared as mergeable, but the merge key is not sufficiently unique, which can cause data corruption when it is merged. Callers should instead use a full-replacement patch. See https://pr.k8s.io/79391 for an example. Consumers should assume that addresses can change during the lifetime of a Node. However, there are some exceptions where this may not be possible, such as Pods that inherit a Node's address in its own status or consumers of the downward API (status.hostIP).
     * 
     */
    addresses?: Array<NodeAddress>

	/**
     * Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
     * 
     */
    allocatable?: Record<string, any>

	/**
     * Capacity represents the total resources of a node. More info: https://kubernetes.io/docs/reference/node/node-status/#capacity
     * 
     */
    capacity?: Record<string, any>

	/**
     * Conditions is an array of current observed node conditions. More info: https://kubernetes.io/docs/reference/node/node-status/#condition
     * 
     */
    conditions?: Array<NodeCondition>

	/**
     * Status of the config assigned to the node via the dynamic Kubelet config feature.
     * 
     */
    config?: NodeConfigStatus

	/**
     * Endpoints of daemons running on the Node.
     * 
     */
    daemonEndpoints?: NodeDaemonEndpoints

	/**
     * Features describes the set of features implemented by the CRI implementation.
     * 
     */
    features?: NodeFeatures

	/**
     * List of container images on this node
     * 
     */
    images?: Array<ContainerImage>

	/**
     * Set of ids/uuids to uniquely identify the node. More info: https://kubernetes.io/docs/reference/node/node-status/#info
     * 
     */
    nodeInfo?: NodeSystemInfo

	/**
     * NodePhase is the recently observed lifecycle phase of the node. More info: https://kubernetes.io/docs/concepts/nodes/node/#phase The field is never populated, and now is deprecated.
     * 
     */
    phase?: string

	/**
     * The available runtime handlers.
     * 
     */
    runtimeHandlers?: Array<NodeRuntimeHandler>

	/**
     * List of volumes that are attached to the node.
     * 
     */
    volumesAttached?: Array<AttachedVolume>

	/**
     * List of attachable volumes in use (mounted) by the node.
     * 
     */
    volumesInUse?: Array<string>
}