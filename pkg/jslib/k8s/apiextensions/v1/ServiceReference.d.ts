// Auto generated code; DO NOT EDIT.

/**
 * ServiceReference holds a reference to Service.legacy.k8s.io
 */
export declare class ServiceReference {
    constructor();
    constructor(spec: Pick<ServiceReference, "name" | "namespace" | "path" | "port">);

	/**
     * Name is the name of the service. Required
     */
    name: string

	/**
     * Namespace is the namespace of the service. Required
     */
    namespace: string

	/**
     * Path is an optional URL path at which the webhook will be contacted.
     */
    path?: string

	/**
     * Port is an optional service port at which the webhook will be contacted. `port` should be a valid port number (1-65535, inclusive). Defaults to 443 for backward compatibility.
     */
    port?: number

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}