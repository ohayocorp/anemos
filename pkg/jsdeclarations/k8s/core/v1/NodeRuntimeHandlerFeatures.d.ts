// Auto generated code; DO NOT EDIT.



/**
 * NodeRuntimeHandlerFeatures is a set of features implemented by the runtime handler.
 * 
 */
export declare class NodeRuntimeHandlerFeatures {
    constructor();
    constructor(spec: NodeRuntimeHandlerFeatures);

	/**
     * RecursiveReadOnlyMounts is set to true if the runtime handler supports RecursiveReadOnlyMounts.
     * 
     */
    recursiveReadOnlyMounts?: boolean

	/**
     * UserNamespaces is set to true if the runtime handler supports UserNamespaces, including for volumes.
     * 
     */
    userNamespaces?: boolean
}