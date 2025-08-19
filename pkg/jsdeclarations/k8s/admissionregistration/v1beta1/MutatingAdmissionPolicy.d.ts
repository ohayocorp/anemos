// Auto generated code; DO NOT EDIT.

import { MutatingAdmissionPolicySpec } from "./MutatingAdmissionPolicySpec"
import { ObjectMeta } from "../../apimachinery/meta/v1"

/**
 * MutatingAdmissionPolicy describes the definition of an admission mutation policy that mutates the object coming into admission chain.
 * 
 */
export declare class MutatingAdmissionPolicy {
    constructor();
    constructor(spec: Omit<MutatingAdmissionPolicy, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     * 
     */
    metadata?: ObjectMeta

	/**
     * Specification of the desired behavior of the MutatingAdmissionPolicy.
     * 
     */
    spec?: MutatingAdmissionPolicySpec
}