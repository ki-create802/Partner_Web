import type { Booleanish } from '../types';
/**
 * Type predicate function to check if the input is Booleanish type
 *
 * @param  {unknown} input
 * @returns inputisBooleanish
 */
export declare const isBooleanish: (input: unknown) => input is Booleanish;
/**
 * Converts a Booleanish type to boolean
 *
 * @param {Booleanish} input
 * @returns
 */
export declare const resolveBooleanish: (input: Booleanish) => boolean;
