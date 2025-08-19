// Auto generated code; DO NOT EDIT.

import { ListMeta } from "../../apimachinery/meta/v1"
import { Namespace } from "./Namespace"

/**
 * NamespaceList is a list of Namespaces.
 * 
 */
export declare class NamespaceList {
    constructor();
    constructor(spec: Omit<NamespaceList, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Items is the list of Namespace objects in the list. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
     * 
     */
    items: Array<Namespace>

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    metadata?: ListMeta
}