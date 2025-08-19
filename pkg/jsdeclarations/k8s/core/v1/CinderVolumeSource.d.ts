// Auto generated code; DO NOT EDIT.

import { LocalObjectReference } from "./LocalObjectReference"

/**
 * Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.
 * 
 */
export declare class CinderVolumeSource {
    constructor();
    constructor(spec: CinderVolumeSource);

	/**
     * FsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     * 
     */
    fsType?: string

	/**
     * ReadOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     * 
     */
    readOnly?: boolean

	/**
     * SecretRef is optional: points to a secret object containing parameters used to connect to OpenStack.
     * 
     */
    secretRef?: LocalObjectReference

	/**
     * VolumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
     * 
     */
    volumeID: string
}