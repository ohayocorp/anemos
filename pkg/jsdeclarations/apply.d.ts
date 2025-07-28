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
        constructor(documents: DocumentGroup);
        constructor(documents: DocumentGroup, namespace: string);
        constructor(documents: Document[], name: string);
        constructor(documents: Document[], name: string, namespace: string);

        documents?: Document[];
        applySetParentName?: string;
        applySetParentNamespace?: string;
        skipConfirmation?: boolean;
    }
}