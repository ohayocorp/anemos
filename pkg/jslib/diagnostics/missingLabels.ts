import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { DiagnosticMetadata, linting, warning } from "@ohayocorp/anemos/diagnostic";
import { KubernetesResource } from "@ohayocorp/anemos/kubernetesResourceInfo";

export const componentType = "diagnostics/missing-labels";

export const diagnosticMetadata: DiagnosticMetadata = {
    id: "missing-labels",
    name: "Missing Labels",
    description: `Labels are used to organize resources and are used by selectors to match resources. Using standard labels can help organizing resources and make it easier to manage them using tools like kubectl.`,
    severity: warning,
    categories: [linting]
};

export class Options {
    labelsToCheck?: string[];
    resourcesToExclude?: KubernetesResource[];
}

export class Component extends AnemosComponent {
    private options: Options;

    constructor(options?: Options) {
        super();

        this.options = options ?? {};

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.sanitize, this.sanitize);
        this.addAction(steps.diagnose, this.diagnose);
    }

    sanitize = (context: BuildContext) => {
        this.options.labelsToCheck ??= [
            "app.kubernetes.io/name",
            "app.kubernetes.io/instance",
        ];

        this.options.resourcesToExclude ??= [
            { apiVersion: "v1", kind: "Namespace", isNamespaced: false },
            { apiVersion: "apiextensions.k8s.io/v1", kind: "CustomResourceDefinition", isNamespaced: false },
        ];
    }

    diagnose = (context: BuildContext) => {
        const checkLabels = (labels?: Record<string, string>): string[] => {
            const missingLabels: Set<string> = new Set();

            for (const label of this.options.labelsToCheck!) {
                if (!labels?.[label]) {
                    missingLabels.add(label);
                }
            }

            return Array.from(missingLabels).sort();
        }

        for (const document of context.getAllDocuments()) {
            const apiVersion = document.apiVersion;
            const kind = document.kind;

            if (!apiVersion || !kind) {
                continue;
            }

            if (this.options.resourcesToExclude!.some(resource => resource.apiVersion === apiVersion && resource.kind === kind)) {
                continue;
            }

            const labels = document.metadata?.labels;
            const missingLabels = checkLabels(labels);

            for (const missingLabel of missingLabels) {
                context.addDiagnostic({
                    metadata: diagnosticMetadata,
                    message: `${missingLabel}`,
                    document: document,
                });
            }

            if (document.asWorkload()) {
                const missingPodTemplateLabels = checkLabels(document.getWorkloadLabels());

                for (const missingLabel of missingPodTemplateLabels) {
                    context.addDiagnostic({
                        metadata: diagnosticMetadata,
                        message: `${missingLabel} *(pod template)*`,
                        document: document,
                    });
                }
            }
        }
    }
}

export function add(builder: Builder, options?: Options): Component {
    const component = new Component(options);
    builder.addComponent(component);

    return component;
}