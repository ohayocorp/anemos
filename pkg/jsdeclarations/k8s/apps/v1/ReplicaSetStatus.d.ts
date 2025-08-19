// Auto generated code; DO NOT EDIT.

import { ReplicaSetCondition } from "./ReplicaSetCondition"

/**
 * ReplicaSetStatus represents the current status of a ReplicaSet.
 * 
 */
export declare class ReplicaSetStatus {
    constructor();
    constructor(spec: ReplicaSetStatus);

	/**
     * The number of available non-terminating pods (ready for at least minReadySeconds) for this replica set.
     * 
     */
    availableReplicas?: number

	/**
     * Represents the latest available observations of a replica set's current state.
     * 
     */
    conditions?: Array<ReplicaSetCondition>

	/**
     * The number of non-terminating pods that have labels matching the labels of the pod template of the replicaset.
     * 
     */
    fullyLabeledReplicas?: number

	/**
     * ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
     * 
     */
    observedGeneration?: number

	/**
     * The number of non-terminating pods targeted by this ReplicaSet with a Ready Condition.
     * 
     */
    readyReplicas?: number

	/**
     * Replicas is the most recently observed number of non-terminating pods. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicaset
     * 
     */
    replicas: number

	/**
     * The number of terminating pods for this replica set. Terminating pods have a non-null .metadata.deletionTimestamp and have not yet reached the Failed or Succeeded .status.phase.
     * 
     * This is an alpha field. Enable DeploymentReplicaSetTerminatingReplicas to be able to use this field.
     * 
     */
    terminatingReplicas?: number
}