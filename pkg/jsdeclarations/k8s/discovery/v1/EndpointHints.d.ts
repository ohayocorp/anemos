// Auto generated code; DO NOT EDIT.

import { ForNode } from "./ForNode"
import { ForZone } from "./ForZone"

/**
 * EndpointHints provides hints describing how an endpoint should be consumed.
 * 
 */
export declare class EndpointHints {
    constructor();
    constructor(spec: EndpointHints);

	/**
     * ForNodes indicates the node(s) this endpoint should be consumed by when using topology aware routing. May contain a maximum of 8 entries. This is an Alpha feature and is only used when the PreferSameTrafficDistribution feature gate is enabled.
     * 
     */
    forNodes?: Array<ForNode>

	/**
     * ForZones indicates the zone(s) this endpoint should be consumed by when using topology aware routing. May contain a maximum of 8 entries.
     * 
     */
    forZones?: Array<ForZone>
}