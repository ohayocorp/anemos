// Auto generated code; DO NOT EDIT.

/**
 * PodReadinessGate contains the reference to a pod condition
 */
export declare class PodReadinessGate {
    constructor();
    constructor(spec: Pick<PodReadinessGate, "conditionType">);

	/**
     * ConditionType refers to a condition in the pod's condition list with matching type.
     */
    conditionType: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}