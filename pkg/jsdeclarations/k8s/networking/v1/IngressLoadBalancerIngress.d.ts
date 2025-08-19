// Auto generated code; DO NOT EDIT.

import { IngressPortStatus } from "./IngressPortStatus"

/**
 * IngressLoadBalancerIngress represents the status of a load-balancer ingress point.
 * 
 */
export declare class IngressLoadBalancerIngress {
    constructor();
    constructor(spec: IngressLoadBalancerIngress);

	/**
     * Hostname is set for load-balancer ingress points that are DNS based.
     * 
     */
    hostname?: string

	/**
     * Ip is set for load-balancer ingress points that are IP based.
     * 
     */
    ip?: string

	/**
     * Ports provides information about the ports exposed by this LoadBalancer.
     * 
     */
    ports?: Array<IngressPortStatus>
}