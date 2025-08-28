// Auto generated code; DO NOT EDIT.
import { TypedLocalObjectReference } from "./../../core/v1"
import { IngressServiceBackend } from "./IngressServiceBackend"

/**
 * IngressBackend describes all endpoints for a given service and port.
 */
export declare class IngressBackend {
    constructor();
    constructor(spec: Pick<IngressBackend, "resource" | "service">);

	/**
     * Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, a service.Name and service.Port must not be specified. This is a mutually exclusive setting with "Service".
     */
    resource?: TypedLocalObjectReference

	/**
     * Service references a service as a backend. This is a mutually exclusive setting with "Resource".
     */
    service?: IngressServiceBackend

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}