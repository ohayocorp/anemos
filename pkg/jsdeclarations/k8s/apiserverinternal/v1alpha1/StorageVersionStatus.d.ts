// Auto generated code; DO NOT EDIT.

import { ServerStorageVersion } from "./ServerStorageVersion"
import { StorageVersionCondition } from "./StorageVersionCondition"

/**
 * API server instances report the versions they can decode and the version they encode objects to when persisting objects in the backend.
 * 
 */
export declare class StorageVersionStatus {
    constructor();
    constructor(spec: StorageVersionStatus);

	/**
     * If all API server instances agree on the same encoding storage version, then this field is set to that version. Otherwise this field is left empty. API servers should finish updating its storageVersionStatus entry before serving write operations, so that this field will be in sync with the reality.
     * 
     */
    commonEncodingVersion?: string

	/**
     * The latest available observations of the storageVersion's state.
     * 
     */
    conditions?: Array<StorageVersionCondition>

	/**
     * The reported versions per API server instance.
     * 
     */
    storageVersions?: Array<ServerStorageVersion>
}