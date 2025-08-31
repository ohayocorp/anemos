// Auto generated code; DO NOT EDIT.
import { CustomResourceDefinitionSpec } from "./CustomResourceDefinitionSpec"
import { ObjectMeta } from "./../../apimachinery/meta/v1"
import {Document} from '@ohayocorp/anemos';

/**
 * CustomResourceDefinition represents a resource that should be exposed on the API server.  Its name MUST be in the format <.spec.name>.<.spec.group>.
 */
export declare class CustomResourceDefinition extends Document {
    constructor();
    constructor(spec: Pick<CustomResourceDefinition, "metadata" | "spec">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     */
    kind?: string

	/**
     * Standard object's metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     */
    metadata?: ObjectMeta

	/**
     * Spec describes how the user wants the resources to appear
     */
    spec: CustomResourceDefinitionSpec
}