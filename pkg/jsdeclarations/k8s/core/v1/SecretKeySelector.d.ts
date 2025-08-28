// Auto generated code; DO NOT EDIT.

/**
 * SecretKeySelector selects a key of a Secret.
 */
export declare class SecretKeySelector {
    constructor();
    constructor(spec: Pick<SecretKeySelector, "key" | "name" | "optional">);

	/**
     * The key of the secret to select from.  Must be a valid secret key.
     */
    key: string

	/**
     * Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
     */
    name?: string

	/**
     * Specify whether the Secret or its key must be defined
     */
    optional?: boolean

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}