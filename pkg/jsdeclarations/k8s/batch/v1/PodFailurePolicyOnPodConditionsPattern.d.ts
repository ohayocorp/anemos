// Auto generated code; DO NOT EDIT.

/**
 * PodFailurePolicyOnPodConditionsPattern describes a pattern for matching an actual pod condition type.
 */
export declare class PodFailurePolicyOnPodConditionsPattern {
    constructor();
    constructor(spec: Pick<PodFailurePolicyOnPodConditionsPattern, "status" | "type">);

	/**
     * Specifies the required Pod condition status. To match a pod condition it is required that the specified status equals the pod condition status. Defaults to True.
     */
    status: string

	/**
     * Specifies the required Pod condition type. To match a pod condition it is required that specified type equals the pod condition type.
     */
    type: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}