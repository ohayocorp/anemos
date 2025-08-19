// Auto generated code; DO NOT EDIT.

import { ParentReference } from "./ParentReference"

/**
 * IPAddressSpec describe the attributes in an IP Address.
 * 
 */
export declare class IPAddressSpec {
    constructor();
    constructor(spec: IPAddressSpec);

	/**
     * ParentRef references the resource that an IPAddress is attached to. An IPAddress must reference a parent object.
     * 
     */
    parentRef: ParentReference
}