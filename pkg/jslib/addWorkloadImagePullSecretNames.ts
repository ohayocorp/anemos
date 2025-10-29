import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";
import { Workload } from "@ohayocorp/anemos/documentExtensionsWorkload";

export const componentType = "add-workload-image-pull-secret-names";

export class Options {
    /** Name of the image pull secret to add to workloads. */
    imagePullSecretNames: string[];
    
    /** Predicate to filter which documents to modify. All workload documents will be modified if not specified. */
    predicate?: (context: BuildContext, document: Document) => boolean;

    constructor(imagePullSecretNames: string[], predicate?: (context: BuildContext, document: Document) => boolean) {
        this.imagePullSecretNames = imagePullSecretNames;
        this.predicate = predicate;
    }
}

export class Component extends AnemosComponent {
    options: Options;

    constructor(options: Options) {
        super();

        this.options = options;

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.modify, this.modify);
    }

    modify = (context: BuildContext) => {
        for (const document of context.getAllDocuments()) {
            if (!document.asWorkload()) {
                continue;
            }

            if (this.options.predicate && !this.options.predicate(context, document)) {
                continue;
            }

            this.addImagePullSecrets(document);
        }
    }

    addImagePullSecrets = (workload: Workload) => {
        if (this.options.imagePullSecretNames.length === 0) {
            return;
        }

        const spec = workload.getWorkloadSpec();
        if (!spec) {
            return;
        }

        const imagePullSecrets = spec.imagePullSecrets ??= [];

        for (const name of this.options.imagePullSecretNames) {
            const imagePullSecret = imagePullSecrets.find(x => x.name === name);
            // Avoid adding duplicate image pull secret names.
            if (imagePullSecret) {
                continue;
            }

            imagePullSecrets.push({ name: name });
        }
    }
}

export function add(builder: Builder, options: Options): Component {
    const component = new Component(options);
    builder.addComponent(component);

    return component;
}

declare module "@ohayocorp/anemos" {
    export interface Builder {
        /**
         * Adds given image pull secret names to the pod spec of all workload resources.
         */
        addWorkloadImagePullSecretNames(options: Options): Component;
        
        /**
         * Adds given image pull secret names to the pod spec of all workload resources.
         */
        addWorkloadImagePullSecretNames(imagePullSecretNames: string[]): Component;
    }
}

Builder.prototype.addWorkloadImagePullSecretNames = function (this: Builder, arg: Options | string[]): Component {
    if (Array.isArray(arg)) {
        return add(this, new Options(arg));
    } else if (typeof arg === "object") {
        return add(this, arg as Options);
    }

    throw new Error("Invalid argument");
}