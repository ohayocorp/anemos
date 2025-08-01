import { Component } from "./component";
import { BuildContext } from "./buildContext";
import { Document } from "./document";

/**
 * Container for multiple {@link Document} instances. It can also contain multiple
 * {@link AdditionalFile} objects that generates additional files beside the documents.
 */
export declare class DocumentGroup {
    constructor(path: string);

    /** The path of the document group. */
    path: string;

    /**
     * Documents in this group. Don't modify this array directly. Use {@link DocumentGroup.addDocument}
     * and {@link DocumentGroup.removeDocument} instead.
     */
    documents: Document[];

    /** Additional files in this group. Don't modify this array directly. Use {@link DocumentGroup.addAdditionalFile} instead. */
    additionalFiles: AdditionalFile[];

    /** Adds the given document to this group and sets its group field to this group. */
    addDocument(document: Document): void;

    /** Adds the given additional file to this group. */
    addAdditionalFile(additionalFile: AdditionalFile): void;

    /**
     * Returns the component that created this document group. Component is set when
     * the document group is added to the builder context.
     */
    getComponent(): Component | null;

    /** Returns the first document that has the given path. */
    getDocument(path: string): Document | null;

    /** Returns the first document that satisfies the given predicate. */
    getDocument(filter: (document: Document) => boolean): Document | null;

    /** Returns the documents in this group sorted by their path. */
    sortedDocuments(): Document[];

    /** Removes all documents and additional files from this group and adds them to the given group. */
    moveTo(group: DocumentGroup): void;

    /** Removes the given document from this group and sets its group field to null. */
    removeDocument(document: Document): void;

    /** Removes the given additional file from this group. */
    removeAdditionalFile(additionalFile: AdditionalFile): void;

    /** 
     * Sets the namespace of documents that are namespaced (e.g. Pod, Job, ...) to the given namespace.
     * This is useful for documents that are generated by Helm charts where the namespace is not set
     * correctly.
     */
    setNamespaces(context: BuildContext, namespace: string): void;

    /** 
     * Fixes duplicate document file paths by adding an index suffix to the end the paths. Sorts the documents
     * by their apiVersion, kind, namespace, and name to ensure consistent naming.
     * This is useful for documents that are generated by Helm charts which may have multiple documents in
     * a single file.
     */
    fixNameClashes(): void;
}

/**
 * Represents a file that is not a document but should be included in the output.
 * It has a path and content, and is typically used for configuration files or other auxiliary files
 * that need to be written alongside the documents.
 */
export declare class AdditionalFile {
    constructor(path: string, content: string);

    /** The file path to write the additional file into. */
    path: string;

    /** The content of the additional file. */
    content: string;
}