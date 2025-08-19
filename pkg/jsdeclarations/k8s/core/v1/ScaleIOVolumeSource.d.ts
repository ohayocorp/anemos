// Auto generated code; DO NOT EDIT.

import { LocalObjectReference } from "./LocalObjectReference"

/**
 * ScaleIOVolumeSource represents a persistent ScaleIO volume
 * 
 */
export declare class ScaleIOVolumeSource {
    constructor();
    constructor(spec: ScaleIOVolumeSource);

	/**
     * FsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Default is "xfs".
     * 
     */
    fsType?: string

	/**
     * Gateway is the host address of the ScaleIO API Gateway.
     * 
     */
    gateway: string

	/**
     * ProtectionDomain is the name of the ScaleIO Protection Domain for the configured storage.
     * 
     */
    protectionDomain?: string

	/**
     * ReadOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     * 
     */
    readOnly?: boolean

	/**
     * SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.
     * 
     */
    secretRef: LocalObjectReference

	/**
     * SslEnabled Flag enable/disable SSL communication with Gateway, default false
     * 
     */
    sslEnabled?: boolean

	/**
     * StorageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.
     * 
     */
    storageMode?: string

	/**
     * StoragePool is the ScaleIO Storage Pool associated with the protection domain.
     * 
     */
    storagePool?: string

	/**
     * System is the name of the storage system as configured in ScaleIO.
     * 
     */
    system: string

	/**
     * VolumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.
     * 
     */
    volumeName?: string
}