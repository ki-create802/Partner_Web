/**
 * Validates a number is an integer. If not return NaN
 *
 * @param value
 * @param defaultValue
 * @returns
 */
export declare const toInteger: (value: number, defaultValue?: number) => number;
/**
 * @param value
 * @param defaultValue
 * @returns
 */
export declare const stringToInteger: (value: string, defaultValue?: number) => number;
/**
 * Validates a number is a float. If not return NaN
 *
 * @param value
 * @param defaultValue
 * @returns
 */
export declare const toFloat: (value: number | string, defaultValue?: number) => number;
/**
 * Float given to a certain decimal place
 *
 * @param val
 * @param precision
 * @returns
 */
export declare const toPercison: (val: number, precision: number) => string;
