import * as anemos from "@ohayocorp/anemos";
import * as kubernetesResources from './kubernetesResources';
import * as resourceRequirements from './resourceRequirements';

export * as kubernetesResources from './kubernetesResources';
export * as resourceRequirements from './resourceRequirements';

export function addDefaultReports(builder: anemos.Builder) {
    kubernetesResources.add(builder);
    resourceRequirements.add(builder);
}