// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { StorageVersionMigrationSpec } from "./StorageVersionMigrationSpec"
import { StorageVersionMigrationStatus } from "./StorageVersionMigrationStatus"

/**
 * StorageVersionMigration represents a migration of stored data to the latest storage version.
 * 
 */
export declare class StorageVersionMigration {
    constructor();
    constructor(spec: Omit<StorageVersionMigration, "apiVersion" | "kind">);

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Specification of the migration.
     * 
     */
    spec?: StorageVersionMigrationSpec

	/**
     * Status of the migration.
     * 
     */
    status?: StorageVersionMigrationStatus
}