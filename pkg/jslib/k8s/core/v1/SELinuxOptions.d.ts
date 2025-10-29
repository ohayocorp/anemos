// Auto generated code; DO NOT EDIT.

/**
 * SELinuxOptions are the labels to be applied to the container
 */
export declare class SELinuxOptions {
    constructor();
    constructor(spec: Pick<SELinuxOptions, "level" | "role" | "type" | "user">);

	/**
     * Level is SELinux level label that applies to the container.
     */
    level?: string

	/**
     * Role is a SELinux role label that applies to the container.
     */
    role?: string

	/**
     * Type is a SELinux type label that applies to the container.
     */
    type?: string

	/**
     * User is a SELinux user label that applies to the container.
     */
    user?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}