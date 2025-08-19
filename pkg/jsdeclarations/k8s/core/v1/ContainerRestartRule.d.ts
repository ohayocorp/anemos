// Auto generated code; DO NOT EDIT.

import { ContainerRestartRuleOnExitCodes } from "./ContainerRestartRuleOnExitCodes"

/**
 * ContainerRestartRule describes how a container exit is handled.
 * 
 */
export declare class ContainerRestartRule {
    constructor();
    constructor(spec: ContainerRestartRule);

	/**
     * Specifies the action taken on a container exit if the requirements are satisfied. The only possible value is "Restart" to restart the container.
     * 
     */
    action: string

	/**
     * Represents the exit codes to check on container exits.
     * 
     */
    exitCodes?: ContainerRestartRuleOnExitCodes
}