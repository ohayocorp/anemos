// Auto generated code; DO NOT EDIT.



/**
 * TokenRequest contains parameters of a service account token.
 * 
 */
export declare class TokenRequest {
    constructor();
    constructor(spec: TokenRequest);

	/**
     * Audience is the intended audience of the token in "TokenRequestSpec". It will default to the audiences of kube apiserver.
     * 
     */
    audience: string

	/**
     * ExpirationSeconds is the duration of validity of the token in "TokenRequestSpec". It has the same default value of "ExpirationSeconds" in "TokenRequestSpec".
     * 
     */
    expirationSeconds?: number
}