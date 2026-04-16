import { Document } from "./document";

/** Parses the given text as an object. */
export declare function parse(yaml: string): Object;

/** Parses the given text as a {@link Document}. */
export declare function parseDocument(yaml: string): Document;

/** Parses the given text as an array of {@link Document}s. */
export declare function parseDocuments(yaml: string): Document[];