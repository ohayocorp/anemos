// Auto generated code; DO NOT EDIT.

import { RollingUpdateDaemonSet } from "./RollingUpdateDaemonSet"

/**
 * DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
 * 
 */
export declare class DaemonSetUpdateStrategy {
    constructor();
    constructor(spec: DaemonSetUpdateStrategy);

	/**
     * Rolling update config params. Present only if type = "RollingUpdate".
     * 
     */
    rollingUpdate?: RollingUpdateDaemonSet

	/**
     * Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
     * 
     */
    type?: string
}