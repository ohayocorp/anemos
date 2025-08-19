// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { TokenRequestSpec } from "./TokenRequestSpec"
import { TokenRequestStatus } from "./TokenRequestStatus"

/**
 * TokenRequest requests a token for a given service account.
 * 
 */
export declare class TokenRequest {
    constructor();
    constructor(spec: Omit<TokenRequest, "apiVersion" | "kind">);

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
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Spec holds information about the request being evaluated
     * 
     */
    spec: TokenRequestSpec

	/**
     * Status is filled in by the server and indicates whether the token can be authenticated.
     * 
     */
    status?: TokenRequestStatus
}