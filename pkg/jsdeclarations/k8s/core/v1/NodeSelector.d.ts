// Auto generated code; DO NOT EDIT.
import { NodeSelectorTerm } from "./NodeSelectorTerm"

/**
 * A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.
 */
export declare class NodeSelector {
    constructor();
    constructor(spec: Pick<NodeSelector, "nodeSelectorTerms">);

	/**
     * Required. A list of node selector terms. The terms are ORed.
     */
    nodeSelectorTerms: Array<NodeSelectorTerm>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}