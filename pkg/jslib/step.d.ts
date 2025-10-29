/**
 * A class that can be used to specify order of things in a flexible manner. This object consists of a
 * list of numbers. When two steps are compared, each number in their lists are compared according to their indexes.
 * If one of the lists do not have a number at an index, it is assumed to have 0.
 * 
 * For example, if (1, 2) is compared to (1), first the numbers at index 0 are compared. Since they are equal then the numbers
 * at index 1 are compared. The second step does not have a number at that index, so it is assumed that it has 0 and the first step
 * is determined to be greater than the second one.
 * 
 * This makes it possible to create a step between any two steps. For example, given two steps (1, 1), (1, 2),
 * the step (1, 1, 1) is between them.
 */
export declare class Step {
    constructor(description: string, numbers: number[]);

    /** A description of the step, useful for debugging and logging. */
    description: string;

    /** The list of numbers that define the step. */
    numbers: number[];

    /** Returns true if the two steps are equal. */
    equals(other: Step): boolean;

    /** Compares this step to another step. */
    compareTo(other: Step): number;
}