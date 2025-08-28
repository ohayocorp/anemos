// Auto generated code; DO NOT EDIT.

/**
 * ServiceAccountSubject holds detailed information for service-account-kind subject.
 */
export declare class ServiceAccountSubject {
    constructor();
    constructor(spec: Pick<ServiceAccountSubject, "name" | "namespace">);

	/**
     * `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
     */
    name: string

	/**
     * `namespace` is the namespace of matching ServiceAccount objects. Required.
     */
    namespace: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}