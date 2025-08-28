// Auto generated code; DO NOT EDIT.
import { ValidatingAdmissionPolicyBindingSpec } from "./ValidatingAdmissionPolicyBindingSpec"
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import {Document} from '@ohayocorp/anemos';

/**
 * ValidatingAdmissionPolicyBinding binds the ValidatingAdmissionPolicy with paramerized resources. ValidatingAdmissionPolicyBinding and parameter CRDs together define how cluster administrators configure policies for clusters.

 * For a given admission request, each binding will cause its policy to be evaluated N times, where N is 1 for policies/bindings that don't use params, otherwise N is the number of parameters selected by the binding.

 * The CEL expressions of a policy must have a computed CEL cost below the maximum CEL budget. Each evaluation of the policy is given an independent CEL cost budget. Adding/removing policies, bindings, or params can not affect whether a given (policy, binding, param) combination is within its own CEL budget.
 */
export declare class ValidatingAdmissionPolicyBinding extends Document {
    constructor();
    constructor(spec: Pick<ValidatingAdmissionPolicyBinding, "metadata" | "spec">);

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
     * Specification of the desired behavior of the ValidatingAdmissionPolicyBinding.
     */
    spec?: ValidatingAdmissionPolicyBindingSpec
}