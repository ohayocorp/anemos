import { Document } from "@ohayocorp/anemos/document";

declare module "@ohayocorp/anemos/document" {
    export interface Document {
        /** Creates a deep copy of the document. */
        clone(): Document;
    }
}

Document.prototype.clone = function (this: Document): Document {
    // Serialize the document to JSON and then parse it back to create a deep copy.
    const jsonString = JSON.stringify(this);
    const clonedObject = JSON.parse(jsonString);
    
    return new Document(clonedObject);
};