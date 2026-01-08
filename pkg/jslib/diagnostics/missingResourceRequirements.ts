import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { DiagnosticMetadata, linting, warning } from "@ohayocorp/anemos/diagnostic";
import { Container } from "@ohayocorp/anemos/k8s/core/v1";
import { Document } from "@ohayocorp/anemos";

export const componentType = "diagnostics/missing-resource-requirements";

export const diagnosticMetadata: DiagnosticMetadata = {
    id: "missing-resource-requirements",
    name: "Missing Resource Requirements",
    description: `Resource requirements are used to specify the amount of resources a container needs. This is important for the scheduler to make decisions about where to place the container and how much resources to allocate to it. Missing resource requirements can lead to performance issues and can make it harder to manage the cluster.`,
    severity: warning,
    categories: [linting]
};

export class Options {
    requestsToCheck?: string[];
    limitsToCheck?: string[];
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

    private sanitize = (context: BuildContext) => {
        this.options.requestsToCheck ??= [
            "cpu",
            "memory",
        ];

        this.options.limitsToCheck ??= [
            "cpu",
            "memory",
        ];
    }

    private diagnose = (context: BuildContext) => {
        for (const document of context.getAllDocuments()) {
            if (!document.asWorkload()) {
                continue;
            }

            this.checkResources(context, document, document.getContainers());
            this.checkResources(context, document, document.getInitContainers());
        }
    }

    private checkResources = (context: BuildContext, document: Document, containers?: Container[]) => {
        if (!containers) {
            return;
        }

        for (const container of containers) {
            const resources = container.resources;

            for (const resource of this.options.requestsToCheck!) {
                if (!resources?.requests?.[resource]) {
                    context.addDiagnostic({
                        metadata: diagnosticMetadata,
                        message: `${container.name} *(request)* **${resource}**`,
                        document: document,
                    });
                }
            }

            for (const resource of this.options.limitsToCheck!) {
                if (!resources?.limits?.[resource]) {
                    context.addDiagnostic({
                        metadata: diagnosticMetadata,
                        message: `${container.name} *(limit)* **${resource}**`,
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