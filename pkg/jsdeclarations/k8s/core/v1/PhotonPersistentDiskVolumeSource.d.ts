// Auto generated code; DO NOT EDIT.



/**
 * Represents a Photon Controller persistent disk resource.
 * 
 */
export declare class PhotonPersistentDiskVolumeSource {
    constructor();
    constructor(spec: PhotonPersistentDiskVolumeSource);

	/**
     * FsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     * 
     */
    fsType?: string

	/**
     * PdID is the ID that identifies Photon Controller persistent disk
     * 
     */
    pdID: string
}