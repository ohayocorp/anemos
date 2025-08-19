// Auto generated code; DO NOT EDIT.



/**
 * ReplicaSetCondition describes the state of a replica set at a certain point.
 * 
 */
export declare class ReplicaSetCondition {
    constructor();
    constructor(spec: ReplicaSetCondition);

	/**
     * The last time the condition transitioned from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * A human readable message indicating details about the transition.
     * 
     */
    message?: string

	/**
     * The reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * Type of replica set condition.
     * 
     */
    type: string
}