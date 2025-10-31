import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";

export type Predicate = (document: Document, context: BuildContext) => boolean;

export const componentType = "set-annotations";

export class Options {
    /** Annotations to set in documents. */
    annotations: Record<string, string>;
    
    /**
     * Predicate to filter which documents to modify. All documents will be modified if not specified.
     */
    predicate?: Predicate;

    constructor(annotations: Record<string, string>, predicate?: Predicate) {
        this.annotations = annotations;
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
        const annotations = this.options.annotations;

        if (Object.keys(annotations).length === 0) {
            return;
        }

        for (const document of context.getAllDocuments()) {
            if (this.options.predicate && !this.options.predicate(document, context)) {
                continue;
            }

            document.setAnnotations(annotations);
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
         * Sets the given annotations in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setAnnotations(options: Options): Component;
        
        /**
         * Sets the given annotations in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setAnnotations(annotations: Record<string, string>, predicate?: Predicate): Component;
        
        /**
         * Sets the given annotation in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setAnnotations(key: string, value: string, predicate?: Predicate): Component;
    }
}

Builder.prototype.setAnnotations = function (
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