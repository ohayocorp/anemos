// Auto generated code; DO NOT EDIT.

import { ClusterTrustBundleProjection } from "./ClusterTrustBundleProjection"
import { ConfigMapProjection } from "./ConfigMapProjection"
import { DownwardAPIProjection } from "./DownwardAPIProjection"
import { PodCertificateProjection } from "./PodCertificateProjection"
import { SecretProjection } from "./SecretProjection"
import { ServiceAccountTokenProjection } from "./ServiceAccountTokenProjection"

/**
 * Projection that may be projected along with other supported volume types. Exactly one of these fields must be set.
 * 
 */
export declare class VolumeProjection {
    constructor();
    constructor(spec: VolumeProjection);

	/**
     * ClusterTrustBundle allows a pod to access the `.spec.trustBundle` field of ClusterTrustBundle objects in an auto-updating file.
     * 
     * Alpha, gated by the ClusterTrustBundleProjection feature gate.
     * 
     * ClusterTrustBundle objects can either be selected by name, or by the combination of signer name and a label selector.
     * 
     * Kubelet performs aggressive normalization of the PEM contents written into the pod filesystem.  Esoteric PEM features such as inter-block comments and block headers are stripped.  Certificates are deduplicated. The ordering of certificates within the file is arbitrary, and Kubelet may change the order over time.
     * 
     */
    clusterTrustBundle?: ClusterTrustBundleProjection

	/**
     * ConfigMap information about the configMap data to project
     * 
     */
    configMap?: ConfigMapProjection

	/**
     * DownwardAPI information about the downwardAPI data to project
     * 
     */
    downwardAPI?: DownwardAPIProjection

	/**
     * Projects an auto-rotating credential bundle (private key and certificate chain) that the pod can use either as a TLS client or server.
     * 
     * Kubelet generates a private key and uses it to send a PodCertificateRequest to the named signer.  Once the signer approves the request and issues a certificate chain, Kubelet writes the key and certificate chain to the pod filesystem.  The pod does not start until certificates have been issued for each podCertificate projected volume source in its spec.
     * 
     * Kubelet will begin trying to rotate the certificate at the time indicated by the signer using the PodCertificateRequest.Status.BeginRefreshAt timestamp.
     * 
     * Kubelet can write a single file, indicated by the credentialBundlePath field, or separate files, indicated by the keyPath and certificateChainPath fields.
     * 
     * The credential bundle is a single file in PEM format.  The first PEM entry is the private key (in PKCS#8 format), and the remaining PEM entries are the certificate chain issued by the signer (typically, signers will return their certificate chain in leaf-to-root order).
     * 
     * Prefer using the credential bundle format, since your application code can read it atomically.  If you use keyPath and certificateChainPath, your application must make two separate file reads. If these coincide with a certificate rotation, it is possible that the private key and leaf certificate you read may not correspond to each other.  Your application will need to check for this condition, and re-read until they are consistent.
     * 
     * The named signer controls chooses the format of the certificate it issues; consult the signer implementation's documentation to learn how to use the certificates it issues.
     * 
     */
    podCertificate?: PodCertificateProjection

	/**
     * Secret information about the secret data to project
     * 
     */
    secret?: SecretProjection

	/**
     * ServiceAccountToken is information about the serviceAccountToken data to project
     * 
     */
    serviceAccountToken?: ServiceAccountTokenProjection
}