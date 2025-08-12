import { Component } from "./component";
import { steps } from "./step";

declare module "./builder" {
    export interface Builder {
        /**
         * Adds a {@link Component} that collects namespace definitions from all
         * the document groups and moves them into a new document group after the {@link steps.modify} step.
         * @param options Options for collecting namespaces.
         */
        collectNamespaces(options?: collectNamespaces.Options): Component;
    }
}

export declare namespace collectNamespaces {
    export const componentType: string;
    
    export class Options {
        constructor(documentGroupPath?: string);

        /** Path of the document group in which namespaces will be collected. Default value is 'namespaces'. */
        documentGroupPath?: string;
    }
}