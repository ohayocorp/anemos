import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { Document } from "@ohayocorp/anemos/document";
import { DiagnosticMetadata, security, warning } from "@ohayocorp/anemos/diagnostic";
import { Container, PodSpec } from "@ohayocorp/anemos/k8s/core/v1";

export const componentType = "diagnostics/run-as-root";

export const diagnosticMetadata: DiagnosticMetadata = {
    id: "run-as-root",
    name: "Run As Root",
    description: `Running containers as root can be a security risk.`,
    severity: warning,
    categories: [security]
};

export class Component extends AnemosComponent {
    constructor() {
        super();

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.diagnose, this.diagnose);
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
            const runAsUser = this.getRunAsUser(document.spec, container);
            if (runAsUser && runAsUser !== 0) {
                continue;
            }

            const runAsNonRoot = this.getRunAsNonRoot(document.spec, container);
            if (runAsNonRoot) {
                if (runAsUser === 0) {
                    context.addDiagnostic({
                        metadata: diagnosticMetadata,
                        message: `**${container.name}** runAsNonRoot is true, but runAsUser is set to 0`,
                        document: document,
                    });

                    continue;
                }
            }

            context.addDiagnostic({
                metadata: diagnosticMetadata,
                message: `**${container.name}** runAsNonRoot is not set to true`,
                document: document,
            });
        }
    }

    private getRunAsUser = (spec: PodSpec, container: Container): number | undefined => {
        if (container.securityContext?.runAsUser) {
            return container.securityContext.runAsUser;
        }

        if (spec.securityContext?.runAsUser) {
            return spec.securityContext.runAsUser;
        }

        return undefined;
    }

    private getRunAsNonRoot = (spec: PodSpec, container: Container): boolean | undefined => {
        if (container.securityContext?.runAsNonRoot) {
            return container.securityContext.runAsNonRoot;
        }

        if (spec.securityContext?.runAsNonRoot) {
            return spec.securityContext.runAsNonRoot;
        }

        return undefined;
    }
}

export function add(builder: Builder): Component {
    const component = new Component();
    builder.addComponent(component);

    return component;
}