// Auto generated code; DO NOT EDIT.



/**
 * HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.
 * 
 */
export declare class HostAlias {
    constructor();
    constructor(spec: HostAlias);

	/**
     * Hostnames for the above IP address.
     * 
     */
    hostnames?: Array<string>

	/**
     * IP address of the host file entry.
     * 
     */
    ip: string
}