// Auto generated code; DO NOT EDIT.

/**
 * LimitRangeItem defines a min/max usage limit for any resource that matches on kind.
 */
export declare class LimitRangeItem {
    constructor();
    constructor(spec: Pick<LimitRangeItem, "default" | "defaultRequest" | "max" | "maxLimitRequestRatio" | "min" | "type">);

	/**
     * Default resource requirement limit value by resource name if resource limit is omitted.
     */
    default?: Record<string, number | string>

	/**
     * DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
     */
    defaultRequest?: Record<string, number | string>

	/**
     * Max usage constraints on this kind by resource name.
     */
    max?: Record<string, number | string>

	/**
     * MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
     */
    maxLimitRequestRatio?: Record<string, number | string>

	/**
     * Min usage constraints on this kind by resource name.
     */
    min?: Record<string, number | string>

	/**
     * Type of resource that this limit applies to.
     */
    type: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}