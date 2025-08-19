// Auto generated code; DO NOT EDIT.



/**
 * NodeCondition contains condition information for a node.
 * 
 */
export declare class NodeCondition {
    constructor();
    constructor(spec: NodeCondition);

	/**
     * Last time we got an update on a given condition.
     * 
     */
    lastHeartbeatTime?: string

	/**
     * Last time the condition transit from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * Human readable message indicating details about last transition.
     * 
     */
    message?: string

	/**
     * (brief) reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * Type of node condition.
     * 
     */
    type: string
}