// Auto generated code; DO NOT EDIT.
import { ClientIPConfig } from "./ClientIPConfig"

/**
 * SessionAffinityConfig represents the configurations of session affinity.
 */
export declare class SessionAffinityConfig {
    constructor();
    constructor(spec: Pick<SessionAffinityConfig, "clientIP">);

	/**
     * ClientIP contains the configurations of Client IP based session affinity.
     */
    clientIP?: ClientIPConfig

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}