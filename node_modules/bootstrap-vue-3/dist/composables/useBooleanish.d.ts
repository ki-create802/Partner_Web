import type { Booleanish } from '../types';
import type { Ref } from 'vue';
/**
 * Resolves a Booleanish type reactively to type boolean
 */
declare function useBooleanish(el: Ref<Booleanish>): Readonly<Ref<boolean>>;
declare function useBooleanish(el: Ref<Booleanish | null>): Readonly<Ref<boolean | null>>;
declare function useBooleanish(el: Ref<Booleanish | undefined>): Readonly<Ref<boolean | undefined>>;
declare function useBooleanish(el: Ref<Booleanish | undefined | null>): Readonly<Ref<boolean | undefined | null>>;
export default useBooleanish;
