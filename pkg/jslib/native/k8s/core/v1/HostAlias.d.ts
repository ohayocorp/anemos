// Auto generated code; DO NOT EDIT.

/**
 * HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
 */
export declare class HostAlias {
    constructor();
    constructor(spec: Pick<HostAlias, "hostnames" | "ip">);

	/**
     * Hostnames for the above IP address.
     */
    hostnames?: Array<string>

	/**
     * IP address of the host file entry.
     */
    ip: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}