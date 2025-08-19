// Auto generated code; DO NOT EDIT.

import { IngressLoadBalancerIngress } from "./IngressLoadBalancerIngress"

/**
 * IngressLoadBalancerStatus represents the status of a load-balancer.
 * 
 */
export declare class IngressLoadBalancerStatus {
    constructor();
    constructor(spec: IngressLoadBalancerStatus);

	/**
     * Ingress is a list containing ingress points for the load-balancer.
     * 
     */
    ingress?: Array<IngressLoadBalancerIngress>
}