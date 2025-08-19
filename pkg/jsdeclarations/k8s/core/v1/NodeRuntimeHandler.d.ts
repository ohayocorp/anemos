// Auto generated code; DO NOT EDIT.

import { NodeRuntimeHandlerFeatures } from "./NodeRuntimeHandlerFeatures"

/**
 * NodeRuntimeHandler is a set of runtime handler information.
 * 
 */
export declare class NodeRuntimeHandler {
    constructor();
    constructor(spec: NodeRuntimeHandler);

	/**
     * Supported features.
     * 
     */
    features?: NodeRuntimeHandlerFeatures

	/**
     * Runtime handler name. Empty for the default runtime handler.
     * 
     */
    name?: string
}