import { Component } from "./component";
import { steps } from "./step";

declare module "./builder" {
    export interface Builder {
        /**
         * Adds a {@link Component} that writes the documents from all document groups to the output directory
         * as YAML files during the {@link steps.output} step.
         * @param options Options for writing documents.
         */
        writeDocuments(options?: writeDocuments.Options): Component;
    }
}

export declare namespace writeDocuments {
    export class Options {
    }
}