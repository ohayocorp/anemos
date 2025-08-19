// Auto generated code; DO NOT EDIT.



/**
 * ResourceQuotaStatus defines the enforced hard limits and observed use.
 * 
 */
export declare class ResourceQuotaStatus {
    constructor();
    constructor(spec: ResourceQuotaStatus);

	/**
     * Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/
     * 
     */
    hard?: any

	/**
     * Used is the current observed total usage of the resource in the namespace.
     * 
     */
    used?: any
}