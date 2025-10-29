// Auto generated code; DO NOT EDIT.

/**
 * A node selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
 */
export declare class NodeSelectorRequirement {
    constructor();
    constructor(spec: Pick<NodeSelectorRequirement, "key" | "operator" | "values">);

	/**
     * The label key that the selector applies to.
     */
    key: string

	/**
     * Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.
     */
    operator: string

	/**
     * An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.
     */
    values?: Array<string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}