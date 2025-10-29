// Auto generated code; DO NOT EDIT.

/**
 * ResourceFieldSelector represents container resources (cpu, memory) and their output format
 */
export declare class ResourceFieldSelector {
    constructor();
    constructor(spec: Pick<ResourceFieldSelector, "containerName" | "divisor" | "resource">);

	/**
     * Container name: required for volumes, optional for env vars
     */
    containerName?: string

	/**
     * Specifies the output format of the exposed resources, defaults to "1"
     */
    divisor?: number | string

	/**
     * Required: resource to select
     */
    resource: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}