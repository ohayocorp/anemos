import { Component } from "./component";
import { Document } from "./document";
import { DocumentGroup } from "./documentGroup";
import { steps } from "./step";

declare module "./builder" {
    export interface Builder {
        /**
         * Applies the generated manifests to the Kubernetes cluster at the {@link steps.apply} step.
         * @param options Options for applying manifests.
         */
        apply(options?: apply.Options): Component;
    }
}

export declare namespace apply {
    export class Options {
        constructor();

        /** Regex patterns of document group paths to include. If not set, all groups will be included. */
        documentGroups?: string[];

        /** Skip confirmation prompt and apply changes directly. */
        skipConfirmation?: boolean;

        /** Forcefully apply changes even if there are conflicts on server side apply. */
        forceConflicts?: boolean;
    }
}