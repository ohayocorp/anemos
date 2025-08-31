// Auto generated code; DO NOT EDIT.

/**
 * Local represents directly-attached storage with node affinity
 */
export declare class LocalVolumeSource {
    constructor();
    constructor(spec: Pick<LocalVolumeSource, "fsType" | "path">);

	/**
     * FsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". The default value is to auto-select a filesystem if unspecified.
     */
    fsType?: string

	/**
     * Path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).
     */
    path: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}