// Auto generated code; DO NOT EDIT.



/**
 * Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.
 * 
 */
export declare class FCVolumeSource {
    constructor();
    constructor(spec: FCVolumeSource);

	/**
     * FsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified.
     * 
     */
    fsType?: string

	/**
     * Lun is Optional: FC target lun number
     * 
     */
    lun?: number

	/**
     * ReadOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     * 
     */
    readOnly?: boolean

	/**
     * TargetWWNs is Optional: FC target worldwide names (WWNs)
     * 
     */
    targetWWNs?: Array<string>

	/**
     * Wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.
     * 
     */
    wwids?: Array<string>
}