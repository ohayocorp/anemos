// Auto generated code; DO NOT EDIT.

/**
 * Adds and removes POSIX capabilities from running containers.
 */
export declare class Capabilities {
    constructor();
    constructor(spec: Pick<Capabilities, "add" | "drop">);

	/**
     * Added capabilities
     */
    add?: Array<string>

	/**
     * Removed capabilities
     */
    drop?: Array<string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}