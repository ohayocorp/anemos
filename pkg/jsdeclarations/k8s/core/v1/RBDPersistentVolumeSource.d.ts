// Auto generated code; DO NOT EDIT.

import { SecretReference } from "./SecretReference"

/**
 * Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.
 * 
 */
export declare class RBDPersistentVolumeSource {
    constructor();
    constructor(spec: RBDPersistentVolumeSource);

	/**
     * FsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: "ext4", "xfs", "ntfs". Implicitly inferred to be "ext4" if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd
     * 
     */
    fsType?: string

	/**
     * Image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     * 
     */
    image: string

	/**
     * Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     * 
     */
    keyring?: string

	/**
     * Monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     * 
     */
    monitors: Array<string>

	/**
     * Pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     * 
     */
    pool?: string

	/**
     * ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     * 
     */
    readOnly?: boolean

	/**
     * SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     * 
     */
    secretRef?: SecretReference

	/**
     * User is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it
     * 
     */
    user?: string
}