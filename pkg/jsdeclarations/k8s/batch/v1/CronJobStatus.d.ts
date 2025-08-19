// Auto generated code; DO NOT EDIT.

import { ObjectReference } from "../../core/v1"

/**
 * CronJobStatus represents the current state of a cron job.
 * 
 */
export declare class CronJobStatus {
    constructor();
    constructor(spec: CronJobStatus);

	/**
     * A list of pointers to currently running jobs.
     * 
     */
    active?: ObjectReference

	/**
     * Information when was the last time the job was successfully scheduled.
     * 
     */
    lastScheduleTime?: string

	/**
     * Information when was the last time the job successfully completed.
     * 
     */
    lastSuccessfulTime?: string
}