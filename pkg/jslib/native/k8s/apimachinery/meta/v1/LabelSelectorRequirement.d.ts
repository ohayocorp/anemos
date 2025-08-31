// Auto generated code; DO NOT EDIT.

/**
 * A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
 */
export declare class LabelSelectorRequirement {
    constructor();
    constructor(spec: Pick<LabelSelectorRequirement, "key" | "operator" | "values">);

	/**
     * Key is the label key that the selector applies to.
     */
    key: string

	/**
     * Operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.
     */
    operator: string

	/**
     * Values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.
     */
    values?: Array<string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}