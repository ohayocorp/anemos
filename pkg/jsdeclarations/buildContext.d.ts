import { Document } from "./document";
import { Mapping } from "./mapping";
import { Component } from "./component";
import { BuilderOptions } from "./builderOptions";
import { KubernetesResourceInfo } from "./kubernetesResourceInfo";
import { DocumentGroup, AdditionalFile } from "./documentGroup";

export declare class AddDocumentOptions {
    path: string;
    content: string | Mapping | object;
    documentGroup?: string;
}

export declare class BuildContext {
    private constructor();

    /** Common options that are used by the builder components. */
    builderOptions: BuilderOptions;

    /** Contains information about the API resources defined in the target cluster. */
    kubernetesResourceInfo: KubernetesResourceInfo;

    /** Custom data that can be used by the builder components. */
    customData: Record<string, any>;

    /**
     * Adds the given document to a {@link DocumentGroup} named "".
     * 
     * Checks for an existing {@link DocumentGroup} with name "" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(document: Document): void;

    /**
     * Adds a new document using the provided options during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name as `options.documentGroup` and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(options: AddDocumentOptions): void;

    /** Adds given group to the document groups list. */
    addDocumentGroup(documentGroup: DocumentGroup): void;

    /**
     * Adds given additional file to a {@link DocumentGroup} named "".
     * 
     * Checks for an existing {@link DocumentGroup} with name "" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addAdditionalFile(additionalFile: AdditionalFile): void;

    /**
     * Adds the given additional file to a {@link DocumentGroup} with the given name during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addAdditionalFile(documentGroupPath: string, additionalFile: AdditionalFile): void;

    /** Removes given group from the document groups list. */
    removeDocumentGroup(documentGroup: DocumentGroup): void;

    /** Returns the document group that matches the specified name. */
    getDocumentGroup(name: string): DocumentGroup | null;

    /** Returns all document groups. If component is provided, only returns groups added by that component. */
    getDocumentGroups(component?: Component): DocumentGroup[];

    /** Returns all documents inside all document groups. */
    getAllDocuments(): Document[];

    /** Returns all documents inside all document groups sorted by their file path. */
    getAllDocumentsSorted(): Document[];

    /** Returns the first document that satisfies the given predicate. Returns null if no document is found. */
    getDocument(predicate: (document: Document) => boolean): Document | null;

    /** Returns the first document that has the given path. Returns null if no document is found. */
    getDocument(path: string): Document | null;
}
