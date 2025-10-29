import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { DocumentGroup } from "@ohayocorp/anemos/documentGroup";
import * as k8s from "@ohayocorp/anemos/k8s";
import * as steps from "@ohayocorp/anemos/steps";

const systemNamespaces = [
    "kube-system",
    "kube-public",
    "kube-node-lease",
    "default",
];

export const componentType = "create-referenced-namespaces";

export class Options {
    /** Path of the document group in which created namespaces will belong. Default value is 'namespaces'. */
    documentGroup?: string;

    /**
     * A predicate function to filter namespaces. Returns true if the namespace should be included.
     * If not specified, all namespaces except the default Kubernetes namespaces will be included.
     */
    predicate?: (namespace: string) => boolean
}

export class Component extends AnemosComponent {
    options: Options;

    constructor(options?: Options) {
        super();

        this.options = options ?? {};

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.sanitize, this.sanitize);
        this.addAction(steps.generateResourcesBasedOnOtherResources, this.generateNamespaces);
    }

    sanitize = (_: BuildContext) => {
        this.options.documentGroup ??= "namespaces";
    }

    generateNamespaces = (context: BuildContext) => {
        const predicate = this.options.predicate;
        const namespaces = new Set<string>();
        const existingNamespaces = new Set<string>();

        for (const document of context.getAllDocuments()) {
            if (document.isNamespace() && document.metadata?.name) {
                existingNamespaces.add(document.metadata.name);
            }

            const namespace = document.metadata?.namespace;
            if (!namespace) {
                continue
            }

            if (systemNamespaces.includes(namespace)) {
                continue;
            }

            if (predicate && !predicate(namespace)) {
                continue;
            }

            namespaces.add(namespace);
        }

        for (const namespace of existingNamespaces) {
            namespaces.delete(namespace);
        }

        if (namespaces.size === 0) {
            return;
        }

        const documentGroup = new DocumentGroup(this.options.documentGroup!);
        context.addDocumentGroup(documentGroup);

        for (const namespace of namespaces) {
            documentGroup.addDocument(new k8s.Namespace({
                metadata: {
                    name: namespace
                }
            }));
        }
    }
}

export function add(builder: Builder, options?: Options): Component {
    const component = new Component(options);
    builder.addComponent(component);

    return component;
}

declare module "@ohayocorp/anemos" {
    export interface Builder {
        /**
         * Create namespace manifests for all the namespaces that are referenced by other resources.
         * @param options The options for creating referenced namespaces.
         */
        createReferencedNamespaces(options?: Options): Component;
    }
}

Builder.prototype.createReferencedNamespaces = function (this: Builder, options?: Options): Component {
    return add(this, options);
}