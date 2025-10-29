// Auto generated code; DO NOT EDIT.

/**
 * NetworkPolicyPort describes a port to allow traffic on
 */
export declare class NetworkPolicyPort {
    constructor();
    constructor(spec: Pick<NetworkPolicyPort, "endPort" | "port" | "protocol">);

	/**
     * EndPort indicates that the range of ports from port to endPort if set, inclusive, should be allowed by the policy. This field cannot be defined if the port field is not defined or if the port field is defined as a named (string) port. The endPort must be equal or greater than port.
     */
    endPort?: number

	/**
     * Port represents the port on the given protocol. This can either be a numerical or named port on a pod. If this field is not provided, this matches all port names and numbers. If present, only traffic on the specified protocol AND port will be matched.
     */
    port?: number | string

	/**
     * Protocol represents the protocol (TCP, UDP, or SCTP) which traffic must match. If not specified, this field defaults to TCP.
     */
    protocol?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}