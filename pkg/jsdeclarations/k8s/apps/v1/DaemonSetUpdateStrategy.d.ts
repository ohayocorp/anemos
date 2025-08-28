// Auto generated code; DO NOT EDIT.
import { RollingUpdateDaemonSet } from "./RollingUpdateDaemonSet"

/**
 * DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
 */
export declare class DaemonSetUpdateStrategy {
    constructor();
    constructor(spec: Pick<DaemonSetUpdateStrategy, "rollingUpdate" | "type">);

	/**
     * Rolling update config params. Present only if type = "RollingUpdate".
     */
    rollingUpdate?: RollingUpdateDaemonSet

	/**
     * Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
     */
    type?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}