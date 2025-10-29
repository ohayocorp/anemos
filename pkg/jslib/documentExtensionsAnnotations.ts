import { Document } from "@ohayocorp/anemos/document";

declare module "@ohayocorp/anemos/document" {
    export interface Document {
        /** Sets an annotation on the document. */
        setAnnotation(key: string, value: string): void;

        /** Sets the given key value pairs as annotations on the document. */
        setAnnotations(annotations: { [key: string]: string }): void;
    }
}

Document.prototype.setAnnotation = function (this: Document, key: string, value: string): void {
    this.metadata ??= {};
    this.metadata.annotations ??= {};
    this.metadata.annotations[key] = value;
}

Document.prototype.setAnnotations = function (this: Document, annotations: { [key: string]: string }): void {
    const metadata = this.metadata ??= {};

    // Some Helm charts may have annotations node with empty string value instead of an object or null.
    if (typeof metadata.annotations !== "object") {
        metadata.annotations = {};
    }

    metadata.annotations ??= {};

    Object.entries(annotations).forEach(([key, value]) => {
        metadata.annotations![key] = value;
    });
};