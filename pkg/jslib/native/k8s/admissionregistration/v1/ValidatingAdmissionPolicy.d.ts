// Auto generated code; DO NOT EDIT.
import { ValidatingAdmissionPolicySpec } from "./ValidatingAdmissionPolicySpec"
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import {Document} from '@ohayocorp/anemos';

/**
 * ValidatingAdmissionPolicy describes the definition of an admission validation policy that accepts or rejects an object without changing it.
 */
export declare class ValidatingAdmissionPolicy extends Document {
    constructor();
    constructor(spec: Pick<ValidatingAdmissionPolicy, "metadata" | "spec">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     */
    metadata?: ObjectMeta

	/**
     * Specification of the desired behavior of the ValidatingAdmissionPolicy.
     */
    spec?: ValidatingAdmissionPolicySpec
}