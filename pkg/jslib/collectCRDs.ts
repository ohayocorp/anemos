import * as anemos from "@ohayocorp/anemos"

export const componentType = "collect-crds";

export class Options {
    /** Path of the document group in which CRDs will be collected. Default value is 'crds'. */
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
        this.addAction(new anemos.Step("Collect CRDs", [...anemos.steps.modify.numbers, 1]), this.modify);
    }

    sanitize = (_: anemos.BuildContext) => {
        this.options.documentGroup ??= "crds";
    }

    modify = (context: anemos.BuildContext) => {
        const crds = new anemos.DocumentGroup(this.options.documentGroup!);
        const documentGroupsToRemove: anemos.DocumentGroup[] = [];

        for (const documentGroup of context.getDocumentGroups()) {
            const documentsToMove: anemos.Document[] = documentGroup.documents.filter(document => document.isCustomResourceDefinition());

            if (documentsToMove.length == 0) {
                continue;
            }

            for (const document of documentsToMove) {
                documentGroup.removeDocument(document);
                crds.addDocument(document);
                
                // Clear the path to enable default naming behavior.
                document.setPath(null);
            }

            if (documentGroup.documents.length == 0) {
                documentGroupsToRemove.push(documentGroup);
            }
        }

        for (const documentGroup of documentGroupsToRemove) {
            context.removeDocumentGroup(documentGroup);
        }

        if (crds.documents.length > 0) {
            context.addDocumentGroup(crds);
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
         * @param options Options for collecting CRDs.
         */
        collectCRDs(options?: Options): Component;
    }
}

anemos.Builder.prototype.collectCRDs = function (this: anemos.Builder, options?: Options): Component {
    return add(this, options);
}