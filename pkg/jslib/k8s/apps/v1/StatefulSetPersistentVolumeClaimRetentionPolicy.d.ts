// Auto generated code; DO NOT EDIT.

/**
 * StatefulSetPersistentVolumeClaimRetentionPolicy describes the policy used for PVCs created from the StatefulSet VolumeClaimTemplates.
 */
export declare class StatefulSetPersistentVolumeClaimRetentionPolicy {
    constructor();
    constructor(spec: Pick<StatefulSetPersistentVolumeClaimRetentionPolicy, "whenDeleted" | "whenScaled">);

	/**
     * WhenDeleted specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is deleted. The default policy of `Retain` causes PVCs to not be affected by StatefulSet deletion. The `Delete` policy causes those PVCs to be deleted.
     */
    whenDeleted?: string

	/**
     * WhenScaled specifies what happens to PVCs created from StatefulSet VolumeClaimTemplates when the StatefulSet is scaled down. The default policy of `Retain` causes PVCs to not be affected by a scaledown. The `Delete` policy causes the associated PVCs for any excess pods above the replica count to be deleted.
     */
    whenScaled?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}