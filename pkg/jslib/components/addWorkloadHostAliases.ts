import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import { HostAlias } from "@ohayocorp/anemos/k8s/core/v1";
import * as steps from "@ohayocorp/anemos/steps";
import { Workload } from "@ohayocorp/anemos/documentExtensions";

export type Predicate = (document: Document, context: BuildContext) => boolean;

export const componentType = "add-workload-host-aliases";

export class Options {
    /** List of host aliases to be added to the pod spec of workload resources. */
    entries: HostAlias[];
    
    /** Predicate to filter which documents to modify. All workload documents will be modified if not specified. */
    predicate?: Predicate;

    constructor(entries: HostAlias[], predicate?: Predicate) {
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

            if (this.options.predicate && !this.options.predicate(document, context)) {
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
        addWorkloadHostAliases(entries: HostAlias[], predicate?: Predicate): Component;
        
        /**
         * Adds given host aliases to the pod spec of all workload resources.
         */
        addWorkloadHostAliases(ip: string, hostname: string | string[], predicate?: Predicate): Component;
    }
}

Builder.prototype.addWorkloadHostAliases = function (
    this: Builder,
    first: Options | HostAlias[] | string,
    second?: string | string[] | Predicate,
    third?: Predicate
): Component {
    if (typeof first === "object") {
        return add(this, first as Options);
    }
    
    if (Array.isArray(first)) {
        if (second && typeof second !== "function") {
            throw new Error("Invalid argument expected function for predicate");
        }

        return add(this, new Options(first, second as Predicate));
    }
    
    if (typeof first === "string") {
        if (!second || (typeof second !== "string" && !Array.isArray(second))) {
            throw new Error("Hostname must be specified when IP address is given");
        }

        if (third && typeof third !== "function") {
            throw new Error("Invalid argument expected function for predicate");
        }

        const entries = [
            {
                ip: first,
                hostnames: Array.isArray(second) ? second : [second]
            }
        ];

        return add(this, new Options(entries, third));
    }

    throw new Error("Invalid argument");
}