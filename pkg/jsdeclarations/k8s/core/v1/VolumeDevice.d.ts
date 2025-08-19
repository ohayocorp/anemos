// Auto generated code; DO NOT EDIT.



/**
 * VolumeDevice describes a mapping of a raw block device within a container.
 * 
 */
export declare class VolumeDevice {
    constructor();
    constructor(spec: VolumeDevice);

	/**
     * DevicePath is the path inside of the container that the device will be mapped to.
     * 
     */
    devicePath: string

	/**
     * Name must match the name of a persistentVolumeClaim in the pod
     * 
     */
    name: string
}