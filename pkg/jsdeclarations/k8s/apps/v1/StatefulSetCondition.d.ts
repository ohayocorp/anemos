// Auto generated code; DO NOT EDIT.



/**
 * StatefulSetCondition describes the state of a statefulset at a certain point.
 * 
 */
export declare class StatefulSetCondition {
    constructor();
    constructor(spec: StatefulSetCondition);

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
     * Status of the condition, one of True, False, Unknown.
     * 
     */
    status: string

	/**
     * Type of statefulset condition.
     * 
     */
    type: string
}