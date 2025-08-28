// Auto generated code; DO NOT EDIT.

/**
 * ContainerPort represents a network port in a single container.
 */
export declare class ContainerPort {
    constructor();
    constructor(spec: Pick<ContainerPort, "containerPort" | "hostIP" | "hostPort" | "name" | "protocol">);

	/**
     * Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
     */
    containerPort: number

	/**
     * What host IP to bind the external port to.
     */
    hostIP?: string

	/**
     * Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
     */
    hostPort?: number

	/**
     * If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
     */
    name?: string

	/**
     * Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
     */
    protocol?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}