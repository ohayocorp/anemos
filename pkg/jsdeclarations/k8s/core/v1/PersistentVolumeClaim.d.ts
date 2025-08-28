// Auto generated code; DO NOT EDIT.
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import { PersistentVolumeClaimSpec } from "./PersistentVolumeClaimSpec"
import {Document} from '@ohayocorp/anemos';

/**
 * PersistentVolumeClaim is a user's request for and claim to a persistent volume
 */
export declare class PersistentVolumeClaim extends Document {
    constructor();
    constructor(spec: Pick<PersistentVolumeClaim, "metadata" | "spec">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta

	/**
     * Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims
     */
    spec?: PersistentVolumeClaimSpec
}