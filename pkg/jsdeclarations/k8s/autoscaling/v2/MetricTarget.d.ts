// Auto generated code; DO NOT EDIT.

/**
 * MetricTarget defines the target value, average value, or average utilization of a specific metric
 */
export declare class MetricTarget {
    constructor();
    constructor(spec: Pick<MetricTarget, "averageUtilization" | "averageValue" | "type" | "value">);

	/**
     * AverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type
     */
    averageUtilization?: number

	/**
     * AverageValue is the target value of the average of the metric across all relevant pods (as a quantity)
     */
    averageValue?: number | string

	/**
     * Type represents whether the metric type is Utilization, Value, or AverageValue
     */
    type: string

	/**
     * Value is the target value of the metric (as a quantity).
     */
    value?: number | string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}