// Auto generated code; DO NOT EDIT.



/**
 * DeploymentCondition describes the state of a deployment at a certain point.
 * 
 */
export declare class DeploymentCondition {
    constructor();
    constructor(spec: DeploymentCondition);

	/**
     * Last time the condition transitioned from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * The last time this condition was updated.
     * 
     */
    lastUpdateTime?: string

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
     * Type of deployment condition.
     * 
     */
    type: string
}