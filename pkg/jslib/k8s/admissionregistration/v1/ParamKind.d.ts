// Auto generated code; DO NOT EDIT.

/**
 * ParamKind is a tuple of Group Kind and Version.
 */
export declare class ParamKind {
    constructor();
    constructor(spec: {});

	/**
     * APIVersion is the API group version the resources belong to. In format of "group/version". Required.
     */
    apiVersion?: string

	/**
     * Kind is the API kind the resources belong to. Required.
     */
    kind?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}