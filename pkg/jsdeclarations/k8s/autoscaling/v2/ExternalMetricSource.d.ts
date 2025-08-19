// Auto generated code; DO NOT EDIT.

import { MetricIdentifier } from "./MetricIdentifier"
import { MetricTarget } from "./MetricTarget"

/**
 * ExternalMetricSource indicates how to scale on a metric not associated with any Kubernetes object (for example length of queue in cloud messaging service, or QPS from loadbalancer running outside of cluster).
 * 
 */
export declare class ExternalMetricSource {
    constructor();
    constructor(spec: ExternalMetricSource);

	/**
     * Metric identifies the target metric by name and selector
     * 
     */
    metric: MetricIdentifier

	/**
     * Target specifies the target value for the given metric
     * 
     */
    target: MetricTarget
}