// Auto generated code; DO NOT EDIT.

import { NamespaceCondition } from "./NamespaceCondition"

/**
 * NamespaceStatus is information about the current status of a Namespace.
 * 
 */
export declare class NamespaceStatus {
    constructor();
    constructor(spec: NamespaceStatus);

	/**
     * Represents the latest available observations of a namespace's current state.
     * 
     */
    conditions?: Array<NamespaceCondition>

	/**
     * Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
     * 
     */
    phase?: string
}