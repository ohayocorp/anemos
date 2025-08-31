// Auto generated code; DO NOT EDIT.
import { Toleration } from "./../../core/v1"

/**
 * Scheduling specifies the scheduling constraints for nodes supporting a RuntimeClass.
 */
export declare class Scheduling {
    constructor();
    constructor(spec: Pick<Scheduling, "nodeSelector" | "tolerations">);

	/**
     * NodeSelector lists labels that must be present on nodes that support this RuntimeClass. Pods using this RuntimeClass can only be scheduled to a node matched by this selector. The RuntimeClass nodeSelector is merged with a pod's existing nodeSelector. Any conflicts will cause the pod to be rejected in admission.
     */
    nodeSelector?: Record<string, string>

	/**
     * Tolerations are appended (excluding duplicates) to pods running with this RuntimeClass during admission, effectively unioning the set of nodes tolerated by the pod and the RuntimeClass.
     */
    tolerations?: Toleration

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}