// Auto generated code; DO NOT EDIT.



/**
 * MetricValueStatus holds the current value for a metric
 * 
 */
export declare class MetricValueStatus {
    constructor();
    constructor(spec: MetricValueStatus);

	/**
     * CurrentAverageUtilization is the current value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods.
     * 
     */
    averageUtilization?: number

	/**
     * AverageValue is the current value of the average of the metric across all relevant pods (as a quantity)
     * 
     */
    averageValue?: any

	/**
     * Value is the current value of the metric (as a quantity).
     * 
     */
    value?: any
}