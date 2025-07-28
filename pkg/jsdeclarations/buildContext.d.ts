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

    /**
     * Adds the given document to a {@link DocumentGroup} named "default".
     * 
     * Checks for an existing {@link DocumentGroup} with name "default" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(document: Document): void;

    /**
     * Adds the given document to a {@link DocumentGroup} with the given name during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(documentGroupName: string, document: Document): void;

    /**
     * Adds a new document to a {@link DocumentGroup} named "default" by parsing given YAML string as a {@link Document}.
     * 
     * Checks for an existing {@link DocumentGroup} with name "default" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(path: string, yamlContent: string): void;

    /**
     * Adds a new document to a {@link DocumentGroup} with the given name during the {@link steps.generateResources}
     * step by parsing given YAML string as a {@link Document}.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(documentGroupName: string, path: string, yamlContent: string): void;

    /**
     * Adds a new document to a {@link DocumentGroup} named "default" by converting given object to a {@link Document}.
     * 
     * Checks for an existing {@link DocumentGroup} with name "default" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(path: string, root: Mapping | object): void;

    /**
     * Adds a new document to a {@link DocumentGroup} with the given name during the {@link steps.generateResources}
     * step by converting given object to a {@link Document}.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(documentGroupName: string, path: string, root: Mapping | object): void;

    /** Adds given group to the document groups list. */
    addDocumentGroup(documentGroup: DocumentGroup): void;

    /**
     * Adds given additional file to a {@link DocumentGroup} named "default".
     * 
     * Checks for an existing {@link DocumentGroup} with name "default" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addAdditionalFile(additionalFile: AdditionalFile): void;

    /**
     * Adds the given additional file to a {@link DocumentGroup} with the given name during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addAdditionalFile(documentGroupName: string, additionalFile: AdditionalFile): void;

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
