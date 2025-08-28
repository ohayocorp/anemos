// Auto generated code; DO NOT EDIT.
import { NodeSelectorRequirement } from "./NodeSelectorRequirement"

/**
 * A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
 */
export declare class NodeSelectorTerm {
    constructor();
    constructor(spec: Pick<NodeSelectorTerm, "matchExpressions" | "matchFields">);

	/**
     * A list of node selector requirements by node's labels.
     */
    matchExpressions?: Array<NodeSelectorRequirement>

	/**
     * A list of node selector requirements by node's fields.
     */
    matchFields?: Array<NodeSelectorRequirement>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}