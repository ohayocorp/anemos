// Auto generated code; DO NOT EDIT.
import { LimitRangeItem } from "./LimitRangeItem"

/**
 * LimitRangeSpec defines a min/max usage limit for resources that match on kind.
 */
export declare class LimitRangeSpec {
    constructor();
    constructor(spec: Pick<LimitRangeSpec, "limits">);

	/**
     * Limits is the list of LimitRangeItem objects that are enforced.
     */
    limits: Array<LimitRangeItem>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}