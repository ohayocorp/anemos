import * as anemos from "@ohayocorp/anemos"

export const componentType = "collect-namespaces";

export class Options {
    /** Path of the document group in which namespaces will be collected. Default value is 'namespaces'. */
    documentGroup?: string;
}

export class Component extends anemos.Component {
    options: Options;

    constructor(options?: Options) {
        super();

        this.options = options ?? {};

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(anemos.steps.sanitize, this.sanitize);
        this.addAction(new anemos.Step("Collect Namespaces", [...anemos.steps.modify.numbers, 1]), this.modify);
    }

    sanitize = (_: anemos.BuildContext) => {
        this.options.documentGroup ??= "namespaces";
    }

    modify = (context: anemos.BuildContext) => {
        const namespaces = new anemos.DocumentGroup(this.options.documentGroup!);
        const documentGroupsToRemove: anemos.DocumentGroup[] = [];

        for (const documentGroup of context.getDocumentGroups()) {
            const documentsToMove: anemos.Document[] = documentGroup.documents.filter(document => document.isNamespace());

            if (documentsToMove.length == 0) {
                continue;
            }

            for (const document of documentsToMove) {
                documentGroup.removeDocument(document);
                namespaces.addDocument(document);
                
                // Set the path to the document's name or use the default naming scheme if name is not available.
                if (document.metadata?.name) {
                    document.setPath(`${document.metadata.name}.yaml`);
                } else {
                    document.setPath(null);
                }
            }

            if (documentGroup.documents.length == 0) {
                documentGroupsToRemove.push(documentGroup);
            }
        }

        for (const documentGroup of documentGroupsToRemove) {
            context.removeDocumentGroup(documentGroup);
        }

        if (namespaces.documents.length > 0) {
            context.addDocumentGroup(namespaces);
        }
    }
}

export function add(builder: anemos.Builder, options?: Options): Component {
    const component = new Component(options);
    builder.addComponent(component);

    return component;
}

declare module "@ohayocorp/anemos" {
    export interface Builder {
        /**
         * Adds a {@link Component} that collects Custom Resource Definitions (CRDs) from all
         * the document groups and moves them into a new document group after the {@link steps.modify} step.
         * @param options Options for collecting namespaces.
         */
        collectNamespaces(options?: Options): Component;
    }
}

anemos.Builder.prototype.collectNamespaces = function (this: anemos.Builder, options?: Options): Component {
    return add(this, options);
}