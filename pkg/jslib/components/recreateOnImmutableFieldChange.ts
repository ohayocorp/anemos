import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";

export type Predicate = (context: BuildContext, document: Document) => boolean;

export const componentType = "recreate-on-immutable-field-change";

export class Options {
    /**
     * Predicate to filter which documents to modify. All documents will be modified if not specified.
     */
    predicate?: Predicate;
}

export class Component extends AnemosComponent {
    options: Options;

    constructor(options?: Options) {
        super();

        this.options = options ?? {};

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.modify, this.modify);
    }

    modify = (context: BuildContext) => {
        for (const document of context.getAllDocuments()) {
            if (this.options.predicate && !this.options.predicate(context, document)) {
                continue;
            }

            document.setAnnotation("anemos.sh/recreate-on-immutable-fields-change", "true");
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
         * Sets "anemos.sh/recreate-on-immutable-fields-change" label to "true" in all documents. It is possible
         * to filter which documents to modify by specifying a predicate.
         */
        recreateOnImmutableFieldChange(options?: Options): Component;
        
        /**
         * Sets "anemos.sh/recreate-on-immutable-fields-change" label to "true" in all documents. It is possible
         * to filter which documents to modify by specifying a predicate.
         */
        recreateOnImmutableFieldChange(predicate?: Predicate): Component;
        
        /**
         * Sets "anemos.sh/recreate-on-immutable-fields-change" label to "true" in all Job documents.
         */
        recreateJobsOnImmutableFieldChange(): Component;
    }
}

Builder.prototype.recreateOnImmutableFieldChange = function (this: Builder, first?: Options | Predicate): Component {
    if (typeof first === "function") {
        return add(this, {
            predicate: first
        });
    }

    return add(this, first as Options);
}

Builder.prototype.recreateJobsOnImmutableFieldChange = function (this: Builder): Component {
    return add(this, {
        predicate: (context, document) => document.isJob()
    });
}