// Auto generated code; DO NOT EDIT.
import { RollingUpdateDeployment } from "./RollingUpdateDeployment"

/**
 * DeploymentStrategy describes how to replace existing pods with new ones.
 */
export declare class DeploymentStrategy {
    constructor();
    constructor(spec: Pick<DeploymentStrategy, "rollingUpdate" | "type">);

	/**
     * Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
     */
    rollingUpdate?: RollingUpdateDeployment

	/**
     * Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
     */
    type?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}