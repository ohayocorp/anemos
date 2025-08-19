// Auto generated code; DO NOT EDIT.



/**
 * ContainerStateTerminated is a terminated state of a container.
 * 
 */
export declare class ContainerStateTerminated {
    constructor();
    constructor(spec: ContainerStateTerminated);

	/**
     * Container's ID in the format '<type>://<container_id>'
     * 
     */
    containerID?: string

	/**
     * Exit status from the last termination of the container
     * 
     */
    exitCode: number

	/**
     * Time at which the container last terminated
     * 
     */
    finishedAt?: string

	/**
     * Message regarding the last termination of the container
     * 
     */
    message?: string

	/**
     * (brief) reason from the last termination of the container
     * 
     */
    reason?: string

	/**
     * Signal from the last termination of the container
     * 
     */
    signal?: number

	/**
     * Time at which previous execution of the container started
     * 
     */
    startedAt?: string
}