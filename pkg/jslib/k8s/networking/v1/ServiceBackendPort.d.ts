// Auto generated code; DO NOT EDIT.

/**
 * ServiceBackendPort is the service port being referenced.
 */
export declare class ServiceBackendPort {
    constructor();
    constructor(spec: Pick<ServiceBackendPort, "name" | "number">);

	/**
     * Name is the name of the port on the Service. This is a mutually exclusive setting with "Number".
     */
    name?: string

	/**
     * Number is the numerical port number (e.g. 80) on the Service. This is a mutually exclusive setting with "Name".
     */
    number?: number

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}