import { Component } from "./component";
import { steps } from "./step";

declare module "./builder" {
    export interface Builder {
        /**
         * Adds a {@link Component} that collects Custom Resource Definitions (CRDs) from all
         * the document groups and moves them into a new document group after the {@link steps.modify} step.
         * @param options Options for collecting CRDs.
         */
        collectCRDs(options?: collectCRDs.Options): Component;
    }
}

export declare namespace collectCRDs {
    export const componentType: string;

    export class Options {
        constructor(documentGroupPath?: string);

        /** Path of the document group in which CRDs will be collected. Default value is 'crds'. */
        documentGroupPath?: string;
    }
}