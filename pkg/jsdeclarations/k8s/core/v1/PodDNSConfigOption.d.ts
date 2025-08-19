// Auto generated code; DO NOT EDIT.



/**
 * PodDNSConfigOption defines DNS resolver options of a pod.
 * 
 */
export declare class PodDNSConfigOption {
    constructor();
    constructor(spec: PodDNSConfigOption);

	/**
     * Name is this DNS resolver option's name. Required.
     * 
     */
    name?: string

	/**
     * Value is this DNS resolver option's value.
     * 
     */
    value?: string
}