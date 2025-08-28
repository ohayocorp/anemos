// Auto generated code; DO NOT EDIT.

/**
 * LocalObjectReference contains enough information to let you locate the referenced object inside the same namespace.
 */
export declare class LocalObjectReference {
    constructor();
    constructor(spec: Pick<LocalObjectReference, "name">);

	/**
     * Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}