// Auto generated code; DO NOT EDIT.
import { SecretReference } from "./SecretReference"

/**
 * Represents storage that is managed by an external CSI volume driver
 */
export declare class CSIPersistentVolumeSource {
    constructor();
    constructor(spec: Pick<CSIPersistentVolumeSource, "controllerExpandSecretRef" | "controllerPublishSecretRef" | "driver" | "fsType" | "nodeExpandSecretRef" | "nodePublishSecretRef" | "nodeStageSecretRef" | "readOnly" | "volumeAttributes" | "volumeHandle">);

	/**
     * ControllerExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    controllerExpandSecretRef?: SecretReference

	/**
     * ControllerPublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    controllerPublishSecretRef?: SecretReference

	/**
     * Driver is the name of the driver to use for this volume. Required.
     */
    driver: string

	/**
     * FsType to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
     */
    fsType?: string

	/**
     * NodeExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeExpandVolume call. This field is optional, may be omitted if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    nodeExpandSecretRef?: SecretReference

	/**
     * NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    nodePublishSecretRef?: SecretReference

	/**
     * NodeStageSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
     */
    nodeStageSecretRef?: SecretReference

	/**
     * ReadOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
     */
    readOnly?: boolean

	/**
     * VolumeAttributes of the volume to publish.
     */
    volumeAttributes?: Record<string, string>

	/**
     * VolumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
     */
    volumeHandle: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}