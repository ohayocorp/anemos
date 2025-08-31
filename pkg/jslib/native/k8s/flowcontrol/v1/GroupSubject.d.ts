// Auto generated code; DO NOT EDIT.

/**
 * GroupSubject holds detailed information for group-kind subject.
 */
export declare class GroupSubject {
    constructor();
    constructor(spec: Pick<GroupSubject, "name">);

	/**
     * Name is the user group that matches, or "*" to match all user groups. See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go for some well-known group names. Required.
     */
    name: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}