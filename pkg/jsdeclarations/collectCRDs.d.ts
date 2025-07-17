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
    export class Options {
        constructor(directory?: string);

        /** Directory where CRDs will be written. Default value is 'crds'. */
        directory?: string;
    }
}