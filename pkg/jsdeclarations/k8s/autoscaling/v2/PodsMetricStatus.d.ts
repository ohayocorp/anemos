// Auto generated code; DO NOT EDIT.

import { MetricIdentifier } from "./MetricIdentifier"
import { MetricValueStatus } from "./MetricValueStatus"

/**
 * PodsMetricStatus indicates the current value of a metric describing each pod in the current scale target (for example, transactions-processed-per-second).
 * 
 */
export declare class PodsMetricStatus {
    constructor();
    constructor(spec: PodsMetricStatus);

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