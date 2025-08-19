// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { DeviceTaintRuleSpec } from "./DeviceTaintRuleSpec"

/**
 * DeviceTaintRule adds one taint to all devices which match the selector. This has the same effect as if the taint was specified directly in the ResourceSlice by the DRA driver.
 * 
 */
export declare class DeviceTaintRule {
    constructor();
    constructor(spec: Omit<DeviceTaintRule, "apiVersion" | "kind">);

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
     * Standard object metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Spec specifies the selector and one taint.
     * 
     * Changing the spec automatically increments the metadata.generation number.
     * 
     */
    spec: DeviceTaintRuleSpec
}