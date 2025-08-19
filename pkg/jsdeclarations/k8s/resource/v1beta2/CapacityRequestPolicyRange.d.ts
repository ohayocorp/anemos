// Auto generated code; DO NOT EDIT.



/**
 * CapacityRequestPolicyRange defines a valid range for consumable capacity values.
 * 
 *   - If the requested amount is less than Min, it is rounded up to the Min value.
 * 
 *   - If Step is set and the requested amount is between Min and Max but not aligned with Step,
 * 
 *     it will be rounded up to the next value equal to Min + (n * Step).
 * 
 *   - If Step is not set, the requested amount is used as-is if it falls within the range Min to Max (if set).
 * 
 *   - If the requested or rounded amount exceeds Max (if set), the request does not satisfy the policy,
 * 
 *     and the device cannot be allocated.
 * 
 */
export declare class CapacityRequestPolicyRange {
    constructor();
    constructor(spec: CapacityRequestPolicyRange);

	/**
     * Max defines the upper limit for capacity that can be requested.
     * 
     * Max must be less than or equal to the capacity value. Min and requestPolicy.default must be less than or equal to the maximum.
     * 
     */
    max?: any

	/**
     * Min specifies the minimum capacity allowed for a consumption request.
     * 
     * Min must be greater than or equal to zero, and less than or equal to the capacity value. requestPolicy.default must be more than or equal to the minimum.
     * 
     */
    min: any

	/**
     * Step defines the step size between valid capacity amounts within the range.
     * 
     * Max (if set) and requestPolicy.default must be a multiple of Step. Min + Step must be less than or equal to the capacity value.
     * 
     */
    step?: any
}