import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { DiagnosticMetadata, error, linting } from "@ohayocorp/anemos/diagnostic";
import { Container } from "@ohayocorp/anemos/k8s/core/v1";
import { Quantity } from "@ohayocorp/anemos/quantity";
import { Document } from "@ohayocorp/anemos/document";

export const componentType = "diagnostics/limit-lower-than-request";

export const diagnosticMetadata: DiagnosticMetadata = {
    id: "limit-lower-than-request",
    name: "Limit Lower Than Request",
    description: `Having a limit lower than the request is an error and Kubernetes will not allow it.`,
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
        for (const document of context.getAllDocuments()) {
            if (!document.asWorkload()) {
                continue;
            }

            const containers = document.getContainers() ?? [];
            const initContainers = document.getInitContainers() ?? [];

            this.visitContainers(context, document, containers);
            this.visitContainers(context, document, initContainers);
        }
    }

    private visitContainers = (context: BuildContext, document: Document, containers: Container[]) => {
        for (const container of containers) {
            const requests = container.resources?.requests;
            const limits = container.resources?.limits;

            if (!requests || !limits) {
                continue;
            }

            const requestNames = Object.keys(requests);
            const limitNames = Object.keys(limits);

            const intersection = requestNames.filter(name => limitNames.includes(name));

            for (const name of intersection) {
                const requestValue = requests[name];
                const limitValue = limits[name];

                if (!requestValue || !limitValue) {
                    continue;
                }

                const requestQuantity = new Quantity(requestValue.toString());
                const limitQuantity = new Quantity(limitValue.toString());

                if (requestQuantity.compare(limitQuantity) > 0) {
                    context.addDiagnostic({
                        metadata: diagnosticMetadata,
                        message: `Limit for ${name} (${limitQuantity}) is lower than request (${requestQuantity}) in container ${container.name} (**${document.fullPath()}**)`,
                        document: document
                    })
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