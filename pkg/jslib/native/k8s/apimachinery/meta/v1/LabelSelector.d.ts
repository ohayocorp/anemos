// Auto generated code; DO NOT EDIT.
import { LabelSelectorRequirement } from "./LabelSelectorRequirement"

/**
 * A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.
 */
export declare class LabelSelector {
    constructor();
    constructor(spec: Pick<LabelSelector, "matchExpressions" | "matchLabels">);

	/**
     * MatchExpressions is a list of label selector requirements. The requirements are ANDed.
     */
    matchExpressions?: Array<LabelSelectorRequirement>

	/**
     * MatchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.
     */
    matchLabels?: Record<string, string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}