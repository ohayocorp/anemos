import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";

export const componentType = "set-ingress-class-name";

export class Options {
    /** Name of the Ingress class to be set on Ingress resources. */
    ingressClassName: string;
    
    /** Predicate to filter which documents to modify. All ingress documents will be modified if not specified. */
    predicate?: (context: BuildContext, document: Document) => boolean;

    constructor(ingressClassName: string, predicate?: (context: BuildContext, document: Document) => boolean) {
        this.ingressClassName = ingressClassName;
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

        this.addAction(steps.sanitize, this.sanitize);
        this.addAction(steps.modify, this.modify);
    }

    sanitize = (_: BuildContext) => {
        if (!this.options.ingressClassName) {
            throw new Error("ingressClassName option must be specified.");
        }
    }

    modify = (context: BuildContext) => {
        for (const document of context.getAllDocuments()) {
            if (!document.asIngress()) {
                continue;
            }

            if (this.options.predicate && !this.options.predicate(context, document)) {
                continue;
            }

            document.spec ??= {};
            document.spec.ingressClassName = this.options.ingressClassName;
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
         * Set the Ingress class name on Ingress resources during the {@link steps.modify} step.
         */
        setIngressClassName(options: Options): Component;
        
        /**
         * Set the Ingress class name on Ingress resources during the {@link steps.modify} step.
         */
        setIngressClassName(ingressClassName: String): Component;
    }
}

Builder.prototype.setIngressClassName = function (this: Builder, arg: Options | String): Component {    
    if (typeof arg === "string") {
        return add(this, new Options(arg));
    } else if (typeof arg === "object") {
        return add(this, arg as Options);
    }

    throw new Error("Invalid argument");
}