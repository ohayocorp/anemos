// Auto generated code; DO NOT EDIT.

import { ConfigMapEnvSource } from "./ConfigMapEnvSource"
import { SecretEnvSource } from "./SecretEnvSource"

/**
 * EnvFromSource represents the source of a set of ConfigMaps or Secrets
 * 
 */
export declare class EnvFromSource {
    constructor();
    constructor(spec: EnvFromSource);

	/**
     * The ConfigMap to select from
     * 
     */
    configMapRef?: ConfigMapEnvSource

	/**
     * Optional text to prepend to the name of each environment variable. May consist of any printable ASCII characters except '='.
     * 
     */
    prefix?: string

	/**
     * The Secret to select from
     * 
     */
    secretRef?: SecretEnvSource
}