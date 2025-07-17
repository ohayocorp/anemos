import { Document } from "./document";
import { Mapping } from "./mapping";
import { Component } from "./component";
import { BuilderOptions } from "./builderOptions";
import { KubernetesResourceInfo } from "./kubernetesResourceInfo";
import { DocumentGroup, AdditionalFile } from "./documentGroup";

export declare class BuildContext {
    private constructor();

    /** Common options that are used by the builder components. */
    builderOptions: BuilderOptions;

    /** Contains information about the API resources defined in the target cluster. */
    kubernetesResourceInfo: KubernetesResourceInfo;

    /** Custom data that can be used by the builder components. */
    customData: Record<string, any>;

    /** Adds the given document to a document group named "". */
    addDocument(document: Document): void;

    /** Adds a new document to a {@link DocumentGroup} named "" by parsing given YAML string as a {@link Document}. */
    addDocument(path: string, yamlContent: string): void;

    /** Adds a new document to a {@link DocumentGroup} named "" by converting given object to a {@link Document}. */
    addDocument(path: string, root: Mapping | object): void;

    /** Adds given group to the document groups list. */
    addDocumentGroup(documentGroup: DocumentGroup): void;

    /** Adds given additional file to a document group named "". */
    addAdditionalFile(additionalFile: AdditionalFile): void;

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
