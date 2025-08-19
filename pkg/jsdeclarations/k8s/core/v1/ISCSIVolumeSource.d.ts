// Auto generated code; DO NOT EDIT.

import { LocalObjectReference } from "./LocalObjectReference"

/**
 * Represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.
 * 
 */
export declare class ISCSIVolumeSource {
    constructor();
    constructor(spec: ISCSIVolumeSource);

	/**
     * ChapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication
     * 
     */
    chapAuthDiscovery?: boolean

	/**
     * ChapAuthSession defines whether support iSCSI Session CHAP authentication
     * 
     */
    chapAuthSession?: boolean

	/**
     * FsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi
     * 
     */
    fsType?: string

	/**
     * InitiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.
     * 
     */
    initiatorName?: string

	/**
     * Iqn is the target iSCSI Qualified Name.
     * 
     */
    iqn: string

	/**
     * IscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).
     * 
     */
    iscsiInterface?: string

	/**
     * Lun represents iSCSI Target Lun number.
     * 
     */
    lun: number

	/**
     * Portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
     * 
     */
    portals?: Array<string>

	/**
     * ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.
     * 
     */
    readOnly?: boolean

	/**
     * SecretRef is the CHAP Secret for iSCSI target and initiator authentication
     * 
     */
    secretRef?: LocalObjectReference

	/**
     * TargetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).
     * 
     */
    targetPortal: string
}