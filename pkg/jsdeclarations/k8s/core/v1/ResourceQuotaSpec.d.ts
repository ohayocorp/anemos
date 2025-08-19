// Auto generated code; DO NOT EDIT.

import { ScopeSelector } from "./ScopeSelector"

/**
 * ResourceQuotaSpec defines the desired hard limits to enforce for Quota.
 * 
 */
export declare class ResourceQuotaSpec {
    constructor();
    constructor(spec: ResourceQuotaSpec);

	/**
     * Hard is the set of desired hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
     * 
     */
    hard?: Record<string, any>

	/**
     * ScopeSelector is also a collection of filters like scopes that must match each object tracked by a quota but expressed using ScopeSelectorOperator in combination with possible values. For a resource to match, both scopes AND scopeSelector (if specified in spec), must be matched.
     * 
     */
    scopeSelector?: ScopeSelector

	/**
     * A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.
     * 
     */
    scopes?: Array<string>
}