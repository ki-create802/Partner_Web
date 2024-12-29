/**
 * Convert a value to a string that can be rendered `undefined`/`null` will be converted to `''` Plain objects and arrays will be JSON stringified
 *
 * @param val
 * @param spaces
 * @returns
 */
export declare const toString: (val: unknown, spaces?: number) => string;
/**
 * @param str
 * @returns
 */
export declare const startCase: (str: string) => string;
/**
 * @param str
 * @returns
 */
export declare const titleCase: (str: string) => string;
/**
 * Uppercases the first letter of a string and returns a new string
 *
 * @param str
 * @returns
 */
export declare const upperFirst: (str: string) => string;
