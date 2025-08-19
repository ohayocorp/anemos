// Auto generated code; DO NOT EDIT.

import { HorizontalPodAutoscalerCondition } from "./HorizontalPodAutoscalerCondition"
import { MetricStatus } from "./MetricStatus"

/**
 * HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.
 * 
 */
export declare class HorizontalPodAutoscalerStatus {
    constructor();
    constructor(spec: HorizontalPodAutoscalerStatus);

	/**
     * Conditions is the set of conditions required for this autoscaler to scale its target, and indicates whether or not those conditions are met.
     * 
     */
    conditions?: Array<HorizontalPodAutoscalerCondition>

	/**
     * CurrentMetrics is the last read state of the metrics used by this autoscaler.
     * 
     */
    currentMetrics?: Array<MetricStatus>

	/**
     * CurrentReplicas is current number of replicas of pods managed by this autoscaler, as last seen by the autoscaler.
     * 
     */
    currentReplicas?: number

	/**
     * DesiredReplicas is the desired number of replicas of pods managed by this autoscaler, as last calculated by the autoscaler.
     * 
     */
    desiredReplicas: number

	/**
     * LastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods, used by the autoscaler to control how often the number of pods is changed.
     * 
     */
    lastScaleTime?: string

	/**
     * ObservedGeneration is the most recent generation observed by this autoscaler.
     * 
     */
    observedGeneration?: number
}