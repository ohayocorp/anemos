// Auto generated code; DO NOT EDIT.



/**
 * Represents a vSphere volume resource.
 * 
 */
export declare class VsphereVirtualDiskVolumeSource {
    constructor();
    constructor(spec: VsphereVirtualDiskVolumeSource);

	/**
     * FsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     * 
     */
    fsType?: string

	/**
     * StoragePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.
     * 
     */
    storagePolicyID?: string

	/**
     * StoragePolicyName is the storage Policy Based Management (SPBM) profile name.
     * 
     */
    storagePolicyName?: string

	/**
     * VolumePath is the path that identifies vSphere volume vmdk
     * 
     */
    volumePath: string
}