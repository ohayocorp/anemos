// Auto generated code; DO NOT EDIT.

import { Condition } from "../../apimachinery/meta/v1"

/**
 * ServiceCIDRStatus describes the current state of the ServiceCIDR.
 * 
 */
export declare class ServiceCIDRStatus {
    constructor();
    constructor(spec: ServiceCIDRStatus);

	/**
     * Conditions holds an array of metav1.Condition that describe the state of the ServiceCIDR. Current service state
     * 
     */
    conditions?: Condition
}