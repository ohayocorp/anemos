import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { DocumentGroup } from "@ohayocorp/anemos/documentGroup";
import { Document } from "@ohayocorp/anemos/document";
import { Step } from "@ohayocorp/anemos/step";
import * as steps from "@ohayocorp/anemos/steps";

export const componentType = "collect-namespaces";

export class Options {
    /** Path of the document group in which namespaces will be collected. Default value is 'namespaces'. */
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
        this.addAction(new Step("Collect Namespaces", [...steps.modify.numbers, 1]), this.modify);
    }

    sanitize = (_: BuildContext) => {
        this.options.documentGroup ??= "namespaces";
    }

    modify = (context: BuildContext) => {
        const namespaces = new DocumentGroup(this.options.documentGroup!);
        const documentGroupsToRemove: DocumentGroup[] = [];

        for (const documentGroup of context.getDocumentGroups()) {
            const documentsToMove: Document[] = documentGroup.documents.filter(document => document.isNamespace());

            if (documentsToMove.length == 0) {
                continue;
            }

            for (const document of documentsToMove) {
                documentGroup.removeDocument(document);
                namespaces.addDocument(document);
                
                // Use the default naming scheme for the file path.
                document.setPath(null);
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

export function add(builder: Builder, options?: Options): Component {
    const component = new Component(options);
    builder.addComponent(component);

    return component;
}

declare module "@ohayocorp/anemos" {
    export interface Builder {
        /**
         * Adds a {@link Component} that collects namespace definitions from all
         * the document groups and moves them into a new document group after the {@link steps.modify} step.
         * @param options Options for collecting namespaces.
         */
        collectNamespaces(options?: Options): Component;
    }
}

Builder.prototype.collectNamespaces = function (this: Builder, options?: Options): Component {
    return add(this, options);
}