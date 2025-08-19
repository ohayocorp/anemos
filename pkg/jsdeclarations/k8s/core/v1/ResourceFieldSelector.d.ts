// Auto generated code; DO NOT EDIT.



/**
 * ResourceFieldSelector represents container resources (cpu, memory) and their output format
 * 
 */
export declare class ResourceFieldSelector {
    constructor();
    constructor(spec: ResourceFieldSelector);

	/**
     * Container name: required for volumes, optional for env vars
     * 
     */
    containerName?: string

	/**
     * Specifies the output format of the exposed resources, defaults to "1"
     * 
     */
    divisor?: any

	/**
     * Required: resource to select
     * 
     */
    resource: string
}