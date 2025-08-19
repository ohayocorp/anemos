// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { StatefulSetSpec } from "./StatefulSetSpec"
import { StatefulSetStatus } from "./StatefulSetStatus"

/**
 * StatefulSet represents a set of pods with consistent identities. Identities are defined as:
 * 
 *   - Network: A single stable DNS and hostname.
 * 
 *   - Storage: As many VolumeClaims as requested.
 * 
 * The StatefulSet guarantees that a given network identity will always map to the same storage identity.
 * 
 */
export declare class StatefulSet {
    constructor();
    constructor(spec: Omit<StatefulSet, "apiVersion" | "kind">);

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
     * Spec defines the desired identities of pods in this set.
     * 
     */
    spec?: StatefulSetSpec

	/**
     * Status is the current status of Pods in this StatefulSet. This data may be out of date by some window of time.
     * 
     */
    status?: StatefulSetStatus
}