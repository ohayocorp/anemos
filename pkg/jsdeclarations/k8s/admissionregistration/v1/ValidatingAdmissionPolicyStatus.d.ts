// Auto generated code; DO NOT EDIT.

import { TypeChecking } from "./TypeChecking"
import { Condition } from "../../apimachinery/meta/v1"

/**
 * ValidatingAdmissionPolicyStatus represents the status of an admission validation policy.
 * 
 */
export declare class ValidatingAdmissionPolicyStatus {
    constructor();
    constructor(spec: ValidatingAdmissionPolicyStatus);

	/**
     * The conditions represent the latest available observations of a policy's current state.
     * 
     */
    conditions?: Condition

	/**
     * The generation observed by the controller.
     * 
     */
    observedGeneration?: number

	/**
     * The results of type checking for each expression. Presence of this field indicates the completion of the type checking.
     * 
     */
    typeChecking?: TypeChecking
}