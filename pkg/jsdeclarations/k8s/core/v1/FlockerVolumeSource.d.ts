// Auto generated code; DO NOT EDIT.



/**
 * Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.
 * 
 */
export declare class FlockerVolumeSource {
    constructor();
    constructor(spec: FlockerVolumeSource);

	/**
     * DatasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated
     * 
     */
    datasetName?: string

	/**
     * DatasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset
     * 
     */
    datasetUUID?: string
}