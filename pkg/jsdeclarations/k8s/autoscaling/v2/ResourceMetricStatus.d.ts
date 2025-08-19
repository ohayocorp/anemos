// Auto generated code; DO NOT EDIT.

import { MetricValueStatus } from "./MetricValueStatus"

/**
 * ResourceMetricStatus indicates the current value of a resource metric known to Kubernetes, as specified in requests and limits, describing each pod in the current scale target (e.g. CPU or memory).  Such metrics are built in to Kubernetes, and have special scaling options on top of those available to normal per-pod metrics using the "pods" source.
 * 
 */
export declare class ResourceMetricStatus {
    constructor();
    constructor(spec: ResourceMetricStatus);

	/**
     * Current contains the current value for the given metric
     * 
     */
    current: MetricValueStatus

	/**
     * Name is the name of the resource in question.
     * 
     */
    name: string
}