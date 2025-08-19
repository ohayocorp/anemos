// Auto generated code; DO NOT EDIT.

import { PodAffinityTerm } from "./PodAffinityTerm"

/**
 * The weights of all of the matched WeightedPodAffinityTerm fields are added per-node to find the most preferred node(s)
 * 
 */
export declare class WeightedPodAffinityTerm {
    constructor();
    constructor(spec: WeightedPodAffinityTerm);

	/**
     * Required. A pod affinity term, associated with the corresponding weight.
     * 
     */
    podAffinityTerm: PodAffinityTerm

	/**
     * Weight associated with matching the corresponding podAffinityTerm, in the range 1-100.
     * 
     */
    weight: number
}