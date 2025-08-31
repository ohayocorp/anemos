// Auto generated code; DO NOT EDIT.

/**
 * ResourceClaim references one entry in PodSpec.ResourceClaims.
 */
export declare class ResourceClaim {
    constructor();
    constructor(spec: Pick<ResourceClaim, "name" | "request">);

	/**
     * Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.
     */
    name: string

	/**
     * Request is the name chosen for a request in the referenced claim. If empty, everything from the claim is made available, otherwise only the result of this request.
     */
    request?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}