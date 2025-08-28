// Auto generated code; DO NOT EDIT.
import { LabelSelector } from "./../../apimachinery/meta/v1"

/**
 * AggregationRule describes how to locate ClusterRoles to aggregate into the ClusterRole
 */
export declare class AggregationRule {
    constructor();
    constructor(spec: Pick<AggregationRule, "clusterRoleSelectors">);

	/**
     * ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole's permissions will be added
     */
    clusterRoleSelectors?: LabelSelector

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}