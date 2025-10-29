// Auto generated code; DO NOT EDIT.
import { CrossVersionObjectReference } from "./CrossVersionObjectReference"
import { MetricIdentifier } from "./MetricIdentifier"
import { MetricTarget } from "./MetricTarget"

/**
 * ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
 */
export declare class ObjectMetricSource {
    constructor();
    constructor(spec: Pick<ObjectMetricSource, "describedObject" | "metric" | "target">);

	/**
     * DescribedObject specifies the descriptions of a object,such as kind,name apiVersion
     */
    describedObject: CrossVersionObjectReference

	/**
     * Metric identifies the target metric by name and selector
     */
    metric: MetricIdentifier

	/**
     * Target specifies the target value for the given metric
     */
    target: MetricTarget

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}