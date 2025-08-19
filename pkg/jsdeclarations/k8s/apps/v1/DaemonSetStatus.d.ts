// Auto generated code; DO NOT EDIT.

import { DaemonSetCondition } from "./DaemonSetCondition"

/**
 * DaemonSetStatus represents the current status of a daemon set.
 * 
 */
export declare class DaemonSetStatus {
    constructor();
    constructor(spec: DaemonSetStatus);

	/**
     * Count of hash collisions for the DaemonSet. The DaemonSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
     * 
     */
    collisionCount?: number

	/**
     * Represents the latest available observations of a DaemonSet's current state.
     * 
     */
    conditions?: Array<DaemonSetCondition>

	/**
     * The number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
     * 
     */
    currentNumberScheduled: number

	/**
     * The total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
     * 
     */
    desiredNumberScheduled: number

	/**
     * The number of nodes that should be running the daemon pod and have one or more of the daemon pod running and available (ready for at least spec.minReadySeconds)
     * 
     */
    numberAvailable?: number

	/**
     * The number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: https://kubernetes.io/docs/concepts/workloads/controllers/daemonset/
     * 
     */
    numberMisscheduled: number

	/**
     * NumberReady is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running with a Ready Condition.
     * 
     */
    numberReady: number

	/**
     * The number of nodes that should be running the daemon pod and have none of the daemon pod running and available (ready for at least spec.minReadySeconds)
     * 
     */
    numberUnavailable?: number

	/**
     * The most recent generation observed by the daemon set controller.
     * 
     */
    observedGeneration?: number

	/**
     * The total number of nodes that are running updated daemon pod
     * 
     */
    updatedNumberScheduled?: number
}