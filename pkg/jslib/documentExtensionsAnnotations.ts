import * as anemos from "@ohayocorp/anemos"

declare module "@ohayocorp/anemos" {
    export interface Document {
        /** Sets an annotation on the document. */
        setAnnotation(key: string, value: string): void;

        /** Sets the given key value pairs as annotations on the document. */
        setAnnotations(annotations: { [key: string]: string }): void;
    }
}

anemos.Document.prototype.setAnnotation = function (this: anemos.Document, key: string, value: string): void {
    this.metadata ??= {};
    this.metadata.annotations ??= {};
    this.metadata.annotations[key] = value;
}

anemos.Document.prototype.setAnnotations = function (this: anemos.Document, annotations: { [key: string]: string }): void {
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