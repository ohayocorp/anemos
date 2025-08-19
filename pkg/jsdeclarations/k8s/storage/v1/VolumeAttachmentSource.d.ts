// Auto generated code; DO NOT EDIT.

import { PersistentVolumeSpec } from "../../core/v1"

/**
 * VolumeAttachmentSource represents a volume that should be attached. Right now only PersistentVolumes can be attached via external attacher, in the future we may allow also inline volumes in pods. Exactly one member can be set.
 * 
 */
export declare class VolumeAttachmentSource {
    constructor();
    constructor(spec: VolumeAttachmentSource);

	/**
     * InlineVolumeSpec contains all the information necessary to attach a persistent volume defined by a pod's inline VolumeSource. This field is populated only for the CSIMigration feature. It contains translated fields from a pod's inline VolumeSource to a PersistentVolumeSpec. This field is beta-level and is only honored by servers that enabled the CSIMigration feature.
     * 
     */
    inlineVolumeSpec?: PersistentVolumeSpec

	/**
     * PersistentVolumeName represents the name of the persistent volume to attach.
     * 
     */
    persistentVolumeName?: string
}