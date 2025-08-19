// Auto generated code; DO NOT EDIT.

import { LimitRangeItem } from "./LimitRangeItem"

/**
 * LimitRangeSpec defines a min/max usage limit for resources that match on kind.
 * 
 */
export declare class LimitRangeSpec {
    constructor();
    constructor(spec: LimitRangeSpec);

	/**
     * Limits is the list of LimitRangeItem objects that are enforced.
     * 
     */
    limits: Array<LimitRangeItem>
}