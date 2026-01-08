import * as anemos from "@ohayocorp/anemos";
import * as duplicateResources from './duplicateResources';
import * as limitLowerThanRequest from './limitLowerThanRequest';
import * as missingLabels from './missingLabels';
import * as missingNamespaces from './missingNamespaces';
import * as missingResourceRequirements from './missingResourceRequirements';
import * as runAsRoot from './runAsRoot';

export * as duplicateResources from './duplicateResources';
export * as limitLowerThanRequest from './limitLowerThanRequest';
export * as missingLabels from './missingLabels';
export * as missingNamespaces from './missingNamespaces';
export * as missingResourceRequirements from './missingResourceRequirements';
export * as runAsRoot from './runAsRoot';

export function addDefaultDiagnostics(builder: anemos.Builder) {
    duplicateResources.add(builder);
    limitLowerThanRequest.add(builder);
    missingLabels.add(builder);
    missingNamespaces.add(builder);
    missingResourceRequirements.add(builder);
    runAsRoot.add(builder);
}