import { type Ref } from 'vue';
import type { TableField, TableFieldObject, TableItem } from '../../types';
declare const _default: () => {
    normaliseFields: (origFields: TableField[], items: TableItem[]) => TableFieldObject[];
    mapItems: (fields: TableField[], items: TableItem[], props: any, flags: Record<string, Ref<boolean>>) => TableItem[];
    internalItems: Ref<{
        [x: string]: unknown;
        _rowVariant?: keyof import("../../types").BaseColorVariant | undefined;
        _cellVariants?: Partial<Record<string, keyof import("../../types").BaseColorVariant>> | undefined;
        _showDetails?: boolean | undefined;
    }[]>;
    updateInternalItems: (items: TableItem<Record<string, any>>[]) => Promise<TableItem[] | undefined>;
    filterEvent: Ref<((items: TableItem[]) => void) | undefined>;
    notifyFilteredItems: () => void;
    formatItem: (item: TableItem, fields: TableFieldObject) => unknown;
};
export default _default;
