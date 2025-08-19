// Auto generated code; DO NOT EDIT.

import { NodeSelector } from "./NodeSelector"

/**
 * VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.
 * 
 */
export declare class VolumeNodeAffinity {
    constructor();
    constructor(spec: VolumeNodeAffinity);

	/**
     * Required specifies hard node constraints that must be met.
     * 
     */
    required?: NodeSelector
}