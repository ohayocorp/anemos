// Auto generated code; DO NOT EDIT.

/**
 * ServiceReference holds a reference to Service.legacy.k8s.io
 */
export declare class ServiceReference {
    constructor();
    constructor(spec: Pick<ServiceReference, "name" | "namespace" | "port">);

	/**
     * Name is the name of the service
     */
    name?: string

	/**
     * Namespace is the namespace of the service
     */
    namespace?: string

	/**
     * If specified, the port on the service that hosting webhook. Default to 443 for backward compatibility. `port` should be a valid port number (1-65535, inclusive).
     */
    port?: number

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}