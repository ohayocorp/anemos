// Auto generated code; DO NOT EDIT.

import { DeviceClassConfiguration } from "./DeviceClassConfiguration"
import { DeviceSelector } from "./DeviceSelector"

/**
 * DeviceClassSpec is used in a [DeviceClass] to define what can be allocated and how to configure it.
 * 
 */
export declare class DeviceClassSpec {
    constructor();
    constructor(spec: DeviceClassSpec);

	/**
     * Config defines configuration parameters that apply to each device that is claimed via this class. Some classses may potentially be satisfied by multiple drivers, so each instance of a vendor configuration applies to exactly one driver.
     * 
     * They are passed to the driver, but are not considered while allocating the claim.
     * 
     */
    config?: Array<DeviceClassConfiguration>

	/**
     * ExtendedResourceName is the extended resource name for the devices of this class. The devices of this class can be used to satisfy a pod's extended resource requests. It has the same format as the name of a pod's extended resource. It should be unique among all the device classes in a cluster. If two device classes have the same name, then the class created later is picked to satisfy a pod's extended resource requests. If two classes are created at the same time, then the name of the class lexicographically sorted first is picked.
     * 
     * This is an alpha field.
     * 
     */
    extendedResourceName?: string

	/**
     * Each selector must be satisfied by a device which is claimed via this class.
     * 
     */
    selectors?: Array<DeviceSelector>
}