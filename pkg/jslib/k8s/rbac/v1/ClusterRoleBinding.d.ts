// Auto generated code; DO NOT EDIT.
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import { RoleRef } from "./RoleRef"
import { Subject } from "./Subject"
import {Document} from '@ohayocorp/anemos';

/**
 * ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.
 */
export declare class ClusterRoleBinding extends Document {
    constructor();
    constructor(spec: Pick<ClusterRoleBinding, "metadata" | "roleRef" | "subjects">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object's metadata.
     */
    metadata?: ObjectMeta

	/**
     * RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error. This field is immutable.
     */
    roleRef: RoleRef

	/**
     * Subjects holds references to the objects the role applies to.
     */
    subjects?: Array<Subject>
}