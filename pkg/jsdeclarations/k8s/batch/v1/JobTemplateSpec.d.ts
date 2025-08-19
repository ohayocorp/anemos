// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { JobSpec } from "./JobSpec"

/**
 * JobTemplateSpec describes the data a Job should have when created from a template
 * 
 */
export declare class JobTemplateSpec {
    constructor();
    constructor(spec: JobTemplateSpec);

	/**
     * Standard object's metadata of the jobs created from this template. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Specification of the desired behavior of the job. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     * 
     */
    spec?: JobSpec
}