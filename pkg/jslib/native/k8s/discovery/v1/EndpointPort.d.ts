// Auto generated code; DO NOT EDIT.

/**
 * EndpointPort represents a Port used by an EndpointSlice
 */
export declare class EndpointPort {
    constructor();
    constructor(spec: Pick<EndpointPort, "appProtocol" | "name" | "port" | "protocol">);

	/**
     * The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:
    
     * * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).
    
     * * Kubernetes-defined prefixed names:
    
     *   * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-
    
     *   * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
    
     *   * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455
    
     * * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
     */
    appProtocol?: string

	/**
     * Name represents the name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is derived from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
     */
    name?: string

	/**
     * Port represents the port number of the endpoint. If the EndpointSlice is derived from a Kubernetes service, this must be set to the service's target port. EndpointSlices used for other purposes may have a nil port.
     */
    port?: number

	/**
     * Protocol represents the IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
     */
    protocol?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}