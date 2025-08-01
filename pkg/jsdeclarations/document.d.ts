import { Mapping } from "./mapping"
import { DocumentGroup } from "./documentGroup"

/**
 * Document corresponds to a single YAML document. Note that even though a YAML file can contain multiple documents,
 * each one of these documents is represented by a separate Document object.
 * 
 * Although the root of the document can be any kind of node, only {@link Mapping} is supported.
 */
export declare class Document {
    /** Creates a new empty document with the given path. */
    constructor(path: string);

    /** Creates a new document with the given path and YAML content. The YAML content must be a valid YAML mapping. */
    constructor(path: string, yamlContent: string);
    
    /** Creates a new document with the given path and root. The root must be a {@link Mapping} or an {@link Object}. */
    constructor(path: string, root: Mapping | Object);

    /** The file path to the document. May contain multiple segments separated by slashes. */
    path: string;

    /** The document group this document belongs to. */
    group?: DocumentGroup;

    /** Returns the file path to write the document into. Adds group name as base directory if it is not null. */
    fullPath(): string;

    /** Return the root of the document as a {@link Mapping}. */
    getRoot(): Mapping;

    /** Returns a deep clone of the document. */
    clone(): Document;

    /** Ensures a {@link Mapping} for ".metadata". */
    ensureMetadata(): Mapping;
}
