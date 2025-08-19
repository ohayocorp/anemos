// Auto generated code; DO NOT EDIT.

import { PriorityLevelConfigurationCondition } from "./PriorityLevelConfigurationCondition"

/**
 * PriorityLevelConfigurationStatus represents the current state of a "request-priority".
 * 
 */
export declare class PriorityLevelConfigurationStatus {
    constructor();
    constructor(spec: PriorityLevelConfigurationStatus);

	/**
     * `conditions` is the current state of "request-priority".
     * 
     */
    conditions?: Array<PriorityLevelConfigurationCondition>
}