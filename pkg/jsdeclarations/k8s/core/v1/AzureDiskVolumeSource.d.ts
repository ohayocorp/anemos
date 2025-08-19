// Auto generated code; DO NOT EDIT.



/**
 * AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
 * 
 */
export declare class AzureDiskVolumeSource {
    constructor();
    constructor(spec: AzureDiskVolumeSource);

	/**
     * CachingMode is the Host Caching mode: None, Read Only, Read Write.
     * 
     */
    cachingMode?: string

	/**
     * DiskName is the Name of the data disk in the blob storage
     * 
     */
    diskName: string

	/**
     * DiskURI is the URI of data disk in the blob storage
     * 
     */
    diskURI: string

	/**
     * FsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     * 
     */
    fsType?: string

	/**
     * Kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared
     * 
     */
    kind?: string

	/**
     * ReadOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     * 
     */
    readOnly?: boolean
}