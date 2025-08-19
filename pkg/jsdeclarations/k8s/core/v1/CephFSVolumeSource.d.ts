// Auto generated code; DO NOT EDIT.

import { LocalObjectReference } from "./LocalObjectReference"

/**
 * Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
 * 
 */
export declare class CephFSVolumeSource {
    constructor();
    constructor(spec: CephFSVolumeSource);

	/**
     * Monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     * 
     */
    monitors: Array<string>

	/**
     * Path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /
     * 
     */
    path?: string

	/**
     * ReadOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     * 
     */
    readOnly?: boolean

	/**
     * SecretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     * 
     */
    secretFile?: string

	/**
     * SecretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     * 
     */
    secretRef?: LocalObjectReference

	/**
     * User is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it
     * 
     */
    user?: string
}