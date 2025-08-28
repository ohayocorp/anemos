// Auto generated code; DO NOT EDIT.
import { WebhookClientConfig } from "./WebhookClientConfig"

/**
 * WebhookConversion describes how to call a conversion webhook
 */
export declare class WebhookConversion {
    constructor();
    constructor(spec: Pick<WebhookConversion, "clientConfig" | "conversionReviewVersions">);

	/**
     * ClientConfig is the instructions for how to call the webhook if strategy is `Webhook`.
     */
    clientConfig?: WebhookClientConfig

	/**
     * ConversionReviewVersions is an ordered list of preferred `ConversionReview` versions the Webhook expects. The API server will use the first version in the list which it supports. If none of the versions specified in this list are supported by API server, conversion will fail for the custom resource. If a persisted Webhook configuration specifies allowed versions and does not include any versions known to the API Server, calls to the webhook will fail.
     */
    conversionReviewVersions: Array<string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}