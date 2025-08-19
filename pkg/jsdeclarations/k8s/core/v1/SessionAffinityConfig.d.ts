// Auto generated code; DO NOT EDIT.

import { ClientIPConfig } from "./ClientIPConfig"

/**
 * SessionAffinityConfig represents the configurations of session affinity.
 * 
 */
export declare class SessionAffinityConfig {
    constructor();
    constructor(spec: SessionAffinityConfig);

	/**
     * ClientIP contains the configurations of Client IP based session affinity.
     * 
     */
    clientIP?: ClientIPConfig
}