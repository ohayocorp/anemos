// Auto generated code; DO NOT EDIT.

/**
 * PodSchedulingGate is associated to a Pod to guard its scheduling.
 */
export declare class PodSchedulingGate {
    constructor();
    constructor(spec: Pick<PodSchedulingGate, "name">);

	/**
     * Name of the scheduling gate. Each scheduling gate must have a unique name field.
     */
    name: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}