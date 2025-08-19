// Auto generated code; DO NOT EDIT.



/**
 * NodeFeatures describes the set of features implemented by the CRI implementation. The features contained in the NodeFeatures should depend only on the cri implementation independent of runtime handlers.
 * 
 */
export declare class NodeFeatures {
    constructor();
    constructor(spec: NodeFeatures);

	/**
     * SupplementalGroupsPolicy is set to true if the runtime supports SupplementalGroupsPolicy and ContainerUser.
     * 
     */
    supplementalGroupsPolicy?: boolean
}