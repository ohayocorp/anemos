import { Document } from "./document";

export type Severity = string;
export type Category = string;

export const info: Severity;
export const warning: Severity;
export const error: Severity;

export const linting: Category;
export const security: Category;
export const specs: Category;

export declare class DiagnosticMetadata {
    id: string;
    name: string;
    description: string;
    severity: Severity;
    categories: Category[];

    constructor(id: string, name: string, description: string, severity: Severity, categories: Category[]);
}

export declare class Diagnostic {
    metadata: DiagnosticMetadata;
    message: string;
    document?: Document;

    constructor(metadata: DiagnosticMetadata, message: string, document?: Document);
}