// Auto generated code; DO NOT EDIT.
import { ExecAction } from "./ExecAction"
import { HTTPGetAction } from "./HTTPGetAction"
import { SleepAction } from "./SleepAction"
import { TCPSocketAction } from "./TCPSocketAction"

/**
 * LifecycleHandler defines a specific action that should be taken in a lifecycle hook. One and only one of the fields, except TCPSocket must be specified.
 */
export declare class LifecycleHandler {
    constructor();
    constructor(spec: Pick<LifecycleHandler, "exec" | "httpGet" | "sleep" | "tcpSocket">);

	/**
     * Exec specifies a command to execute in the container.
     */
    exec?: ExecAction

	/**
     * HTTPGet specifies an HTTP GET request to perform.
     */
    httpGet?: HTTPGetAction

	/**
     * Sleep represents a duration that the container should sleep.
     */
    sleep?: SleepAction

	/**
     * Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for backward compatibility. There is no validation of this field and lifecycle hooks will fail at runtime when it is specified.
     */
    tcpSocket?: TCPSocketAction

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}