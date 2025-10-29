import { KubernetesResource } from "./kubernetesResourceInfo";
import { EnvironmentType } from "./environmentType";
import { KubernetesDistribution } from "./kubernetesDistribution";

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
 * Contains information about the target Kubernetes cluster.
 */
export declare class KubernetesCluster {
    constructor(version: Version, distribution: KubernetesDistribution, additionalResources?: KubernetesResource[]);

    distribution: KubernetesDistribution
    version: Version
    additionalResources?: KubernetesResource[]
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
