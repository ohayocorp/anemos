// Auto generated code; DO NOT EDIT.

import { ValidatingAdmissionPolicyBinding } from "./ValidatingAdmissionPolicyBinding"
import { ListMeta } from "../../apimachinery/meta/v1"

/**
 * ValidatingAdmissionPolicyBindingList is a list of ValidatingAdmissionPolicyBinding.
 * 
 */
export declare class ValidatingAdmissionPolicyBindingList {
    constructor();
    constructor(spec: Omit<ValidatingAdmissionPolicyBindingList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * List of PolicyBinding.
     * 
     */
    items: Array<ValidatingAdmissionPolicyBinding>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    metadata?: ListMeta
}