// Auto generated code; DO NOT EDIT.

import { DownwardAPIVolumeFile } from "./DownwardAPIVolumeFile"

/**
 * Represents downward API info for projecting into a projected volume. Note that this is identical to a downwardAPI volume source without the default mode.
 * 
 */
export declare class DownwardAPIProjection {
    constructor();
    constructor(spec: DownwardAPIProjection);

	/**
     * Items is a list of DownwardAPIVolume file
     * 
     */
    items?: Array<DownwardAPIVolumeFile>
}