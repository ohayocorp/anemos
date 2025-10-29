import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import { HostAlias } from "@ohayocorp/anemos/k8s/core/v1";
import * as steps from "@ohayocorp/anemos/steps";
import { Workload } from "@ohayocorp/anemos/documentExtensionsWorkload";

export const componentType = "add-workload-host-aliases";

export class Options {
    /** List of host aliases to be added to the pod spec of workload resources. */
    entries: HostAlias[];
    
    /** Predicate to filter which documents to modify. All workload documents will be modified if not specified. */
    predicate?: (context: BuildContext, document: Document) => boolean;

    constructor(entries: HostAlias[], predicate?: (context: BuildContext, document: Document) => boolean) {
        this.entries = entries;
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

            this.addHostAliases(document);
        }
    }

    addHostAliases = (workload: Workload) => {
        if (this.options.entries.length === 0) {
            return;
        }

        const spec = workload.getWorkloadSpec();
        if (!spec) {
            return;
        }

        const hostAliases = spec.hostAliases ??= [];

        for (const entry of this.options.entries) {
            const hostAlias = hostAliases.find(x => x.ip === entry.ip);

            // If the host alias for the IP address does not exist, add the given entry directly.
            if (!hostAlias) {
                hostAliases.push(entry);
                continue;
            }

            // If the host alias for the IP address exists, add the hostnames to it.
            const hostnames = hostAlias.hostnames ??= [];

            for (const hostname of entry.hostnames ?? []) {
                if (!hostnames.includes(hostname)) {
                    hostnames.push(hostname);
                }
            }
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
         * Adds given host aliases to the pod spec of all workload resources.
         */
        addWorkloadHostAliases(options: Options): Component;
        
        /**
         * Adds given host aliases to the pod spec of all workload resources.
         */
        addWorkloadHostAliases(entries: HostAlias[]): Component;
    }
}

Builder.prototype.addWorkloadHostAliases = function (this: Builder, arg: Options | HostAlias[]): Component {
    if (Array.isArray(arg)) {
        return add(this, new Options(arg));
    } else if (typeof arg === "object") {
        return add(this, arg as Options);
    }

    throw new Error("Invalid argument");
}