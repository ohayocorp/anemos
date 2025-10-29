import { KubernetesResource } from "./kubernetesResourceInfo";
import { Step } from "./step";

/**
 * Built-in steps in execution order:
 * - {@link populateKubernetesResources}            -> 1
 * - {@link sanitize}                               -> 2
 * - {@link generateResources}                      -> 5
 * - {@link generateResourcesBasedOnOtherResources} -> 5,1
 * - {@link modify}                                 -> 6
 * - {@link specifyProvisionerDependencies}         -> 7
 * - {@link output}                                 -> 99
 */

/**
 * Use this step to populate {@link KubernetesResource} resources so that other components can rely on this
 * information to modify existing resources or generate extra resources.
 * E.g. when ServiceMonitor is added via this func, other components can generate ServiceMonitor resources
 * to monitor the services.
 */
export const populateKubernetesResources: Step;

/** Sanitize the options and the component properties in this step. */
export const sanitize: Step;

/** Use this step to generate documents and additional files. */
export const generateResources: Step;

/**
 * Use this step to generate documents and additional files based on other documents and additional files
 * that were generated in the {@link generateResources} step.
 */
export const generateResourcesBasedOnOtherResources: Step;

/**
 * Use this step to modify the generated documents, e.g. set labels, annotations, etc.
 */
export const modify: Step;

/** Specify provisioner dependencies in this step. */
export const specifyProvisionerDependencies: Step;

/** Write the outputs, e.g. documents and additional files in this step. */
export const output: Step;
