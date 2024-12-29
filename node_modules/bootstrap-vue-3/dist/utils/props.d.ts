/**
 * Use data last parameters to allow for currying
 *
 * @param suffix
 * @param value
 * @returns
 */
export declare const suffixPropName: (suffix: string, value: string) => string;
/**
 * Given an array of properties or an object of property keys, plucks all the values off the target object, returning a new object that has props that reference the original prop values.
 * An object equivelent to TS Pick<>
 *
 * @param {ReadonlyArray<PropertyKey> | Record<PropertyKey, unknown>} keysToPluck
 * @param {Record<PropertyKey, unknown>} objToPluck
 * @returns
 */
export declare const pluckProps: <A extends Record<PropertyKey, unknown>, B extends Record<PropertyKey, unknown> | readonly PropertyKey[]>(objToPluck: A, keysToPluck: B) => B extends readonly PropertyKey[] ? Pick<A, B[number]> : Pick<A, keyof B>;
