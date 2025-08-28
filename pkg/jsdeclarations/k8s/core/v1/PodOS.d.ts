// Auto generated code; DO NOT EDIT.

/**
 * PodOS defines the OS parameters of a pod.
 */
export declare class PodOS {
    constructor();
    constructor(spec: Pick<PodOS, "name">);

	/**
     * Name is the name of the operating system. The currently supported values are linux and windows. Additional value may be defined in future and can be one of: https://github.com/opencontainers/runtime-spec/blob/master/config.md#platform-specific-configuration Clients should expect to handle additional values and treat unrecognized values in this field as os: null
     */
    name: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}