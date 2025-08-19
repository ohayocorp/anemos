// Auto generated code; DO NOT EDIT.



/**
 * LimitRangeItem defines a min/max usage limit for any resource that matches on kind.
 * 
 */
export declare class LimitRangeItem {
    constructor();
    constructor(spec: LimitRangeItem);

	/**
     * Default resource requirement limit value by resource name if resource limit is omitted.
     * 
     */
    default?: any

	/**
     * DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
     * 
     */
    defaultRequest?: any

	/**
     * Max usage constraints on this kind by resource name.
     * 
     */
    max?: any

	/**
     * MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
     * 
     */
    maxLimitRequestRatio?: any

	/**
     * Min usage constraints on this kind by resource name.
     * 
     */
    min?: any

	/**
     * Type of resource that this limit applies to.
     * 
     */
    type: string
}