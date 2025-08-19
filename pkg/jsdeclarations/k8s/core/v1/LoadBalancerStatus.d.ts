// Auto generated code; DO NOT EDIT.

import { LoadBalancerIngress } from "./LoadBalancerIngress"

/**
 * LoadBalancerStatus represents the status of a load-balancer.
 * 
 */
export declare class LoadBalancerStatus {
    constructor();
    constructor(spec: LoadBalancerStatus);

	/**
     * Ingress is a list containing ingress points for the load-balancer. Traffic intended for the service should be sent to these ingress points.
     * 
     */
    ingress?: Array<LoadBalancerIngress>
}