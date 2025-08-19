// Auto generated code; DO NOT EDIT.



/**
 * VolumeError captures an error encountered during a volume operation.
 * 
 */
export declare class VolumeError {
    constructor();
    constructor(spec: VolumeError);

	/**
     * ErrorCode is a numeric gRPC code representing the error encountered during Attach or Detach operations.
     * 
     * This is an optional, beta field that requires the MutableCSINodeAllocatableCount feature gate being enabled to be set.
     * 
     */
    errorCode?: number

	/**
     * Message represents the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
     * 
     */
    message?: string

	/**
     * Time represents the time the error was encountered.
     * 
     */
    time?: string
}