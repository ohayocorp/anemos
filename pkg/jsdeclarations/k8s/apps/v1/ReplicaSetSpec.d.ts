// Auto generated code; DO NOT EDIT.

import { LabelSelector } from "../../apimachinery/meta/v1"
import { PodTemplateSpec } from "../../core/v1"

/**
 * ReplicaSetSpec is the specification of a ReplicaSet.
 * 
 */
export declare class ReplicaSetSpec {
    constructor();
    constructor(spec: ReplicaSetSpec);

	/**
     * Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
     * 
     */
    minReadySeconds?: number

	/**
     * Replicas is the number of desired pods. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicaset
     * 
     */
    replicas?: number

	/**
     * Selector is a label query over pods that should match the replica count. Label keys and values that must match in order to be controlled by this replica set. It must match the pod template's labels. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#label-selectors
     * 
     */
    selector: LabelSelector

	/**
     * Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: https://kubernetes.io/docs/concepts/workloads/controllers/replicaset/#pod-template
     * 
     */
    template?: PodTemplateSpec
}