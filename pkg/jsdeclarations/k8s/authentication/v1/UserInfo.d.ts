// Auto generated code; DO NOT EDIT.



/**
 * UserInfo holds the information about the user needed to implement the user.Info interface.
 * 
 */
export declare class UserInfo {
    constructor();
    constructor(spec: UserInfo);

	/**
     * Any additional information provided by the authenticator.
     * 
     */
    extra?: Record<string, Array<string>>

	/**
     * The names of groups this user is a part of.
     * 
     */
    groups?: Array<string>

	/**
     * A unique value that identifies this user across time. If this user is deleted and another user by the same name is added, they will have different UIDs.
     * 
     */
    uid?: string

	/**
     * The name that uniquely identifies this user among all active users.
     * 
     */
    username?: string
}