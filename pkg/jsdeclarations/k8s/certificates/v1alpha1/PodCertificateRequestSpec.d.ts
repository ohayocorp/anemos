// Auto generated code; DO NOT EDIT.



/**
 * PodCertificateRequestSpec describes the certificate request.  All fields are immutable after creation.
 * 
 */
export declare class PodCertificateRequestSpec {
    constructor();
    constructor(spec: PodCertificateRequestSpec);

	/**
     * MaxExpirationSeconds is the maximum lifetime permitted for the certificate.
     * 
     * If omitted, kube-apiserver will set it to 86400(24 hours). kube-apiserver will reject values shorter than 3600 (1 hour).  The maximum allowable value is 7862400 (91 days).
     * 
     * The signer implementation is then free to issue a certificate with any lifetime *shorter* than MaxExpirationSeconds, but no shorter than 3600 seconds (1 hour).  This constraint is enforced by kube-apiserver. `kubernetes.io` signers will never issue certificates with a lifetime longer than 24 hours.
     * 
     */
    maxExpirationSeconds?: number

	/**
     * NodeName is the name of the node the pod is assigned to.
     * 
     */
    nodeName: string

	/**
     * NodeUID is the UID of the node the pod is assigned to.
     * 
     */
    nodeUID: string

	/**
     * PkixPublicKey is the PKIX-serialized public key the signer will issue the certificate to.
     * 
     * The key must be one of RSA3072, RSA4096, ECDSAP256, ECDSAP384, ECDSAP521, or ED25519. Note that this list may be expanded in the future.
     * 
     * Signer implementations do not need to support all key types supported by kube-apiserver and kubelet.  If a signer does not support the key type used for a given PodCertificateRequest, it must deny the request by setting a status.conditions entry with a type of "Denied" and a reason of "UnsupportedKeyType". It may also suggest a key type that it does support in the message field.
     * 
     */
    pkixPublicKey: string

	/**
     * PodName is the name of the pod into which the certificate will be mounted.
     * 
     */
    podName: string

	/**
     * PodUID is the UID of the pod into which the certificate will be mounted.
     * 
     */
    podUID: string

	/**
     * ProofOfPossession proves that the requesting kubelet holds the private key corresponding to pkixPublicKey.
     * 
     * It is contructed by signing the ASCII bytes of the pod's UID using `pkixPublicKey`.
     * 
     * kube-apiserver validates the proof of possession during creation of the PodCertificateRequest.
     * 
     * If the key is an RSA key, then the signature is over the ASCII bytes of the pod UID, using RSASSA-PSS from RFC 8017 (as implemented by the golang function crypto/rsa.SignPSS with nil options).
     * 
     * If the key is an ECDSA key, then the signature is as described by [SEC 1, Version 2.0](https://www.secg.org/sec1-v2.pdf) (as implemented by the golang library function crypto/ecdsa.SignASN1)
     * 
     * If the key is an ED25519 key, the the signature is as described by the [ED25519 Specification](https://ed25519.cr.yp.to/) (as implemented by the golang library crypto/ed25519.Sign).
     * 
     */
    proofOfPossession: string

	/**
     * ServiceAccountName is the name of the service account the pod is running as.
     * 
     */
    serviceAccountName: string

	/**
     * ServiceAccountUID is the UID of the service account the pod is running as.
     * 
     */
    serviceAccountUID: string

	/**
     * SignerName indicates the requested signer.
     * 
     * All signer names beginning with `kubernetes.io` are reserved for use by the Kubernetes project.  There is currently one well-known signer documented by the Kubernetes project, `kubernetes.io/kube-apiserver-client-pod`, which will issue client certificates understood by kube-apiserver.  It is currently unimplemented.
     * 
     */
    signerName: string
}