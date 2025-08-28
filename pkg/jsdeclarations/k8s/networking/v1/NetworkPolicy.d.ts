// Auto generated code; DO NOT EDIT.
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import { NetworkPolicySpec } from "./NetworkPolicySpec"
import {Document} from '@ohayocorp/anemos';

/**
 * NetworkPolicy describes what network traffic is allowed for a set of Pods
 */
export declare class NetworkPolicy extends Document {
    constructor();
    constructor(spec: Pick<NetworkPolicy, "metadata" | "spec">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta

	/**
     * Spec represents the specification of the desired behavior for this NetworkPolicy.
     */
    spec?: NetworkPolicySpec
}