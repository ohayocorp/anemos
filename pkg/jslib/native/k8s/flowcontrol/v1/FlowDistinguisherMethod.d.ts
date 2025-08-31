// Auto generated code; DO NOT EDIT.

/**
 * FlowDistinguisherMethod specifies the method of a flow distinguisher.
 */
export declare class FlowDistinguisherMethod {
    constructor();
    constructor(spec: Pick<FlowDistinguisherMethod, "type">);

	/**
     * `type` is the type of flow distinguisher method The supported types are "ByUser" and "ByNamespace". Required.
     */
    type: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}