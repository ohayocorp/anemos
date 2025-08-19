// Auto generated code; DO NOT EDIT.



/**
 * NamespaceCondition contains details about state of namespace.
 * 
 */
export declare class NamespaceCondition {
    constructor();
    constructor(spec: NamespaceCondition);

	/**
     * Last time the condition transitioned from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * Human-readable message indicating details about last transition.
     * 
     */
    message?: string

	/**
     * Unique, one-word, CamelCase reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * Type of namespace controller condition.
     * 
     */
    type: string
}