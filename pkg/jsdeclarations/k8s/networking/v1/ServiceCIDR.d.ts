// Auto generated code; DO NOT EDIT.
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import { ServiceCIDRSpec } from "./ServiceCIDRSpec"
import {Document} from '@ohayocorp/anemos';

/**
 * ServiceCIDR defines a range of IP addresses using CIDR format (e.g. 192.168.0.0/24 or 2001:db2::/64). This range is used to allocate ClusterIPs to Service objects.
 */
export declare class ServiceCIDR extends Document {
    constructor();
    constructor(spec: Pick<ServiceCIDR, "metadata" | "spec">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta

	/**
     * Spec is the desired state of the ServiceCIDR. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
     */
    spec?: ServiceCIDRSpec
}