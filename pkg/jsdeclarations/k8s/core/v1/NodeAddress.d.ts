// Auto generated code; DO NOT EDIT.



/**
 * NodeAddress contains information for the node's address.
 * 
 */
export declare class NodeAddress {
    constructor();
    constructor(spec: NodeAddress);

	/**
     * The node address.
     * 
     */
    address: string

	/**
     * Node address type, one of Hostname, ExternalIP or InternalIP.
     * 
     */
    type: string
}