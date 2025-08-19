// Auto generated code; DO NOT EDIT.



/**
 * PersistentVolumeClaimCondition contains details about state of pvc
 * 
 */
export declare class PersistentVolumeClaimCondition {
    constructor();
    constructor(spec: PersistentVolumeClaimCondition);

	/**
     * LastProbeTime is the time we probed the condition.
     * 
     */
    lastProbeTime?: string

	/**
     * LastTransitionTime is the time the condition transitioned from one status to another.
     * 
     */
    lastTransitionTime?: string

	/**
     * Message is the human-readable message indicating details about last transition.
     * 
     */
    message?: string

	/**
     * Reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports "Resizing" that means the underlying persistent volume is being resized.
     * 
     */
    reason?: string

	/**
     * Type is the type of the condition. More info: https://kubernetes.io/docs/reference/kubernetes-api/config-and-storage-resources/persistent-volume-claim-v1/#:~:text=set%20to%20%27ResizeStarted%27.-,PersistentVolumeClaimCondition,-contains%20details%20about
     * 
     */
    type: string
}