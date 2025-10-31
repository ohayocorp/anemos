import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";

export type Predicate = (document: Document, context: BuildContext) => boolean;

export const componentType = "set-labels";

export class Options {
    /** Labels to set in documents. */
    labels: Record<string, string>;
    
    /**
     * Predicate to filter which documents to modify. All documents will be modified if not specified.
     */
    predicate?: Predicate;

    constructor(labels: Record<string, string>, predicate?: Predicate) {
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
        setLabels(labels: Record<string, string>, predicate?: Predicate): Component;
        
        /**
         * Sets the given label in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setLabels(key: string, value: string, predicate?: Predicate): Component;
    }
}

Builder.prototype.setLabels = function (
    this: Builder,
    first: Options | Record<string, string> | string,
    second?: string | Predicate,
    third?: Predicate
): Component {
    if (typeof first === "string") {
        if (typeof second !== "string") {
            throw new Error("Invalid argument expected string for value");
        }

        if (third && typeof third !== "function") {
            throw new Error("Invalid argument expected function for predicate");
        }

        return add(this, new Options({ [first]: second }, third));
    }

    if (typeof first !== "object") {
        throw new Error("Invalid argument");
    }

    if (Object.getOwnPropertyNames(first).every(property => typeof Object.getOwnPropertyDescriptor(first, property)?.value === "string")) {
        if (second && typeof second !== "function") {
            throw new Error("Invalid argument expected function for predicate");
        }

        return add(this, new Options(first as Record<string, string>, second as Predicate));
    }

    return add(this, first as Options);
}