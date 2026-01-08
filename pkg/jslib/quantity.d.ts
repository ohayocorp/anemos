/**
 * Represents a quantity of a resource, such as CPU or memory.
 * It supports basic arithmetic operations like addition, subtraction, and multiplication.
 */
export declare class Quantity {
    /**
     * Creates a new Quantity instance with the given value.
     * The value should be a string representing the quantity, e.g. "100m" for 100 milli-units.
     * @param value The quantity value as a string.
     */
    constructor(value: string);

    /**
     * Adds another quantity to this quantity. Does not modify the original quantity,
     * returns a new Quantity instance instead.
     * @param other The quantity to add.
     * @returns A new Quantity instance representing the sum of this quantity and the other quantity.
     */
    add(other: Quantity): Quantity;

    /**
     * Subtracts another quantity from this quantity. Does not modify the original quantity,
     * returns a new Quantity instance instead.
     * @param other The quantity to subtract.
     * @returns A new Quantity instance representing the difference between this quantity and the other quantity.
     */
    subtract(other: Quantity): Quantity;

    /**
     * Multiplies this quantity by a factor. Does not modify the original quantity,
     * returns a new Quantity instance instead.
     * @param factor The factor to multiply by.
     * @returns A new Quantity instance representing the product of this quantity and the factor.
     */
    multiply(factor: number): Quantity;

    /**
     * Checks if this quantity is equal to another quantity.
     * @param other The quantity to compare to.
     * @returns True if the quantities are equal, false otherwise.
     */
    equals(other: Quantity): boolean;

    /**
     * Compares this quantity to another quantity.
     * @param other The quantity to compare to.
     * @returns A negative number if this quantity is less than the other quantity,
     *          a positive number if this quantity is greater than the other quantity,
     *          and 0 if they are equal.
     */
    compare(other: Quantity): number;

    /**
     * Returns a deep clone of this quantity.
     * @returns A new Quantity instance with the same value as this quantity.
     */
    clone(): Quantity;

    /**
     * Returns a string representation of this quantity.
     * @returns The string representation of the quantity.
     */
    toString(): string;
}