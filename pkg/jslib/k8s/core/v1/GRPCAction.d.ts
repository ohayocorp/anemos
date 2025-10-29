// Auto generated code; DO NOT EDIT.

/**
 * GRPCAction specifies an action involving a GRPC service.
 */
export declare class GRPCAction {
    constructor();
    constructor(spec: Pick<GRPCAction, "port" | "service">);

	/**
     * Port number of the gRPC service. Number must be in the range 1 to 65535.
     */
    port: number

	/**
     * Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
    
     * If this is not specified, the default behavior is defined by gRPC.
     */
    service?: string

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}