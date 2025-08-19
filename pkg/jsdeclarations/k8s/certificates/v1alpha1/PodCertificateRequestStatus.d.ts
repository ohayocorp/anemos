// Auto generated code; DO NOT EDIT.

import { Condition } from "../../apimachinery/meta/v1"

/**
 * PodCertificateRequestStatus describes the status of the request, and holds the certificate data if the request is issued.
 * 
 */
export declare class PodCertificateRequestStatus {
    constructor();
    constructor(spec: PodCertificateRequestStatus);

	/**
     * BeginRefreshAt is the time at which the kubelet should begin trying to refresh the certificate.  This field is set via the /status subresource, and must be set at the same time as certificateChain.  Once populated, this field is immutable.
     * 
     * This field is only a hint.  Kubelet may start refreshing before or after this time if necessary.
     * 
     */
    beginRefreshAt?: string

	/**
     * CertificateChain is populated with an issued certificate by the signer. This field is set via the /status subresource. Once populated, this field is immutable.
     * 
     * If the certificate signing request is denied, a condition of type "Denied" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type "Failed" is added and this field remains empty.
     * 
     * Validation requirements:
     * 
     *  1. certificateChain must consist of one or more PEM-formatted certificates.
     * 
     *  2. Each entry must be a valid PEM-wrapped, DER-encoded ASN.1 Certificate as
     * 
     *     described in section 4 of RFC5280.
     * 
     * If more than one block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.  When projecting the chain into a pod volume, kubelet will drop any data in-between the PEM blocks, as well as any PEM block headers.
     * 
     */
    certificateChain?: string

	/**
     * Conditions applied to the request.
     * 
     * The types "Issued", "Denied", and "Failed" have special handling.  At most one of these conditions may be present, and they must have status "True".
     * 
     * If the request is denied with `Reason=UnsupportedKeyType`, the signer may suggest a key type that will work in the message field.
     * 
     */
    conditions?: Condition

	/**
     * NotAfter is the time at which the certificate expires.  The value must be the same as the notAfter value in the leaf certificate in certificateChain.  This field is set via the /status subresource.  Once populated, it is immutable.  The signer must set this field at the same time it sets certificateChain.
     * 
     */
    notAfter?: string

	/**
     * NotBefore is the time at which the certificate becomes valid.  The value must be the same as the notBefore value in the leaf certificate in certificateChain.  This field is set via the /status subresource.  Once populated, it is immutable. The signer must set this field at the same time it sets certificateChain.
     * 
     */
    notBefore?: string
}