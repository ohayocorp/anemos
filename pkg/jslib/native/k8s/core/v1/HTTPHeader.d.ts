// Auto generated code; DO NOT EDIT.

/**
 * HTTPHeader describes a custom header to be used in HTTP probes
 */
export declare class HTTPHeader {
    constructor();
    constructor(spec: Pick<HTTPHeader, "name" | "value">);

	/**
     * The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.
     */
    name: string

	/**
     * The header field value
     */
    value: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}