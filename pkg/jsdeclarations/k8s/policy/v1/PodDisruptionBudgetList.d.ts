// Auto generated code; DO NOT EDIT.

import { ListMeta } from "../../apimachinery/meta/v1"
import { PodDisruptionBudget } from "./PodDisruptionBudget"

/**
 * PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
 * 
 */
export declare class PodDisruptionBudgetList {
    constructor();
    constructor(spec: Omit<PodDisruptionBudgetList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Items is a list of PodDisruptionBudgets
     * 
     */
    items: Array<PodDisruptionBudget>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    metadata?: ListMeta
}