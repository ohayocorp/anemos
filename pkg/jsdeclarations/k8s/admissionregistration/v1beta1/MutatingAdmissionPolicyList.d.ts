// Auto generated code; DO NOT EDIT.

import { MutatingAdmissionPolicy } from "./MutatingAdmissionPolicy"
import { ListMeta } from "../../apimachinery/meta/v1"

/**
 * MutatingAdmissionPolicyList is a list of MutatingAdmissionPolicy.
 * 
 */
export declare class MutatingAdmissionPolicyList {
    constructor();
    constructor(spec: Omit<MutatingAdmissionPolicyList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * List of ValidatingAdmissionPolicy.
     * 
     */
    items: Array<MutatingAdmissionPolicy>

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