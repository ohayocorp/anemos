// Auto generated code; DO NOT EDIT.

/**
 * CrossVersionObjectReference contains enough information to let you identify the referred resource.
 */
export declare class CrossVersionObjectReference {
    constructor();
    constructor(spec: Pick<CrossVersionObjectReference, "name">);

	/**
     * ApiVersion is the API version of the referent
     */
    apiVersion?: string

	/**
     * Kind is the kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind: string

	/**
     * Name is the name of the referent; More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}