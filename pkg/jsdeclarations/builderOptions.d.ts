import { KubernetesResource } from "./kubernetesResourceInfo"

/**
 * Contains common options that are used by all components.
 */
export declare class BuilderOptions {
    constructor(kubernetesCluster: KubernetesCluster, environment: Environment, outputConfiguration?: OutputConfiguration);

    environment: Environment
    kubernetesCluster: KubernetesCluster
    outputConfiguration?: OutputConfiguration
}

/**
 * Contains information about the target environment.
 */
export declare class Environment {
    constructor(name: string, type: EnvironmentType);

    name: string
    type: EnvironmentType
}

/**
 * Predefined types for the environment. These are used to determine the default configurations of the components.
 * For example, in the development environment, some components may disable high availability.
 */
export declare enum EnvironmentType {
    Development,
    Testing,
    Production,
}

/**
 * Contains information about the target Kubernetes cluster.
 */
export declare class KubernetesCluster {
    constructor(version: Version, distribution: KubernetesDistribution, additionalResources?: KubernetesResource[]);

    distribution: KubernetesDistribution
    version: Version
    additionalResources?: KubernetesResource[]
}

/**
 * Predefined types for the Kubernetes distribution. These are used to determine the behavior of some components.
 */
export declare enum KubernetesDistribution {
    AKS,
    EKS,
    GKE,
    K3S,
    Kubeadm,
    MicroK8S,
    Minikube,
    OpenShift
}

/**
 * Specifies the output directory structure.
 */
export declare class OutputConfiguration {
    /** Default value is "output" under the current working directory. */
    outputPath?: string
}

/**
 * Represents a single semantic version.
 */
export declare class Version {
    constructor(version: string);
}
