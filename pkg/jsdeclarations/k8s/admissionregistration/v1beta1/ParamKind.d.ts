// Auto generated code; DO NOT EDIT.



/**
 * ParamKind is a tuple of Group Kind and Version.
 * 
 */
export declare class ParamKind {
    constructor();
    constructor(spec: ParamKind);

	/**
     * APIVersion is the API group version the resources belong to. In format of "group/version". Required.
     * 
     */
    apiVersion?: string

	/**
     * Kind is the API kind the resources belong to. Required.
     * 
     */
    kind?: string
}