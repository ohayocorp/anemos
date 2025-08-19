// Auto generated code; DO NOT EDIT.

import { DaemonEndpoint } from "./DaemonEndpoint"

/**
 * NodeDaemonEndpoints lists ports opened by daemons running on the Node.
 * 
 */
export declare class NodeDaemonEndpoints {
    constructor();
    constructor(spec: NodeDaemonEndpoints);

	/**
     * Endpoint on which Kubelet is listening.
     * 
     */
    kubeletEndpoint?: DaemonEndpoint
}