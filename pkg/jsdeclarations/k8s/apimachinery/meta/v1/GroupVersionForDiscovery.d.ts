// Auto generated code; DO NOT EDIT.



/**
 * GroupVersion contains the "group/version" and "version" string of a version. It is made a struct to keep extensibility.
 * 
 */
export declare class GroupVersionForDiscovery {
    constructor();
    constructor(spec: GroupVersionForDiscovery);

	/**
     * GroupVersion specifies the API group and version in the form "group/version"
     * 
     */
    groupVersion: string

	/**
     * Version specifies the version in the form of "version". This is to save the clients the trouble of splitting the GroupVersion.
     * 
     */
    version: string
}