// Auto generated code; DO NOT EDIT.
import { GroupSubject } from "./GroupSubject"
import { ServiceAccountSubject } from "./ServiceAccountSubject"
import { UserSubject } from "./UserSubject"

/**
 * Subject matches the originator of a request, as identified by the request authentication system. There are three ways of matching an originator; by user, group, or service account.
 */
export declare class Subject {
    constructor();
    constructor(spec: Pick<Subject, "group" | "serviceAccount" | "user">);

	/**
     * `group` matches based on user group name.
     */
    group?: GroupSubject

	/**
     * `kind` indicates which one of the other fields is non-empty. Required
     */
    kind: string

	/**
     * `serviceAccount` matches ServiceAccounts.
     */
    serviceAccount?: ServiceAccountSubject

	/**
     * `user` matches based on username.
     */
    user?: UserSubject

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}