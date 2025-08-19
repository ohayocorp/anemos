// Auto generated code; DO NOT EDIT.



/**
 * MetricTarget defines the target value, average value, or average utilization of a specific metric
 * 
 */
export declare class MetricTarget {
    constructor();
    constructor(spec: MetricTarget);

	/**
     * AverageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type
     * 
     */
    averageUtilization?: number

	/**
     * AverageValue is the target value of the average of the metric across all relevant pods (as a quantity)
     * 
     */
    averageValue?: any

	/**
     * Type represents whether the metric type is Utilization, Value, or AverageValue
     * 
     */
    type: string

	/**
     * Value is the target value of the metric (as a quantity).
     * 
     */
    value?: any
}