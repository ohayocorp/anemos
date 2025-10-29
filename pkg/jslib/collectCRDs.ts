import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { DocumentGroup } from "@ohayocorp/anemos/documentGroup";
import { Document } from "@ohayocorp/anemos/document";
import { Step } from "@ohayocorp/anemos/step";
import * as steps from "@ohayocorp/anemos/steps";

export const componentType = "collect-crds";

export class Options {
    /** Path of the document group in which CRDs will be collected. Default value is 'crds'. */
    documentGroup?: string;
}

export class Component extends AnemosComponent {
    options: Options;

    constructor(options?: Options) {
        super();

        this.options = options ?? {};

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.sanitize, this.sanitize);
        this.addAction(new Step("Collect CRDs", [...steps.modify.numbers, 1]), this.modify);
    }

    sanitize = (_: BuildContext) => {
        this.options.documentGroup ??= "crds";
    }

    modify = (context: BuildContext) => {
        const crds = new DocumentGroup(this.options.documentGroup!);
        const documentGroupsToRemove: DocumentGroup[] = [];

        for (const documentGroup of context.getDocumentGroups()) {
            const documentsToMove: Document[] = documentGroup.documents.filter(document => document.isCRD());

            if (documentsToMove.length == 0) {
                continue;
            }

            for (const document of documentsToMove) {
                documentGroup.removeDocument(document);
                crds.addDocument(document);

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

        if (crds.documents.length > 0) {
            context.addDocumentGroup(crds);
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
         * Adds a {@link Component} that collects Custom Resource Definitions (CRDs) from all
         * the document groups and moves them into a new document group after the {@link steps.modify} step.
         * @param options Options for collecting CRDs.
         */
        collectCRDs(options?: Options): Component;
    }
}

Builder.prototype.collectCRDs = function (this: Builder, options?: Options): Component {
    return add(this, options);
}