// Auto generated code; DO NOT EDIT.

import { ValidatingWebhook } from "./ValidatingWebhook"
import { ObjectMeta } from "../../apimachinery/meta/v1"

/**
 * ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.
 * 
 */
export declare class ValidatingWebhookConfiguration {
    constructor();
    constructor(spec: Omit<ValidatingWebhookConfiguration, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     * 
     */
    metadata?: ObjectMeta

	/**
     * Webhooks is a list of webhooks and the affected resources and operations.
     * 
     */
    webhooks?: Array<ValidatingWebhook>
}