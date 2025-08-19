// Auto generated code; DO NOT EDIT.

import { UserInfo } from "./UserInfo"

/**
 * SelfSubjectReviewStatus is filled by the kube-apiserver and sent back to a user.
 * 
 */
export declare class SelfSubjectReviewStatus {
    constructor();
    constructor(spec: SelfSubjectReviewStatus);

	/**
     * User attributes of the user making this request.
     * 
     */
    userInfo?: UserInfo
}