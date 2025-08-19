// Auto generated code; DO NOT EDIT.



/**
 * The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
 * 
 */
export declare class Taint {
    constructor();
    constructor(spec: Taint);

	/**
     * Required. The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
     * 
     */
    effect: string

	/**
     * Required. The taint key to be applied to a node.
     * 
     */
    key: string

	/**
     * TimeAdded represents the time at which the taint was added.
     * 
     */
    timeAdded?: string

	/**
     * The taint value corresponding to the taint key.
     * 
     */
    value?: string
}