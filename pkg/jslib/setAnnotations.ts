import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";

export const componentType = "set-annotations";

export class Options {
    /** Annotations to set in documents. */
    annotations: Record<string, string>;
    
    /**
     * Predicate to filter which documents to modify. All documents will be modified if not specified.
     */
    predicate?: (document: Document, context: BuildContext) => boolean;

    constructor(
        annotations: Record<string, string>,
        predicate?: (document: Document, context: BuildContext) => boolean
    ) {
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
        setAnnotations(
            environmentVariables: Record<string, string>,
            predicate?: (document: Document, context: BuildContext) => boolean
        ): Component;
        
        /**
         * Sets the given annotation in all documents. It is possible to filter which documents to modify
         * by specifying a predicate.
         */
        setAnnotation(
            key: string,
            value: string,
            predicate?: (document: Document, context: BuildContext) => boolean
        ): Component;
    }
}

Builder.prototype.setAnnotations = function (
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

Builder.prototype.setAnnotation = function (
    this: Builder,
    key: string,
    value: string,
    predicate?: (document: Document, context: BuildContext) => boolean
): Component {
    return add(this, new Options({ [key]: value }, predicate));
}