// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { PodDisruptionBudgetSpec } from "./PodDisruptionBudgetSpec"
import { PodDisruptionBudgetStatus } from "./PodDisruptionBudgetStatus"

/**
 * PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
 * 
 */
export declare class PodDisruptionBudget {
    constructor();
    constructor(spec: Omit<PodDisruptionBudget, "apiVersion" | "kind">);

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
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Specification of the desired behavior of the PodDisruptionBudget.
     * 
     */
    spec?: PodDisruptionBudgetSpec

	/**
     * Most recently observed status of the PodDisruptionBudget.
     * 
     */
    status?: PodDisruptionBudgetStatus
}