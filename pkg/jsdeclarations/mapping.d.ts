import { Scalar } from "./scalar"
import { Sequence } from "./sequence"

export declare class Mapping {
    /**
     * Creates an empty mapping.
     */
    constructor();

    /** 
     * Creates a mapping with the given key and value pairs. If a value contains a newline character,
     * its style will be set to literal. Otherwise, the style will be set to double quoted.
     */
    constructor(record: Record<string, string>);

    /** 
     * Creates a mapping with the given object. Object keys will be converted to scalars and object
     * values will be converted to scalars, sequences or mappings.
     */
    constructor(object: Object);

    /**
     * Returns the {@link Mapping} corresponding to the given key. Returns null if the key doesn't
     * exist or it doesn't correspond to a {@link Mapping}.
     */
    getMapping(key: string): Mapping | null;

    /**
     * Returns a {@link Mapping} by following each one of the keys. Expects that each intermediate node is a {@link Mapping},
     * and last node is a {@link Mapping}, otherwise returns null. Returns null if a key doesn't exist at any point in the chain.
     */
    getMapping(keys: string[]): Mapping | null;

    /**
     * Returns the {@link Sequence} corresponding to the given key. Returns null if the key doesn't
     * exist or it doesn't correspond to a {@link Sequence}.
     */
    getSequence(key: string): Sequence | null;

    /**
     * Returns a {@link Sequence} by following each one of the keys. Expects that each intermediate node is a {@link Mapping},
     * and last node is a {@link Sequence}, otherwise returns null. Returns null if a key doesn't exist at any point in the chain.
     */
    getSequence(keys: string[]): Sequence | null;

    /**
     * Returns the {@link Scalar} corresponding to the given key. Returns null if the key doesn't
     * exist or it doesn't correspond to a {@link Scalar}.
     */
    getScalar(key: string): Scalar | null;

    /**
     * Returns a {@link Scalar} by following each one of the keys. Expects that each intermediate node is a {@link Mapping},
     * and last node is a {@link Scalar}, otherwise returns null. Returns null if a key doesn't exist at any point in the chain.
     */
    getScalar(keys: string[]): Scalar | null;

    /**
     * Returns the value of the given key as a string. Returns null if the key doesn't exist
     * or it doesn't correspond to a {@link Scalar}.
     */
    getValue(key: string): string | null;

    /**
     * Returns the value as a string by following each one of the keys. Expects that each intermediate node is a {@link Mapping},
     * and last node is a {@link Scalar}, otherwise returns null. Returns null if a key doesn't exist at any point in the chain.
     */
    getValue(keys: string[]): string | null;

    /** Sets the given key to the given {@link Mapping} or {@link Object}. */
    set(key: string, value: Mapping | Object): void;

    /** Sets the given key to the given {@link Scalar}. */
    set(key: string, value: Scalar): void;

    /** Sets the given key to the given {@link Sequence} or array. */
    set(key: string, value: Sequence | any[]): void;

    /** Sets the given key to the given string as a {@link Scalar} and returns the created {@link Scalar}. */
    set(key: string, value: string): Scalar;

    /** Inserts the given key value pair at the given index. */
    insert(index: number, key: string, value: Mapping): void;

    /** Inserts the given key value pair at the given index. */
    insert(index: number, key: string, value: Scalar): void;

    /** Inserts the given key value pair at the given index. */
    insert(index: number, key: string, value: Sequence): void;

    /** Inserts the given key value pair at the given index and returns the created {@link Scalar}. */
    insert(index: number, key: string, value: string): Scalar;

    /** Returns true if the key exists in the keys list. */
    containsKey(key: string): boolean;

    /**
     * Ensures that the given key corresponds to a {@link Mapping} and returns it. Creates an empty {@link Mapping}
     * if the key doesn't exist or it exists but correspond to an empty {@link Scalar}. Throws if the key exists but
     * corresponds to a {@link Sequence} or a non-empty {@link Scalar}.
     */
    ensureMapping(key: string): Mapping;

    /**
     * Ensures that the last key corresponds to a {@link Mapping} and returns it. Expects that each intermediate node
     * is a {@link Mapping}. Creates an empty {@link Mapping} if the key doesn't exist or it exists but correspond
     * to an empty {@link Scalar}. Throws if the last node exists but corresponds to a {@link Sequence} or a non-empty {@link Scalar}.
     */
    ensureMapping(keys: string[]): Mapping;

    /**
     * Ensures that the given key corresponds to a {@link Sequence} and returns it. Creates an empty {@link Sequence}
     * if the key doesn't exist or it exists but correspond to an empty {@link Scalar}. Throws if the key exists but
     * corresponds to a {@link Mapping} or a non-empty {@link Scalar}.
     */
    ensureSequence(key: string): Sequence;

    /**
     * Ensures that the last key corresponds to a {@link Sequence} and returns it. Expects that each intermediate node
     * is a {@link Mapping}. Creates an empty {@link Sequence} if the key doesn't exist or it exists but correspond
     * to an empty {@link Scalar}. Throws if the last node exists but corresponds to a {@link Mapping} or a non-empty {@link Scalar}.
     */
    ensureSequence(keys: string[]): Sequence;

    /** Returns the index of the key, or -1 if it doesn't exist. */
    indexOfKey(key: string): number;

    /** Returns the key at given index as a {@link Scalar}. */
    keyAt(index: number): Scalar;

    /** Returns the keys as an array of {@link Scalar}. */
    keys(): Scalar[];

    /** Returns the number of children. */
    length(): number;

    /** Moves the key and corresponding value to the given index. */
    move(key: string, index: number): void;

    /** Removes the key and corresponding value. */
    remove(key: string): void;

    /** Clears all the contents of the mapping. */
    clear(): void;

    /** Sorts the child nodes by their keys. */
    sortByKey(): void;

    /** Returns a deep clone of the mapping. */
    clone(): Mapping;
}