// Auto generated code; DO NOT EDIT.

/**
 * FileKeySelector selects a key of the env file.
 */
export declare class FileKeySelector {
    constructor();
    constructor(spec: Pick<FileKeySelector, "key" | "optional" | "path" | "volumeName">);

	/**
     * The key within the env file. An invalid key will prevent the pod from starting. The keys defined within a source may consist of any printable ASCII characters except '='. During Alpha stage of the EnvFiles feature gate, the key size is limited to 128 characters.
     */
    key: string

	/**
     * Specify whether the file or its key must be defined. If the file or key does not exist, then the env var is not published. If optional is set to true and the specified key does not exist, the environment variable will not be set in the Pod's containers.
    
     * If optional is set to false and the specified key does not exist, an error will be returned during Pod creation.
     */
    optional?: boolean

	/**
     * The path within the volume from which to select the file. Must be relative and may not contain the '..' path or start with '..'.
     */
    path: string

	/**
     * The name of the volume mount containing the env file.
     */
    volumeName: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}