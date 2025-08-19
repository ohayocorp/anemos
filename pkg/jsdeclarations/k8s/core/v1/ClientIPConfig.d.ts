// Auto generated code; DO NOT EDIT.



/**
 * ClientIPConfig represents the configurations of Client IP based session affinity.
 * 
 */
export declare class ClientIPConfig {
    constructor();
    constructor(spec: ClientIPConfig);

	/**
     * TimeoutSeconds specifies the seconds of ClientIP type session sticky time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP". Default value is 10800(for 3 hours).
     * 
     */
    timeoutSeconds?: number
}