// Auto generated code; DO NOT EDIT.

/**
 * Sysctl defines a kernel parameter to be set
 */
export declare class Sysctl {
    constructor();
    constructor(spec: Pick<Sysctl, "name" | "value">);

	/**
     * Name of a property to set
     */
    name: string

	/**
     * Value of a property to set
     */
    value: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}