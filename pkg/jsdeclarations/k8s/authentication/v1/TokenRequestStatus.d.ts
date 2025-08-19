// Auto generated code; DO NOT EDIT.



/**
 * TokenRequestStatus is the result of a token request.
 * 
 */
export declare class TokenRequestStatus {
    constructor();
    constructor(spec: TokenRequestStatus);

	/**
     * ExpirationTimestamp is the time of expiration of the returned token.
     * 
     */
    expirationTimestamp: string

	/**
     * Token is the opaque bearer token.
     * 
     */
    token: string
}