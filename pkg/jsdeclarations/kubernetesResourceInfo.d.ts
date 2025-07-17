/**
 * Represents a Kubernetes API resource, such as Pod, Deployment, Service, etc.
 * It contains the API version, kind, and whether the resource is namespaced.
 */
export declare class KubernetesResource {
    constructor(apiVersion: string, kind: string, isNamespaced: boolean);
    
    /** The API version of the resource, e.g. "v1", "apps/v1". */
    apiVersion: string;

    /** The kind of the resource, e.g. "Pod", "Deployment". */
    kind: string;

    /** Indicates whether the resource is namespaced (true) or cluster-scoped (false). */
    isNamespaced: boolean;
}

/**
 * Contains all the API resources defined in the target cluster and enables listing them
 * and querying their existence.
 */
export declare class KubernetesResourceInfo {
    private constructor();

    /** Adds the given API resource to the available resources list. */
    addResource(resource: KubernetesResource): void;

    /** Returns true if the given API resource exists in the target cluster. */
    contains(apiVersion: string, kind: string): boolean;

    /** Returns true if the given kind exists in the target cluster. This ignores the apiVersion field. */
    containsKind(kind: string): boolean;

    /**
     * Returns true if the given API resource is namespaced. E.g. returns true for v1/Pod,
     * false for rbac.authorization.k8s.io/v1/ClusterRole.
     */
    isNamespaced(apiVersion: string, kind: string): boolean;
}