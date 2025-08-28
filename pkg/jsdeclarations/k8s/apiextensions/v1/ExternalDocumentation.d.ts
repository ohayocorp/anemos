// Auto generated code; DO NOT EDIT.

/**
 * ExternalDocumentation allows referencing an external resource for extended documentation.
 */
export declare class ExternalDocumentation {
    constructor();
    constructor(spec: Pick<ExternalDocumentation, "description" | "url">);

	/**
     */
    description?: string

	/**
     */
    url?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}