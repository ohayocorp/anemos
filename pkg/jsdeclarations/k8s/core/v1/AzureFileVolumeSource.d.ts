// Auto generated code; DO NOT EDIT.



/**
 * AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
 * 
 */
export declare class AzureFileVolumeSource {
    constructor();
    constructor(spec: AzureFileVolumeSource);

	/**
     * ReadOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     * 
     */
    readOnly?: boolean

	/**
     * SecretName is the  name of secret that contains Azure Storage Account Name and Key
     * 
     */
    secretName: string

	/**
     * ShareName is the azure share Name
     * 
     */
    shareName: string
}