import { Mapping } from "./mapping"
import { Scalar } from "./scalar"

/**
 * Represents a YAML sequence, which is an ordered collection of items.
 * A sequence can contain mappings, other sequences, or scalars.
 * It provides methods to add, set, insert, get, and manipulate its items.
 */
export declare class Sequence {
    constructor();
    constructor(items: any[]);

    /** Appends the given {@link Mapping} to the sequence. */
    add(mapping: Mapping | Object): void;

    /** Appends the given {@link Mapping} array to the sequence. */
    add(mappings: Mapping[] | Object[]): void;

    /** Appends the given {@link Sequence} to the sequence. */
    add(sequence: Sequence | any[]): void;

    /** Appends the given {@link Sequence} array to the sequence. */
    add(sequences: Sequence[] | any[]): void;

    /** Appends the given {@link Scalar} to the sequence. */
    add(scalar: Scalar): void;

    /** Appends the given {@link Scalar} array to the sequence. */
    add(scalars: Scalar[]): void;

    /** Appends the given value to the sequence as a {@link Scalar} and returns the created scalar. */
    add(value: string): Scalar;

    /** Appends the given values to the sequence as {@link Scalar}s. */
    add(values: string[]): void;

    /** Sets the child at given index to the given {@link Mapping}. */
    set(index: number, mapping: Mapping): void;

    /** Sets the child at given index to the given {@link Sequence}. */
    set(index: number, sequence: Sequence): void;

    /** Sets the child at given index to the given {@link Scalar}. */
    set(index: number, scalar: Scalar): void;

    /** Sets the value at given index as a {@link Scalar} and returns the created scalar. */
    set(index: number, value: string): Scalar;

    /** Inserts the given {@link Mapping} at given index. */
    insert(index: number, mapping: Mapping): void;

    /** Inserts the given {@link Sequence} at given index. */
    insert(index: number, sequence: Sequence): void;

    /** Inserts the given {@link Scalar} at given index. */
    insert(index: number, scalar: Scalar): void;

    /** Inserts the value as a {@link Scalar} at given index and returns the created scalar. */
    insert(index: number, value: string): Scalar;

    /** Returns the child at given index as {@link Mapping}. Throws if the child is not a {@link Mapping}. */
    getMapping(index: number): Mapping;

    /** Returns the first child that returns true for the filter function. Returns null if the node is not found. */
    getMapping(filter: (mapping: Mapping) => boolean): Mapping | null;

    /** Returns the child at given index as {@link Sequence}. Throws if the child is not a {@link Sequence}. */
    getSequence(index: number): Sequence;

    /** Returns the first child that returns true for the filter function. Returns null if the node is not found. */
    getSequence(filter: (sequence: Sequence) => boolean): Sequence | null;

    /** Returns the child at given index as {@link Scalar}. Throws if the child is not a {@link Scalar}. */
    getScalar(index: number): Scalar;

    /** Returns the first child that returns true for the filter function. Returns null if the node is not found. */
    getScalar(filter: (scalar: Scalar) => boolean): Scalar | null;

    /** Returns the value at given index. Throws if the child is not a {@link Scalar}. */
    getValue(index: number): string;

    /** Returns the number of elements. */
    length(): number;

    /** Removes the node at the given index. */
    remove(index: number): void;

    /** Returns true if the sequence contains a {@link Scalar} with the given value. */
    contains(value: string): boolean;

    /** Removes all elements. */
    clear(): void;

    /** Returns a deep clone of the sequence. */
    clone(): Sequence;
}