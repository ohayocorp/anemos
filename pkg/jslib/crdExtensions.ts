import * as anemos from "@ohayocorp/anemos";
import { CustomResourceDefinition } from "./native/k8s/apiextensions/v1";

declare module "@ohayocorp/anemos" {
    export interface Document {
        asCRD(): this is anemos.Document & CustomResourceDefinition;
    }
}

anemos.Document.prototype.asCRD = function (this: anemos.Document): this is anemos.Document & CustomResourceDefinition {
    return this.isCRD();
}
