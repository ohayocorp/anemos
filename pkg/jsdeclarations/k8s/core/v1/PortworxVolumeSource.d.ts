// Auto generated code; DO NOT EDIT.



/**
 * PortworxVolumeSource represents a Portworx volume resource.
 * 
 */
export declare class PortworxVolumeSource {
    constructor();
    constructor(spec: PortworxVolumeSource);

	/**
     * FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs". Implicitly inferred to be "ext4" if unspecified.
     * 
     */
    fsType?: string

	/**
     * ReadOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     * 
     */
    readOnly?: boolean

	/**
     * VolumeID uniquely identifies a Portworx volume
     * 
     */
    volumeID: string
}