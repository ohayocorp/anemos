// Auto generated code; DO NOT EDIT.
import { NodeSelector } from "./NodeSelector"

/**
 * VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
 */
export declare class VolumeNodeAffinity {
    constructor();
    constructor(spec: Pick<VolumeNodeAffinity, "required">);

	/**
     * Required specifies hard node constraints that must be met.
     */
    required?: NodeSelector

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}