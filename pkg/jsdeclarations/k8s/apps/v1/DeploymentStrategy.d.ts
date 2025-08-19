// Auto generated code; DO NOT EDIT.

import { RollingUpdateDeployment } from "./RollingUpdateDeployment"

/**
 * DeploymentStrategy describes how to replace existing pods with new ones.
 * 
 */
export declare class DeploymentStrategy {
    constructor();
    constructor(spec: DeploymentStrategy);

	/**
     * Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
     * 
     */
    rollingUpdate?: RollingUpdateDeployment

	/**
     * Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
     * 
     */
    type?: string
}