// Auto generated code; DO NOT EDIT.

/**
 * Overhead structure represents the resource overhead associated with running a pod.
 */
export declare class Overhead {
    constructor();
    constructor(spec: Pick<Overhead, "podFixed">);

	/**
     * PodFixed represents the fixed resource overhead associated with running a pod.
     */
    podFixed?: number | string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}