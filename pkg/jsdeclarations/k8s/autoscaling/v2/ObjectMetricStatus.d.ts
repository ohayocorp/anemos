// Auto generated code; DO NOT EDIT.

import { CrossVersionObjectReference } from "./CrossVersionObjectReference"
import { MetricIdentifier } from "./MetricIdentifier"
import { MetricValueStatus } from "./MetricValueStatus"

/**
 * ObjectMetricStatus indicates the current value of a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
 * 
 */
export declare class ObjectMetricStatus {
    constructor();
    constructor(spec: ObjectMetricStatus);

	/**
     * Current contains the current value for the given metric
     * 
     */
    current: MetricValueStatus

	/**
     * DescribedObject specifies the descriptions of a object,such as kind,name apiVersion
     * 
     */
    describedObject: CrossVersionObjectReference

	/**
     * Metric identifies the target metric by name and selector
     * 
     */
    metric: MetricIdentifier
}