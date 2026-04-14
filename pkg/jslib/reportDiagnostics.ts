import { Component } from "./component";
import * as steps from "./steps";

declare module "./builder" {
    export interface Builder {
        /**
         * Adds a {@link Component} that reports diagnostics.
         * This component is used to generate diagnostic reports during the build process.
         * It is executed during the {@link steps.report} step.
         * @param options Options for reporting diagnostics.
         */
        reportDiagnostics(options?: reportDiagnostics.Options): Component;
    }
}

export declare namespace reportDiagnostics {
    export const componentType: string;
    
    export class Options {
    }
}