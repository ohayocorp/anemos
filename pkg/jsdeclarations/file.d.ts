export declare namespace file {
    /**
     * Reads the file from the file system and returns the content as a string.
     * File path must be under the main script directory.
     */
    export function readAllText(filePath: string): string;

    /**
     * Reads the file from the file system and returns the content as a byte array.
     * File path must be under the main script directory.
     */
    export function readAllBytes(filePath: string): Uint8Array;

    /**
     * Writes the content to the file system.
     * File path must be under the main script directory.
     */
    export function writeAllText(filePath: string, content: string): void;

    /**
     * Writes the content to the file system.
     * File path must be under the main script directory.
     */
    export function writeAllBytes(filePath: string, bytes: Uint8Array): void;

    /** Gets the main script path. */
    export function mainScriptPath(): string;

    /** Gets the main script directory. */
    export function mainScriptDirectory(): string;

    /** Gets the current script path. */
    export function currentScriptPath(): string;

    /** Gets the current script directory. */
    export function currentScriptDirectory(): string;
}