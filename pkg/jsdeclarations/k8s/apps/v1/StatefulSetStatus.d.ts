// Auto generated code; DO NOT EDIT.

import { StatefulSetCondition } from "./StatefulSetCondition"

/**
 * StatefulSetStatus represents the current state of a StatefulSet.
 * 
 */
export declare class StatefulSetStatus {
    constructor();
    constructor(spec: StatefulSetStatus);

	/**
     * Total number of available pods (ready for at least minReadySeconds) targeted by this statefulset.
     * 
     */
    availableReplicas?: number

	/**
     * CollisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
     * 
     */
    collisionCount?: number

	/**
     * Represents the latest available observations of a statefulset's current state.
     * 
     */
    conditions?: Array<StatefulSetCondition>

	/**
     * CurrentReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by currentRevision.
     * 
     */
    currentReplicas?: number

	/**
     * CurrentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [0,currentReplicas).
     * 
     */
    currentRevision?: string

	/**
     * ObservedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
     * 
     */
    observedGeneration?: number

	/**
     * ReadyReplicas is the number of pods created for this StatefulSet with a Ready Condition.
     * 
     */
    readyReplicas?: number

	/**
     * Replicas is the number of Pods created by the StatefulSet controller.
     * 
     */
    replicas: number

	/**
     * UpdateRevision, if not empty, indicates the version of the StatefulSet used to generate Pods in the sequence [replicas-updatedReplicas,replicas)
     * 
     */
    updateRevision?: string

	/**
     * UpdatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
     * 
     */
    updatedReplicas?: number
}