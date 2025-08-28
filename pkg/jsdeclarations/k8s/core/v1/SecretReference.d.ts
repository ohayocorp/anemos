// Auto generated code; DO NOT EDIT.

/**
 * SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace
 */
export declare class SecretReference {
    constructor();
    constructor(spec: Pick<SecretReference, "name" | "namespace">);

	/**
     * Name is unique within a namespace to reference a secret resource.
     */
    name?: string

	/**
     * Namespace defines the space within which the secret name must be unique.
     */
    namespace?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}