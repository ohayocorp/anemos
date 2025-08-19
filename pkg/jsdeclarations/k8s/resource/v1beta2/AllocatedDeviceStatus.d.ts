// Auto generated code; DO NOT EDIT.

import { Condition } from "../../apimachinery/meta/v1"
import { NetworkDeviceData } from "./NetworkDeviceData"

/**
 * AllocatedDeviceStatus contains the status of an allocated device, if the driver chooses to report it. This may include driver-specific information.
 * 
 * The combination of Driver, Pool, Device, and ShareID must match the corresponding key in Status.Allocation.Devices.
 * 
 */
export declare class AllocatedDeviceStatus {
    constructor();
    constructor(spec: AllocatedDeviceStatus);

	/**
     * Conditions contains the latest observation of the device's state. If the device has been configured according to the class and claim config references, the `Ready` condition should be True.
     * 
     * Must not contain more than 8 entries.
     * 
     */
    conditions?: Condition

	/**
     * Data contains arbitrary driver-specific data.
     * 
     * The length of the raw data must be smaller or equal to 10 Ki.
     * 
     */
    data?: any

	/**
     * Device references one device instance via its name in the driver's resource pool. It must be a DNS label.
     * 
     */
    device: string

	/**
     * Driver specifies the name of the DRA driver whose kubelet plugin should be invoked to process the allocation once the claim is needed on a node.
     * 
     * Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
     * 
     */
    driver: string

	/**
     * NetworkData contains network-related information specific to the device.
     * 
     */
    networkData?: NetworkDeviceData

	/**
     * This name together with the driver name and the device name field identify which device was allocated (`<driver name>/<pool name>/<device name>`).
     * 
     * Must not be longer than 253 characters and may contain one or more DNS sub-domains separated by slashes.
     * 
     */
    pool: string

	/**
     * ShareID uniquely identifies an individual allocation share of the device.
     * 
     */
    shareID?: string
}