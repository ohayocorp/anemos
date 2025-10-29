// Auto generated code; DO NOT EDIT.
import { RollingUpdateStatefulSetStrategy } from "./RollingUpdateStatefulSetStrategy"

/**
 * StatefulSetUpdateStrategy indicates the strategy that the StatefulSet controller will use to perform updates. It includes any additional parameters necessary to perform the update for the indicated strategy.
 */
export declare class StatefulSetUpdateStrategy {
    constructor();
    constructor(spec: Pick<StatefulSetUpdateStrategy, "rollingUpdate" | "type">);

	/**
     * RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.
     */
    rollingUpdate?: RollingUpdateStatefulSetStrategy

	/**
     * Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
     */
    type?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}