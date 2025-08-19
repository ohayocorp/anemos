// Auto generated code; DO NOT EDIT.

import { VolumeAttachmentSource } from "./VolumeAttachmentSource"

/**
 * VolumeAttachmentSpec is the specification of a VolumeAttachment request.
 * 
 */
export declare class VolumeAttachmentSpec {
    constructor();
    constructor(spec: VolumeAttachmentSpec);

	/**
     * Attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().
     * 
     */
    attacher: string

	/**
     * NodeName represents the node that the volume should be attached to.
     * 
     */
    nodeName: string

	/**
     * Source represents the volume that should be attached.
     * 
     */
    source: VolumeAttachmentSource
}