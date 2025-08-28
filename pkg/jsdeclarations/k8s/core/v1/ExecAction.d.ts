// Auto generated code; DO NOT EDIT.

/**
 * ExecAction describes a "run in container" action.
 */
export declare class ExecAction {
    constructor();
    constructor(spec: Pick<ExecAction, "command">);

	/**
     * Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.
     */
    command?: Array<string>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}