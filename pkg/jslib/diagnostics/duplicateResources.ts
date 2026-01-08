import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { Document } from "@ohayocorp/anemos/document";
import { DiagnosticMetadata, error, linting } from "@ohayocorp/anemos/diagnostic";

export const componentType = "diagnostics/duplicate-resources";

export const diagnosticMetadata: DiagnosticMetadata = {
    id: "duplicate-resources",
    name: "Duplicate Resources",
    description: `Multiple definitions of the same resource may override each other and cause unexpected behavior.`,
    severity: error,
    categories: [linting]
};

export class Component extends AnemosComponent {
    constructor() {
        super();

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.diagnose, this.diagnose);
    }

    diagnose = (context: BuildContext) => {
        const allDocuments = new Map<string, Document[]>();

        for (const document of context.getAllDocuments()) {
            const apiVersion = document.apiVersion;
            const kind = document.kind;
            const name = document.metadata?.name;
            const namespace = document.metadata?.namespace;

            if (!apiVersion || !kind || !name) {
                continue;
            }

            const key = `${apiVersion}#${kind}#${namespace ?? ""}#${name}`;
            const documents = allDocuments.get(key) ?? [];
            documents.push(document);
            allDocuments.set(key, documents);
        }

        for (const documents of allDocuments.values()) {
            if (documents.length > 1) {
                const firstDocument = documents[0];

                const apiVersion = firstDocument.apiVersion;
                const kind = firstDocument.kind;
                const name = firstDocument.metadata?.name;
                const namespace = firstDocument.metadata?.namespace;

                const message = `Duplicate resource detected: ${apiVersion}/${kind} ${name}` + (namespace ? ` in namespace ${namespace}` : '');

                for (const document of documents) {
                    context.addDiagnostic({
                        metadata: diagnosticMetadata,
                        message: message,
                        document: document,
                    });
                }
            }
        }
    }
}

export function add(builder: Builder): Component {
    const component = new Component();
    builder.addComponent(component);

    return component;
}