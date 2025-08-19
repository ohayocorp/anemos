// Auto generated code; DO NOT EDIT.

import { MigrationCondition } from "./MigrationCondition"

/**
 * Status of the storage version migration.
 * 
 */
export declare class StorageVersionMigrationStatus {
    constructor();
    constructor(spec: StorageVersionMigrationStatus);

	/**
     * The latest available observations of the migration's current state.
     * 
     */
    conditions?: Array<MigrationCondition>

	/**
     * ResourceVersion to compare with the GC cache for performing the migration. This is the current resource version of given group, version and resource when kube-controller-manager first observes this StorageVersionMigration resource.
     * 
     */
    resourceVersion?: string
}