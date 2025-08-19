// Auto generated code; DO NOT EDIT.

import { Condition } from "../../apimachinery/meta/v1"
import { LoadBalancerStatus } from "./LoadBalancerStatus"

/**
 * ServiceStatus represents the current status of a service.
 * 
 */
export declare class ServiceStatus {
    constructor();
    constructor(spec: ServiceStatus);

	/**
     * Current service state
     * 
     */
    conditions?: Condition

	/**
     * LoadBalancer contains the current status of the load-balancer, if one is present.
     * 
     */
    loadBalancer?: LoadBalancerStatus
}