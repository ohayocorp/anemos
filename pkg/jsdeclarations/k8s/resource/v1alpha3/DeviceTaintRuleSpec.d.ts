// Auto generated code; DO NOT EDIT.

import { DeviceTaint } from "./DeviceTaint"
import { DeviceTaintSelector } from "./DeviceTaintSelector"

/**
 * DeviceTaintRuleSpec specifies the selector and one taint.
 * 
 */
export declare class DeviceTaintRuleSpec {
    constructor();
    constructor(spec: DeviceTaintRuleSpec);

	/**
     * DeviceSelector defines which device(s) the taint is applied to. All selector criteria must be satified for a device to match. The empty selector matches all devices. Without a selector, no devices are matches.
     * 
     */
    deviceSelector?: DeviceTaintSelector

	/**
     * The taint that gets applied to matching devices.
     * 
     */
    taint: DeviceTaint
}