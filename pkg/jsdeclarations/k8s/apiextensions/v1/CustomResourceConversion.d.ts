// Auto generated code; DO NOT EDIT.
import { WebhookConversion } from "./WebhookConversion"

/**
 * CustomResourceConversion describes how to convert different versions of a CR.
 */
export declare class CustomResourceConversion {
    constructor();
    constructor(spec: Pick<CustomResourceConversion, "strategy" | "webhook">);

	/**
     * Strategy specifies how custom resources are converted between versions. Allowed values are: - `"None"`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `"Webhook"`: API Server will call to an external webhook to do the conversion. Additional information
    
     *   is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
     */
    strategy: string

	/**
     * Webhook describes how to call the conversion webhook. Required when `strategy` is set to `"Webhook"`.
     */
    webhook?: WebhookConversion

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}