// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { PodCertificateRequestSpec } from "./PodCertificateRequestSpec"

/**
 * PodCertificateRequest encodes a pod requesting a certificate from a given signer.
 * 
 * Kubelets use this API to implement podCertificate projected volumes
 * 
 */
export declare class PodCertificateRequest {
    constructor();
    constructor(spec: Omit<PodCertificateRequest, "apiVersion" | "kind">);

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
     * Metadata contains the object metadata.
     * 
     */
    metadata?: ObjectMeta

	/**
     * Spec contains the details about the certificate being requested.
     * 
     */
    spec: PodCertificateRequestSpec
}