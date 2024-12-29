/**
 * Removes properties from an object, based on the values in an array, and returns the new object.
 * Equivalent to an object version of TS Omit<>
 *
 * @param {Record<PropertyKey, unknown>} objToPluck
 * @param {ReadonlyArray<PropertyKey>} keysToPluck
 * @returns
 */
export declare const omit: <A extends Record<PropertyKey, unknown>, B extends readonly PropertyKey[]>(objToPluck: A, keysToPluck: B) => Omit<A, B[number]>;
/**
 * Convenience method to create a read-only descriptor
 *
 * @returns
 */
export declare const readonlyDescriptor: () => {
    enumerable: boolean;
    configurable: boolean;
    writable: boolean;
};
/**
 * Deeply clones an item
 *
 * @param {T} source
 * @returns
 */
export declare const cloneDeep: <T>(source: T) => T;
export declare const cloneDeepAsync: <T>(source: T) => Promise<T>;
