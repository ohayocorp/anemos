import * as anemos from "@ohayocorp/anemos";

export const componentType = "set-default-provisioner-dependencies";

export class Component extends anemos.Component {
    constructor() {
        super();

        this.addAction(anemos.steps.specifyProvisionerDependencies, this.specifyProvisionerDependencies);
    }

    specifyProvisionerDependencies = (context: anemos.BuildContext) => {
        // Find document groups that depend on other resources in other document groups such as
        // namespaces or CRDs and add provisioner dependencies accordingly.
        const resources: Resource[] = [];

        for (const document of context.getAllDocuments()) {
            const apiVersion = document.apiVersion;
            const kind = document.kind;

            if (!apiVersion || !kind) {
                continue;
            }

            if (apiVersion === "v1" && kind === "Namespace") {
                resources.push({
                    apiVersion,
                    kind,
                    document,
                    checkDependency: checkNamespace
                });
            }

            if (apiVersion === "apiextensions.k8s.io/v1" && kind === "CustomResourceDefinition") {
                resources.push({
                    apiVersion,
                    kind,
                    document,
                    checkDependency: checkCRD
                });
            }
        }

        // Check all documents with all collected resources.
        for (const document of context.getAllDocuments()) {
            for (const resource of resources) {
                if (!resource.checkDependency(resource, document)) {
                    continue;
                }

                const prerequisite = resource.document.group;
                const dependent = document.group;

                // Don't add self-dependencies in the same group.
                if (prerequisite && dependent && prerequisite !== dependent) {
                    prerequisite.provisionBefore(dependent);
                }
            }
        }
    }
}

function checkNamespace(resource: Resource, document: anemos.Document): boolean {
    const namespace = document.metadata?.namespace;
    if (!namespace) {
        return false;
    }

	return resource.document.metadata?.name === namespace
}

function checkCRD(resource: Resource, document: anemos.Document): boolean {
    const apiVersion = document.apiVersion;
    const kind = document.kind;

    if (!apiVersion || !kind) {
        return false;
    }

    // Compare kinds.
    const resourceKind = resource.document.spec?.names?.kind;
    if (!resourceKind || resourceKind !== kind) {
        return false
    }

    const group = resource.document.spec?.group;
    if (!group) {
        return false;
    }

    // CRD can have multiple versions.
    for (const element of resource.document.spec?.versions || []) {
        const version = element.name;

        if (version && apiVersion === `${group}/${version}`) {
            return true;
        }
    }

    return false;
}

class Resource {
    apiVersion!: string;
    kind!: string;
    document!: anemos.Document;
    checkDependency!: (resource: Resource, document: anemos.Document) => boolean;
}
