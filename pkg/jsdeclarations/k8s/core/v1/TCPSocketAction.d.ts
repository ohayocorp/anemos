// Auto generated code; DO NOT EDIT.

/**
 * TCPSocketAction describes an action based on opening a socket
 */
export declare class TCPSocketAction {
    constructor();
    constructor(spec: Pick<TCPSocketAction, "host" | "port">);

	/**
     * Optional: Host name to connect to, defaults to the pod IP.
     */
    host?: string

	/**
     * Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
     */
    port: number | string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}