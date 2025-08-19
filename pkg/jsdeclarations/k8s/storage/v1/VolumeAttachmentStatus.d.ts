// Auto generated code; DO NOT EDIT.

import { VolumeError } from "./VolumeError"

/**
 * VolumeAttachmentStatus is the status of a VolumeAttachment request.
 * 
 */
export declare class VolumeAttachmentStatus {
    constructor();
    constructor(spec: VolumeAttachmentStatus);

	/**
     * AttachError represents the last error encountered during attach operation, if any. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
     * 
     */
    attachError?: VolumeError

	/**
     * Attached indicates the volume is successfully attached. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
     * 
     */
    attached: boolean

	/**
     * AttachmentMetadata is populated with any information returned by the attach operation, upon successful attach, that must be passed into subsequent WaitForAttach or Mount calls. This field must only be set by the entity completing the attach operation, i.e. the external-attacher.
     * 
     */
    attachmentMetadata?: Record<string, string>

	/**
     * DetachError represents the last error encountered during detach operation, if any. This field must only be set by the entity completing the detach operation, i.e. the external-attacher.
     * 
     */
    detachError?: VolumeError
}