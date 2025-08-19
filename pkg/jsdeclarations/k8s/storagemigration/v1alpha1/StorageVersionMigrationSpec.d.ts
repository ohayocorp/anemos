// Auto generated code; DO NOT EDIT.

import { GroupVersionResource } from "./GroupVersionResource"

/**
 * Spec of the storage version migration.
 * 
 */
export declare class StorageVersionMigrationSpec {
    constructor();
    constructor(spec: StorageVersionMigrationSpec);

	/**
     * The token used in the list options to get the next chunk of objects to migrate. When the .status.conditions indicates the migration is "Running", users can use this token to check the progress of the migration.
     * 
     */
    continueToken?: string

	/**
     * The resource that is being migrated. The migrator sends requests to the endpoint serving the resource. Immutable.
     * 
     */
    resource: GroupVersionResource
}