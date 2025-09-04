/** Indents the data with the given number of spaces. Note that the first line is not indented. */
export declare function indent(text: string, numberOfSpaces: number): string;

/** Dedents the string by removing the common leading whitespace from each line. */
export declare function dedent(text: string): string;

/**
 * Converts a string to a valid Kubernetes identifier by replacing invalid characters with a dash.
 * Truncates the string to 63 characters if it exceeds that length.
 */
export declare function toKubernetesIdentifier(text: string): string;

/** Base64 encodes the string. */
export declare function base64Encode(text: string): string;

/** Base64 decodes the string. */
export declare function base64Decode(text: string): string;

declare global {
    interface String {
        /** Indents the string with the given number of spaces. Note that the first line is not indented. */
        indent(numberOfSpaces: number): string;

        /** Dedents the string by removing the common leading whitespace from each line. */
        dedent(): string;

        /**
         * Converts the string to a valid Kubernetes identifier by replacing invalid characters with a dash.
         * Truncates the string to 63 characters if it exceeds that length.
         */
        toKubernetesIdentifier(): string;

        /** Base64 encodes the string. */
        base64Encode(): string;

        /** Base64 decodes the string. */
        base64Decode(): string;
    }
}