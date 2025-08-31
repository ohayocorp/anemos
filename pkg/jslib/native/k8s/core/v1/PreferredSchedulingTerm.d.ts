// Auto generated code; DO NOT EDIT.
import { NodeSelectorTerm } from "./NodeSelectorTerm"

/**
 * An empty preferred scheduling term matches all objects with implicit weight 0 (i.e. it's a no-op). A null preferred scheduling term matches no objects (i.e. is also a no-op).
 */
export declare class PreferredSchedulingTerm {
    constructor();
    constructor(spec: Pick<PreferredSchedulingTerm, "preference" | "weight">);

	/**
     * A node selector term, associated with the corresponding weight.
     */
    preference: NodeSelectorTerm

	/**
     * Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.
     */
    weight: number

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}