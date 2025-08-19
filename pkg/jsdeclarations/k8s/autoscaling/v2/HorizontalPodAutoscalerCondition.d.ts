// Auto generated code; DO NOT EDIT.



/**
 * HorizontalPodAutoscalerCondition describes the state of a HorizontalPodAutoscaler at a certain point.
 * 
 */
export declare class HorizontalPodAutoscalerCondition {
    constructor();
    constructor(spec: HorizontalPodAutoscalerCondition);

	/**
     * LastTransitionTime is the last time the condition transitioned from one status to another
     * 
     */
    lastTransitionTime?: string

	/**
     * Message is a human-readable explanation containing details about the transition
     * 
     */
    message?: string

	/**
     * Reason is the reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * Status is the status of the condition (True, False, Unknown)
     * 
     */
    status: string

	/**
     * Type describes the current condition
     * 
     */
    type: string
}