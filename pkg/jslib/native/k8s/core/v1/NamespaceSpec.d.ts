// Auto generated code; DO NOT EDIT.

/**
 * NamespaceSpec describes the attributes on a Namespace.
 */
export declare class NamespaceSpec {
    constructor();
    constructor(spec: Pick<NamespaceSpec, "finalizers">);

	/**
     * Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/
     */
    finalizers?: Array<string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}