import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";

export const componentType = "set-labels";

export class Options {
    /** Labels to set in documents. */
    labels: Record<string, string>;
    
    /**
     * Predicate to filter which documents to modify. All documents will be modified if not specified.
     */
    predicate?: (document: Document, context: BuildContext) => boolean;

    constructor(
        labels: Record<string, string>,
        predicate?: (document: Document, context: BuildContext) => boolean
    ) {
        this.labels = labels;
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
        const labels = this.options.labels;

        if (Object.keys(labels).length === 0) {
            return;
        }

        for (const document of context.getAllDocuments()) {
            if (this.options.predicate && !this.options.predicate(document, context)) {
                continue;
            }

            document.setLabels(labels);
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
         * Sets the given labels in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setLabels(options: Options): Component;
        
        /**
         * Sets the given labels in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setLabels(
            environmentVariables: Record<string, string>,
            predicate?: (document: Document, context: BuildContext) => boolean
        ): Component;
        
        /**
         * Sets the given label in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setLabel(
            key: string,
            value: string,
            predicate?: (document: Document, context: BuildContext) => boolean
        ): Component;
    }
}

Builder.prototype.setLabels = function (
    this: Builder,
    arg: Options | Record<string, string>,
    predicate?: (document: Document, context: BuildContext) => boolean
): Component {
    if (typeof arg !== "object") {
        throw new Error("Invalid argument");
    }

    if (Object.getOwnPropertyNames(arg).every(property => typeof Object.getOwnPropertyDescriptor(arg, property)?.value === "string")) {
        return add(this, new Options(arg as Record<string, string>, predicate));
    }

    return add(this, arg as Options);
}

Builder.prototype.setLabel = function (
    this: Builder,
    key: string,
    value: string,
    predicate?: (document: Document, context: BuildContext) => boolean
): Component {
    return add(this, new Options({ [key]: value }, predicate));
}