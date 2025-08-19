// Auto generated code; DO NOT EDIT.



/**
 * DaemonSetCondition describes the state of a DaemonSet at a certain point.
 * 
 */
export declare class DaemonSetCondition {
    constructor();
    constructor(spec: DaemonSetCondition);

	/**
     * Last time the condition transitioned from one status to another.
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
     * Type of DaemonSet condition.
     * 
     */
    type: string
}