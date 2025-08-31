// Auto generated code; DO NOT EDIT.
import { MutatingWebhook } from "./MutatingWebhook"
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import {Document} from '@ohayocorp/anemos';

/**
 * MutatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and may change the object.
 */
export declare class MutatingWebhookConfiguration extends Document {
    constructor();
    constructor(spec: Pick<MutatingWebhookConfiguration, "metadata" | "webhooks">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
     */
    metadata?: ObjectMeta

	/**
     * Webhooks is a list of webhooks and the affected resources and operations.
     */
    webhooks?: Array<MutatingWebhook>
}