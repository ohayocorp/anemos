// Auto generated code; DO NOT EDIT.



/**
 * Current status of a horizontal pod autoscaler
 * 
 */
export declare class HorizontalPodAutoscalerStatus {
    constructor();
    constructor(spec: HorizontalPodAutoscalerStatus);

	/**
     * CurrentCPUUtilizationPercentage is the current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
     * 
     */
    currentCPUUtilizationPercentage?: number

	/**
     * CurrentReplicas is the current number of replicas of pods managed by this autoscaler.
     * 
     */
    currentReplicas: number

	/**
     * DesiredReplicas is the  desired number of replicas of pods managed by this autoscaler.
     * 
     */
    desiredReplicas: number

	/**
     * LastScaleTime is the last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
     * 
     */
    lastScaleTime?: string

	/**
     * ObservedGeneration is the most recent generation observed by this autoscaler.
     * 
     */
    observedGeneration?: number
}