// Auto generated code; DO NOT EDIT.

/**
 * CustomResourceColumnDefinition specifies a column for server side printing.
 */
export declare class CustomResourceColumnDefinition {
    constructor();
    constructor(spec: Pick<CustomResourceColumnDefinition, "description" | "format" | "jsonPath" | "name" | "priority" | "type">);

	/**
     * Description is a human readable description of this column.
     */
    description?: string

	/**
     * Format is an optional OpenAPI type definition for this column. The 'name' format is applied to the primary identifier column to assist in clients identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
     */
    format?: string

	/**
     * JsonPath is a simple JSON path (i.e. with array notation) which is evaluated against each custom resource to produce the value for this column.
     */
    jsonPath: string

	/**
     * Name is a human readable name for the column.
     */
    name: string

	/**
     * Priority is an integer defining the relative importance of this column compared to others. Lower numbers are considered higher priority. Columns that may be omitted in limited space scenarios should be given a priority greater than 0.
     */
    priority?: number

	/**
     * Type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types for details.
     */
    type: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}