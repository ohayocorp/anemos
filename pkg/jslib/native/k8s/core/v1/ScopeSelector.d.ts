// Auto generated code; DO NOT EDIT.
import { ScopedResourceSelectorRequirement } from "./ScopedResourceSelectorRequirement"

/**
 * A scope selector represents the AND of the selectors represented by the scoped-resource selector requirements.
 */
export declare class ScopeSelector {
    constructor();
    constructor(spec: Pick<ScopeSelector, "matchExpressions">);

	/**
     * A list of scope selector requirements by scope of the resources.
     */
    matchExpressions?: Array<ScopedResourceSelectorRequirement>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}