// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { PodSpec } from "./PodSpec"

/**
 * PodTemplateSpec describes the data a pod should have when created from a template
 * 
 */
export declare class PodTemplateSpec {
    constructor();
    constructor(spec: PodTemplateSpec);

	/**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    spec?: PodSpec
}