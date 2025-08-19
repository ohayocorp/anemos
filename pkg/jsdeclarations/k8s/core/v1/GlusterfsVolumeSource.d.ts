// Auto generated code; DO NOT EDIT.



/**
 * Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.
 * 
 */
export declare class GlusterfsVolumeSource {
    constructor();
    constructor(spec: GlusterfsVolumeSource);

	/**
     * Endpoints is the endpoint name that details Glusterfs topology.
     * 
     */
    endpoints: string

	/**
     * Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     * 
     */
    path: string

	/**
     * ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod
     * 
     */
    readOnly?: boolean
}