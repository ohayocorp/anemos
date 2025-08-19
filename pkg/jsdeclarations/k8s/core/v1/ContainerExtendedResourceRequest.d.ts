// Auto generated code; DO NOT EDIT.



/**
 * ContainerExtendedResourceRequest has the mapping of container name, extended resource name to the device request name.
 * 
 */
export declare class ContainerExtendedResourceRequest {
    constructor();
    constructor(spec: ContainerExtendedResourceRequest);

	/**
     * The name of the container requesting resources.
     * 
     */
    containerName: string

	/**
     * The name of the request in the special ResourceClaim which corresponds to the extended resource.
     * 
     */
    requestName: string

	/**
     * The name of the extended resource in that container which gets backed by DRA.
     * 
     */
    resourceName: string
}