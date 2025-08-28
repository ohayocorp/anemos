// Auto generated code; DO NOT EDIT.
import { ObjectReference } from "./ObjectReference"

/**
 * EndpointAddress is a tuple that describes single IP address. Deprecated: This API is deprecated in v1.33+.
 */
export declare class EndpointAddress {
    constructor();
    constructor(spec: Pick<EndpointAddress, "hostname" | "ip" | "nodeName" | "targetRef">);

	/**
     * The Hostname of this endpoint
     */
    hostname?: string

	/**
     * The IP of this endpoint. May not be loopback (127.0.0.0/8 or ::1), link-local (169.254.0.0/16 or fe80::/10), or link-local multicast (224.0.0.0/24 or ff02::/16).
     */
    ip: string

	/**
     * Optional: Node hosting this endpoint. This can be used to determine endpoints local to a node.
     */
    nodeName?: string

	/**
     * Reference to object providing the endpoint.
     */
    targetRef?: ObjectReference

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}