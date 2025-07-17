/**
 * Represents a YAML scalar value with a specific style.
 * The scalar can be a string, number, or boolean and can have different styles such as
 * plain, single-quoted, double-quoted, literal, or folded.
 */
export declare enum YamlStyle {
    /** No quotes or special formatting. */
    Plain,

    /** Content is wrapped in single quotes. */
    SingleQuoted,

    /** Content is wrapped in double quotes. */
    DoubleQuoted,

    /**
     * Content is represented as a literal block (preserving newlines).
     * E.g.:
     * ```
     * literal: |
     *   This is a literal block.
     *   It preserves newlines.
     * ```
     */
    Literal,

    /**
     * Content is represented as a folded block (folding newlines).
     * E.g.:
     * ```
     * folded: >
     *   This is a folded block.
     *   It folds newlines.
     * ```
     */
    Folded,
}

/**
 * Represents a YAML scalar value. A scalar has a string value and an optional style.
 */
export declare class Scalar {
    /**
     * Creates a scalar with the given value and style.
     * If no value is provided, it defaults to an empty string.
     * If no style is provided, it defaults to plain style.
     * @param value The scalar value (string, number, or boolean).
     * @param style The style of the scalar (YamlStyle).
     */
    constructor(value?: string | number | boolean, style?: YamlStyle);

    /** Returns the value of the scalar. */
    getValue(): string;

    /** Sets the value of the scalar. */
    setValue(value: string | number | boolean): void;

    /** Sets the style of the scalar. */
    setStyle(style: YamlStyle): void;
    
    /** Returns a deep clone of the scalar. */
    clone(): Scalar;
}