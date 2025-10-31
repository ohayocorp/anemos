import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import { Document } from "@ohayocorp/anemos/document";
import * as steps from "@ohayocorp/anemos/steps";
import { Container } from "@ohayocorp/anemos/k8s/core/v1";

export type Predicate = (document: Document, container: Container, context: BuildContext) => boolean;

export const componentType = "override-environment-variables";

export class Options {
    /** Environment variables to override in workloads. */
    environmentVariables: Record<string, string>;
    
    /**
     * Predicate to filter which documents and containers to modify. All containers of all workload documents
     * will be modified if not specified.
     */
    predicate?: Predicate;

    constructor(
        environmentVariables: Record<string, string>,
        predicate?: Predicate
    ) {
        this.environmentVariables = environmentVariables;
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
        if (Object.keys(this.options.environmentVariables).length === 0) {
            return;
        }

        for (const document of context.getAllDocuments()) {
            if (!document.asWorkload()) {
                continue;
            }

            const containers = [
                ...(document.getContainers() ?? []),
                ...(document.getInitContainers() ?? [])
            ];

            for (const container of containers) {
                if (this.options.predicate && !this.options.predicate(document, container, context)) {
                    continue;
                }

                this.overrideEnvironmentVariables(container);
            }
        }
    }

    overrideEnvironmentVariables = (container: Container) => {
        const envs = container.env ?? [];

        for (const env of envs ) {
            for (const key of Object.keys(this.options.environmentVariables)) {
                if (env.name === key) {
                    // Delete all other properties of the env var and set only the value.
                    Object.keys(env).forEach(prop => {
                        if (prop !== "name" && prop !== "value") {
                            delete env[prop];
                        }
                    });

                    env.value = this.options.environmentVariables[key];
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
         * Overrides the environment variables in the container specs of all workload resources. Sets the given
         * environment variable values only if the environment variable already exists in the container spec.
         * It is possible to filter which documents and containers to modify by specifying a predicate in the options.
         */
        overrideEnvironmentVariables(options: Options): Component;
        
        /**
         * Overrides the environment variables in the container specs of all workload resources. Sets the given
         * environment variable values only if the environment variable already exists in the container spec.
         * It is possible to filter which documents and containers to modify by specifying a predicate in the options.
         */
        overrideEnvironmentVariables(environmentVariables: Record<string, string>, predicate?: Predicate): Component;
        
        /**
         * Overrides the environment variables in the container specs of all workload resources. Sets the given
         * environment variable values only if the environment variable already exists in the container spec.
         * It is possible to filter which documents and containers to modify by specifying a predicate in the options.
         */
        overrideEnvironmentVariables(name: string, value: string, predicate?: Predicate): Component;
    }
}

Builder.prototype.overrideEnvironmentVariables = function (
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