import { DocumentGroup } from "./documentGroup"

/**
 * Document corresponds to a single YAML document. Note that even though a YAML file can contain multiple documents,
 * each one of these documents is represented by a separate Document object.
 */
export declare class Document {
    /** Creates a new empty document. */
    constructor();

    /** Creates a new document with the given content. Content must be a valid YAML string or a JavaScript object. */
    constructor(content: string | Object);

    /**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;

    /** The document group this document belongs to. */
    group?: DocumentGroup;

    /** Returns the file path to the document. May contain multiple segments separated by slashes. */
    getPath(): string | null;

    /** Sets the file path to the document. May contain multiple segments separated by slashes. */
    setPath(path: string | null): void;

    /** Returns the file path to write the document into. Adds group name as base directory if it is not null. */
    fullPath(): string;

    /** Apply and wait for this document after the given document. Documents must be in the same group. */
    provisionAfter(other: Document): void;

    /** Apply and wait for this document before the given document. Documents must be in the same group. */
    provisionBefore(other: Document): void;
}

export declare class NewDocumentOptions {
    content: string | Object;
    path?: string;
    documentGroup?: string;
}
