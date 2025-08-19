// Auto generated code; DO NOT EDIT.

import { CSINodeDriver } from "./CSINodeDriver"

/**
 * CSINodeSpec holds information about the specification of all CSI drivers installed on a node
 * 
 */
export declare class CSINodeSpec {
    constructor();
    constructor(spec: CSINodeSpec);

	/**
     * Drivers is a list of information of all CSI Drivers existing on a node. If all drivers in the list are uninstalled, this can become empty.
     * 
     */
    drivers: Array<CSINodeDriver>
}