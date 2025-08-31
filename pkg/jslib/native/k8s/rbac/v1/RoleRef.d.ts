// Auto generated code; DO NOT EDIT.

/**
 * RoleRef contains information that points to the role being used
 */
export declare class RoleRef {
    constructor();
    constructor(spec: Pick<RoleRef, "apiGroup" | "name">);

	/**
     * APIGroup is the group for the resource being referenced
     */
    apiGroup: string

	/**
     * Kind is the type of resource being referenced
     */
    kind: string

	/**
     * Name is the name of resource being referenced
     */
    name: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}