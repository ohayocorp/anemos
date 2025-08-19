// Auto generated code; DO NOT EDIT.

import { HPAScalingRules } from "./HPAScalingRules"

/**
 * HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
 * 
 */
export declare class HorizontalPodAutoscalerBehavior {
    constructor();
    constructor(spec: HorizontalPodAutoscalerBehavior);

	/**
     * ScaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).
     * 
     */
    scaleDown?: HPAScalingRules

	/**
     * ScaleUp is scaling policy for scaling Up. If not set, the default value is the higher of:
     * 
     *   * increase no more than 4 pods per 60 seconds
     * 
     *   * double the number of pods per 60 seconds
     * 
     * No stabilization is used.
     * 
     */
    scaleUp?: HPAScalingRules
}