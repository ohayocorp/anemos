// Auto generated code; DO NOT EDIT.

import { ApplyConfiguration } from "./ApplyConfiguration"
import { JSONPatch } from "./JSONPatch"

/**
 * Mutation specifies the CEL expression which is used to apply the Mutation.
 * 
 */
export declare class Mutation {
    constructor();
    constructor(spec: Mutation);

	/**
     * ApplyConfiguration defines the desired configuration values of an object. The configuration is applied to the admission object using [structured merge diff](https://github.com/kubernetes-sigs/structured-merge-diff). A CEL expression is used to create apply configuration.
     * 
     */
    applyConfiguration?: ApplyConfiguration

	/**
     * JsonPatch defines a [JSON patch](https://jsonpatch.com/) operation to perform a mutation to the object. A CEL expression is used to create the JSON patch.
     * 
     */
    jsonPatch?: JSONPatch

	/**
     * PatchType indicates the patch strategy used. Allowed values are "ApplyConfiguration" and "JSONPatch". Required.
     * 
     */
    patchType: string
}