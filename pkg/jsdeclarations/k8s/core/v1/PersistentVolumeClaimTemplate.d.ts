// Auto generated code; DO NOT EDIT.
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import { PersistentVolumeClaimSpec } from "./PersistentVolumeClaimSpec"

/**
 * PersistentVolumeClaimTemplate is used to produce PersistentVolumeClaim objects as part of an EphemeralVolumeSource.
 */
export declare class PersistentVolumeClaimTemplate {
    constructor();
    constructor(spec: Pick<PersistentVolumeClaimTemplate, "metadata" | "spec">);

	/**
     * May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.
     */
    metadata?: ObjectMeta

	/**
     * The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.
     */
    spec: PersistentVolumeClaimSpec

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}