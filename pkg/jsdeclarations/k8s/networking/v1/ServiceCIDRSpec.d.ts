// Auto generated code; DO NOT EDIT.

/**
 * ServiceCIDRSpec define the CIDRs the user wants to use for allocating ClusterIPs for Services.
 */
export declare class ServiceCIDRSpec {
    constructor();
    constructor(spec: Pick<ServiceCIDRSpec, "cidrs">);

	/**
     * CIDRs defines the IP blocks in CIDR notation (e.g. "192.168.0.0/24" or "2001:db8::/64") from which to assign service cluster IPs. Max of two CIDRs is allowed, one of each IP family. This field is immutable.
     */
    cidrs?: Array<string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}