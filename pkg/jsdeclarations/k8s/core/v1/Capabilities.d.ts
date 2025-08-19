// Auto generated code; DO NOT EDIT.



/**
 * Adds and removes POSIX capabilities from running containers.
 * 
 */
export declare class Capabilities {
    constructor();
    constructor(spec: Capabilities);

	/**
     * Added capabilities
     * 
     */
    add?: Array<string>

	/**
     * Removed capabilities
     * 
     */
    drop?: Array<string>
}