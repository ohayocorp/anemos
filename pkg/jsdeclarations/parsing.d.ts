import { Document } from "./document";
import { Mapping } from "./mapping";
import { Sequence } from "./sequence";

/** Parses the given text as a {@link Document}. */
export declare function parseDocument(path: string, text: string): Document;

/** Parses the given text as a {@link Mapping}. */
export declare function parseMapping(text: string): Mapping;

/** Parses the given text as a {@link Sequence}. */
export declare function parseSequence(text: string): Sequence;

/** Serializes given object to a YAML string. */
export declare function serializeToYaml(data: any): string;