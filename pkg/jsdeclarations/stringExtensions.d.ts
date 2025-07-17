/** Indents the data with the given number of spaces. Note that the first line is not indented. */
export declare function indent(text: string, numberOfSpaces: number): string;

/** Dedents the string by removing the common leading whitespace from each line. */
export declare function dedent(text: string): string;

/**
 * Converts a string to a valid Kubernetes identifier by replacing invalid characters with a dash.
 * Truncates the string to 63 characters if it exceeds that length.
 */
export declare function toKubernetesIdentifier(text: string): string;