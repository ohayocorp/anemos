// Auto generated code; DO NOT EDIT.



/**
 * PodResourceClaimStatus is stored in the PodStatus for each PodResourceClaim which references a ResourceClaimTemplate. It stores the generated name for the corresponding ResourceClaim.
 * 
 */
export declare class PodResourceClaimStatus {
    constructor();
    constructor(spec: PodResourceClaimStatus);

	/**
     * Name uniquely identifies this resource claim inside the pod. This must match the name of an entry in pod.spec.resourceClaims, which implies that the string must be a DNS_LABEL.
     * 
     */
    name: string

	/**
     * ResourceClaimName is the name of the ResourceClaim that was generated for the Pod in the namespace of the Pod. If this is unset, then generating a ResourceClaim was not necessary. The pod.spec.resourceClaims entry can be ignored in this case.
     * 
     */
    resourceClaimName?: string
}