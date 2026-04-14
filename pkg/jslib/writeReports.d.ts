import { Component } from "./component";
import * as steps from "./steps";

declare module "./builder" {
    export interface Builder {
        /**
         * Adds a {@link Component} that writes the reports to the output directory
         * as requested file formats during the {@link steps.output} step.
         * @param options Options for writing reports.
         */
        writeReports(options?: writeReports.Options): Component;
    }
}

export type ReportOutputType = string;

export const markdown: ReportOutputType;
export const html: ReportOutputType;

export declare namespace writeReports {
    export const componentType: string;
    
    export class Options {
        outputTypes?: ReportOutputType[];
    }
}