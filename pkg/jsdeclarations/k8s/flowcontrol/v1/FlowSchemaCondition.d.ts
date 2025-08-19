// Auto generated code; DO NOT EDIT.



/**
 * FlowSchemaCondition describes conditions for a FlowSchema.
 * 
 */
export declare class FlowSchemaCondition {
    constructor();
    constructor(spec: FlowSchemaCondition);

	/**
     * `lastTransitionTime` is the last time the condition transitioned from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * `message` is a human-readable message indicating details about last transition.
     * 
     */
    message?: string

	/**
     * `reason` is a unique, one-word, CamelCase reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * `status` is the status of the condition. Can be True, False, Unknown. Required.
     * 
     */
    status?: string

	/**
     * `type` is the type of the condition. Required.
     * 
     */
    type?: string
}