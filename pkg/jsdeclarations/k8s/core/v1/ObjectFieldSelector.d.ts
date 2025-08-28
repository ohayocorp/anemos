// Auto generated code; DO NOT EDIT.

/**
 * ObjectFieldSelector selects an APIVersioned field of an object.
 */
export declare class ObjectFieldSelector {
    constructor();
    constructor(spec: Pick<ObjectFieldSelector, "fieldPath">);

	/**
     * Version of the schema the FieldPath is written in terms of, defaults to "v1".
     */
    apiVersion?: string

	/**
     * Path of the field to select in the specified API version.
     */
    fieldPath: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}