// Auto generated code; DO NOT EDIT.

/**
 * TokenRequest contains parameters of a service account token.
 */
export declare class TokenRequest {
    constructor();
    constructor(spec: Pick<TokenRequest, "audience" | "expirationSeconds">);

	/**
     * Audience is the intended audience of the token in "TokenRequestSpec". It will default to the audiences of kube apiserver.
     */
    audience: string

	/**
     * ExpirationSeconds is the duration of validity of the token in "TokenRequestSpec". It has the same default value of "ExpirationSeconds" in "TokenRequestSpec".
     */
    expirationSeconds?: number

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}