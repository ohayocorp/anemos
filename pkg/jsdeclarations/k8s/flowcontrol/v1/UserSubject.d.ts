// Auto generated code; DO NOT EDIT.

/**
 * UserSubject holds detailed information for user-kind subject.
 */
export declare class UserSubject {
    constructor();
    constructor(spec: Pick<UserSubject, "name">);

	/**
     * `name` is the username that matches, or "*" to match all usernames. Required.
     */
    name: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}