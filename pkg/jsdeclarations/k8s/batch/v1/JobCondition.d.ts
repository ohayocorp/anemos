// Auto generated code; DO NOT EDIT.



/**
 * JobCondition describes current state of a job.
 * 
 */
export declare class JobCondition {
    constructor();
    constructor(spec: JobCondition);

	/**
     * Last time the condition was checked.
     * 
     */
    lastProbeTime?: string

	/**
     * Last time the condition transit from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * Human readable message indicating details about last transition.
     * 
     */
    message?: string

	/**
     * (brief) reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * Status of the condition, one of True, False, Unknown.
     * 
     */
    status: string

	/**
     * Type of job condition, Complete or Failed.
     * 
     */
    type: string
}