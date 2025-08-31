import { Component } from "./component";
import { steps } from "./step";

declare module "./builder" {
    export interface Builder {
        /**
         * Adds a {@link Component} that deletes the output directory.
         * This component is used to clean up the output directory before generating new content.
         * It is executed before the {@link steps.output} step.
         * @param options Options for deleting the output directory.
         */
        deleteOutputDirectory(options?: deleteOutputDirectory.Options): Component;
    }
}

export declare namespace deleteOutputDirectory {
    export class Options {
    }
}