// Auto generated code; DO NOT EDIT.

/**
 * ContainerRestartRuleOnExitCodes describes the condition for handling an exited container based on its exit codes.
 */
export declare class ContainerRestartRuleOnExitCodes {
    constructor();
    constructor(spec: Pick<ContainerRestartRuleOnExitCodes, "operator" | "values">);

	/**
     * Represents the relationship between the container exit code(s) and the specified values. Possible values are: - In: the requirement is satisfied if the container exit code is in the
    
     *   set of specified values.
    
     * - NotIn: the requirement is satisfied if the container exit code is
    
     *   not in the set of specified values.
     */
    operator: string

	/**
     * Specifies the set of values to check for container exit codes. At most 255 elements are allowed.
     */
    values?: Array<number>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}