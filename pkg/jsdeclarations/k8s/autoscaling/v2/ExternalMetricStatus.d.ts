// Auto generated code; DO NOT EDIT.

import { MetricIdentifier } from "./MetricIdentifier"
import { MetricValueStatus } from "./MetricValueStatus"

/**
 * ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
 * 
 */
export declare class ExternalMetricStatus {
    constructor();
    constructor(spec: ExternalMetricStatus);

	/**
     * Current contains the current value for the given metric
     * 
     */
    current: MetricValueStatus

	/**
     * Metric identifies the target metric by name and selector
     * 
     */
    metric: MetricIdentifier
}