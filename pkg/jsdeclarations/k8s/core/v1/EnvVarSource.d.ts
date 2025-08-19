// Auto generated code; DO NOT EDIT.

import { ConfigMapKeySelector } from "./ConfigMapKeySelector"
import { FileKeySelector } from "./FileKeySelector"
import { ObjectFieldSelector } from "./ObjectFieldSelector"
import { ResourceFieldSelector } from "./ResourceFieldSelector"
import { SecretKeySelector } from "./SecretKeySelector"

/**
 * EnvVarSource represents a source for the value of an EnvVar.
 * 
 */
export declare class EnvVarSource {
    constructor();
    constructor(spec: EnvVarSource);

	/**
     * Selects a key of a ConfigMap.
     * 
     */
    configMapKeyRef?: ConfigMapKeySelector

	/**
     * Selects a field of the pod: supports metadata.name, metadata.namespace, `metadata.labels['<KEY>']`, `metadata.annotations['<KEY>']`, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.
     * 
     */
    fieldRef?: ObjectFieldSelector

	/**
     * FileKeyRef selects a key of the env file. Requires the EnvFiles feature gate to be enabled.
     * 
     */
    fileKeyRef?: FileKeySelector

	/**
     * Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
     * 
     */
    resourceFieldRef?: ResourceFieldSelector

	/**
     * Selects a key of a secret in the pod's namespace
     * 
     */
    secretKeyRef?: SecretKeySelector
}