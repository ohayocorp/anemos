// Auto generated code; DO NOT EDIT.

import { FlowSchemaCondition } from "./FlowSchemaCondition"

/**
 * FlowSchemaStatus represents the current state of a FlowSchema.
 * 
 */
export declare class FlowSchemaStatus {
    constructor();
    constructor(spec: FlowSchemaStatus);

	/**
     * `conditions` is a list of the current states of FlowSchema.
     * 
     */
    conditions?: Array<FlowSchemaCondition>
}