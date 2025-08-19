// Auto generated code; DO NOT EDIT.



/**
 * PodCondition contains details for the current condition of this pod.
 * 
 */
export declare class PodCondition {
    constructor();
    constructor(spec: PodCondition);

	/**
     * Last time we probed the condition.
     * 
     */
    lastProbeTime?: string

	/**
     * Last time the condition transitioned from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * Human-readable message indicating details about last transition.
     * 
     */
    message?: string

	/**
     * If set, this represents the .metadata.generation that the pod condition was set based upon. This is an alpha field. Enable PodObservedGenerationTracking to be able to use this field.
     * 
     */
    observedGeneration?: number

	/**
     * Unique, one-word, CamelCase reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions
     * 
     */
    type: string
}