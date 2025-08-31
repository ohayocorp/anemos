// Auto generated code; DO NOT EDIT.

/**
 * PriorityLevelConfigurationReference contains information that points to the "request-priority" being used.
 */
export declare class PriorityLevelConfigurationReference {
    constructor();
    constructor(spec: Pick<PriorityLevelConfigurationReference, "name">);

	/**
     * `name` is the name of the priority level configuration being referenced Required.
     */
    name: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}