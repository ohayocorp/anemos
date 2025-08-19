// Auto generated code; DO NOT EDIT.



/**
 * Describes the state of a migration at a certain point.
 * 
 */
export declare class MigrationCondition {
    constructor();
    constructor(spec: MigrationCondition);

	/**
     * The last time this condition was updated.
     * 
     */
    lastUpdateTime?: string

	/**
     * A human readable message indicating details about the transition.
     * 
     */
    message?: string

	/**
     * The reason for the condition's last transition.
     * 
     */
    reason?: string

	/**
     * Status of the condition, one of True, False, Unknown.
     * 
     */
    status: string

	/**
     * Type of the condition.
     * 
     */
    type: string
}