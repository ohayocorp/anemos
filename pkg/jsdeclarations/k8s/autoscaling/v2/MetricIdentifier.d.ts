// Auto generated code; DO NOT EDIT.

import { LabelSelector } from "../../apimachinery/meta/v1"

/**
 * MetricIdentifier defines the name and optionally selector for a metric
 * 
 */
export declare class MetricIdentifier {
    constructor();
    constructor(spec: MetricIdentifier);

	/**
     * Name is the name of the given metric
     * 
     */
    name: string

	/**
     * Selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
     * 
     */
    selector?: LabelSelector
}