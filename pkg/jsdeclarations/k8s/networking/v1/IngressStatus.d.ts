// Auto generated code; DO NOT EDIT.

import { IngressLoadBalancerStatus } from "./IngressLoadBalancerStatus"

/**
 * IngressStatus describe the current state of the Ingress.
 * 
 */
export declare class IngressStatus {
    constructor();
    constructor(spec: IngressStatus);

	/**
     * LoadBalancer contains the current status of the load-balancer.
     * 
     */
    loadBalancer?: IngressLoadBalancerStatus
}