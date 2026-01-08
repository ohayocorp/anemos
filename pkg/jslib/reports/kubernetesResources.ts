import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { Document } from "@ohayocorp/anemos/document";
import { Report } from "@ohayocorp/anemos/report";

export const componentType = "report-kubernetes-resources";

export class Component extends AnemosComponent {
    constructor() {
        super();

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.report, this.report);
    }

    report = (context: BuildContext) => {
        const documentGroups = new Map<string, Document[]>();

        for (const document of context.getAllDocuments()) {
            const apiVersion = document.apiVersion;
            const kind = document.kind;
            const name = document.metadata?.name;

            if (!apiVersion || !kind || !name) {
                continue;
            }

            const key = `${apiVersion}#${kind}`;
            const documents = documentGroups.get(key) ?? [];
            documents.push(document);
            documentGroups.set(key, documents);
        }

        if (documentGroups.size === 0) {
            return;
        }

        // Sort keys to ensure deterministic output order.
        const sortedKeys = Array.from(documentGroups.keys()).sort();

        var content = "# Kubernetes Resources\n\n";

        for (const key of sortedKeys) {
            const documentGroup = documentGroups.get(key)!;
            documentGroup.sort((a, b) => {
                const namespaceA = a.metadata?.namespace ?? "";
                const namespaceB = b.metadata?.namespace ?? "";

                const nameA = a.metadata?.name ?? "";
                const nameB = b.metadata?.name ?? "";

                return namespaceA.localeCompare(namespaceB) || nameA.localeCompare(nameB);
            });

            content += "## " + documentGroup[0].apiVersion + "/" + documentGroup[0].kind + "\n\n";

            for (const document of documentGroup) {
                const name = document.metadata?.name;
                const namespace = document.metadata?.namespace;

                content += `- ${(namespace ? namespace + "/" : "")}${name}\n`;
            }

            content += "\n";
        }

        const report = new Report({ filePath: "kubernetes-resources.md" }, content);
        context.addReport(report);
    }
}

export function add(builder: Builder): Component {
    const component = new Component();
    builder.addComponent(component);

    return component;
}