// Auto generated code; DO NOT EDIT.



/**
 * ServiceAccountSubject holds detailed information for service-account-kind subject.
 * 
 */
export declare class ServiceAccountSubject {
    constructor();
    constructor(spec: ServiceAccountSubject);

	/**
     * `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
     * 
     */
    name: string

	/**
     * `namespace` is the namespace of matching ServiceAccount objects. Required.
     * 
     */
    namespace: string
}