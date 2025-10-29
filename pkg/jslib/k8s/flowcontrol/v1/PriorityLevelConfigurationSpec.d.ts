// Auto generated code; DO NOT EDIT.
import { ExemptPriorityLevelConfiguration } from "./ExemptPriorityLevelConfiguration"
import { LimitedPriorityLevelConfiguration } from "./LimitedPriorityLevelConfiguration"

/**
 * PriorityLevelConfigurationSpec specifies the configuration of a priority level.
 */
export declare class PriorityLevelConfigurationSpec {
    constructor();
    constructor(spec: Pick<PriorityLevelConfigurationSpec, "exempt" | "limited" | "type">);

	/**
     * `exempt` specifies how requests are handled for an exempt priority level. This field MUST be empty if `type` is `"Limited"`. This field MAY be non-empty if `type` is `"Exempt"`. If empty and `type` is `"Exempt"` then the default values for `ExemptPriorityLevelConfiguration` apply.
     */
    exempt?: ExemptPriorityLevelConfiguration

	/**
     * `limited` specifies how requests are handled for a Limited priority level. This field must be non-empty if and only if `type` is `"Limited"`.
     */
    limited?: LimitedPriorityLevelConfiguration

	/**
     * `type` indicates whether this priority level is subject to limitation on request execution.  A value of `"Exempt"` means that requests of this priority level are not subject to a limit (and thus are never queued) and do not detract from the capacity made available to other priority levels.  A value of `"Limited"` means that (a) requests of this priority level _are_ subject to limits and (b) some of the server's limited capacity is made available exclusively to this priority level. Required.
     */
    type: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}