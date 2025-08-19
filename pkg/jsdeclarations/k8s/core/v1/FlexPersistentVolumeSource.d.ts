// Auto generated code; DO NOT EDIT.

import { SecretReference } from "./SecretReference"

/**
 * FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.
 * 
 */
export declare class FlexPersistentVolumeSource {
    constructor();
    constructor(spec: FlexPersistentVolumeSource);

	/**
     * Driver is the name of the driver to use for this volume.
     * 
     */
    driver: string

	/**
     * FsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default filesystem depends on FlexVolume script.
     * 
     */
    fsType?: string

	/**
     * Options is Optional: this field holds extra command options if any.
     * 
     */
    options?: Record<string, string>

	/**
     * ReadOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     * 
     */
    readOnly?: boolean

	/**
     * SecretRef is Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.
     * 
     */
    secretRef?: SecretReference
}