// Auto generated code; DO NOT EDIT.



/**
 * VolumeResourceRequirements describes the storage resource requirements for a volume.
 * 
 */
export declare class VolumeResourceRequirements {
    constructor();
    constructor(spec: VolumeResourceRequirements);

	/**
     * Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     * 
     */
    limits?: any

	/**
     * Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
     * 
     */
    requests?: any
}