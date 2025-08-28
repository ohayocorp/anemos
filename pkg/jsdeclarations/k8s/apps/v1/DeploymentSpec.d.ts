// Auto generated code; DO NOT EDIT.
import { LabelSelector } from "./../../apimachinery/meta/v1"
import { DeploymentStrategy } from "./DeploymentStrategy"
import { PodTemplateSpec } from "./../../core/v1"

/**
 * DeploymentSpec is the specification of the desired behavior of the Deployment.
 */
export declare class DeploymentSpec {
    constructor();
    constructor(spec: Pick<DeploymentSpec, "minReadySeconds" | "paused" | "progressDeadlineSeconds" | "replicas" | "revisionHistoryLimit" | "selector" | "strategy" | "template">);

	/**
     * Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
     */
    minReadySeconds?: number

	/**
     * Indicates that the deployment is paused.
     */
    paused?: boolean

	/**
     * The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.
     */
    progressDeadlineSeconds?: number

	/**
     * Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
     */
    replicas?: number

	/**
     * The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.
     */
    revisionHistoryLimit?: number

	/**
     * Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment. It must match the pod template's labels.
     */
    selector: LabelSelector

	/**
     * The deployment strategy to use to replace existing pods with new ones.
     */
    strategy?: DeploymentStrategy

	/**
     * Template describes the pods that will be created. The only allowed template.spec.restartPolicy value is "Always".
     */
    template: PodTemplateSpec

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}