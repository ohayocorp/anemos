// Auto generated code; DO NOT EDIT.

import { MetricIdentifier } from "./MetricIdentifier"
import { MetricTarget } from "./MetricTarget"

/**
 * PodsMetricSource indicates how to scale on a metric describing each pod in the current scale target (for example, transactions-processed-per-second). The values will be averaged together before being compared to the target value.
 * 
 */
export declare class PodsMetricSource {
    constructor();
    constructor(spec: PodsMetricSource);

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