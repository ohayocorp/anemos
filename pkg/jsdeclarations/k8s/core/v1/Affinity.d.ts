// Auto generated code; DO NOT EDIT.

import { NodeAffinity } from "./NodeAffinity"
import { PodAffinity } from "./PodAffinity"
import { PodAntiAffinity } from "./PodAntiAffinity"

/**
 * Affinity is a group of affinity scheduling rules.
 * 
 */
export declare class Affinity {
    constructor();
    constructor(spec: Affinity);

	/**
     * Describes node affinity scheduling rules for the pod.
     * 
     */
    nodeAffinity?: NodeAffinity

	/**
     * Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).
     * 
     */
    podAffinity?: PodAffinity

	/**
     * Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).
     * 
     */
    podAntiAffinity?: PodAntiAffinity
}