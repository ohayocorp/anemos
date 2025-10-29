// Auto generated code; DO NOT EDIT.
import { CustomResourceSubresourceScale } from "./CustomResourceSubresourceScale"
import { CustomResourceSubresourceStatus } from "./CustomResourceSubresourceStatus"

/**
 * CustomResourceSubresources defines the status and scale subresources for CustomResources.
 */
export declare class CustomResourceSubresources {
    constructor();
    constructor(spec: Pick<CustomResourceSubresources, "scale" | "status">);

	/**
     * Scale indicates the custom resource should serve a `/scale` subresource that returns an `autoscaling/v1` Scale object.
     */
    scale?: CustomResourceSubresourceScale

	/**
     * Status indicates the custom resource should serve a `/status` subresource. When enabled: 1. requests to the custom resource primary endpoint ignore changes to the `status` stanza of the object. 2. requests to the custom resource `/status` subresource ignore changes to anything other than the `status` stanza of the object.
     */
    status?: CustomResourceSubresourceStatus

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}