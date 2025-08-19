// Auto generated code; DO NOT EDIT.

import { MutatingWebhookConfiguration } from "./MutatingWebhookConfiguration"
import { ListMeta } from "../../apimachinery/meta/v1"

/**
 * MutatingWebhookConfigurationList is a list of MutatingWebhookConfiguration.
 * 
 */
export declare class MutatingWebhookConfigurationList {
    constructor();
    constructor(spec: Omit<MutatingWebhookConfigurationList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * List of MutatingWebhookConfiguration.
     * 
     */
    items: Array<MutatingWebhookConfiguration>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    metadata?: ListMeta
}