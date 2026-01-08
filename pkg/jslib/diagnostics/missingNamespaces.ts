import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { DiagnosticMetadata, error, linting, warning } from "@ohayocorp/anemos/diagnostic";
import { KubernetesResource } from "@ohayocorp/anemos/kubernetesResourceInfo";

export const componentType = "diagnostics/missing-namespaces";

export const diagnosticMetadata: DiagnosticMetadata = {
    id: "missing-namespaces",
    name: "Missing Namespaces",
    description: `Not having a namespace can lead to unexpected results when applying resources to a cluster since resources will be applied in the current namespace of the kube-context.`,
    severity: warning,
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
        for (const document of context.getAllDocuments()) {
            const apiVersion = document.apiVersion;
            const kind = document.kind;

            if (!apiVersion || !kind) {
                continue;
            }

            if (!context.kubernetesResourceInfo.isNamespaced(apiVersion, kind)) {
                continue;
            }

            if (!document.metadata?.namespace) {
                context.addDiagnostic({
                    metadata: diagnosticMetadata,
                    message: ``,
                    document: document,
                });
            }
        }
    }
}

export function add(builder: Builder): Component {
    const component = new Component();
    builder.addComponent(component);

    return component;
}