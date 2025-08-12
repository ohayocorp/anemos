import { Component } from "./component";

declare module "./builder" {
    export interface Builder {
        /**
         * Create namespace manifests for all the namespaces that are referenced by other resources.
         * @param options The options for creating referenced namespaces.
         */
        createReferencedNamespaces(options?: createReferencedNamespaces.Options): Component;
    }
}

export declare namespace createReferencedNamespaces {
    export const componentType: string;

    export class Options {
        /**
         * A predicate function to filter namespaces. Returns true if the namespace should be included.
         * If not specified, all namespaces except the default Kubernetes namespaces will be included.
         */
        predicate: (namespace: string) => boolean;
    }
}