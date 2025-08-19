// Auto generated code; DO NOT EDIT.



/**
 * AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
 * 
 */
export declare class AzureFilePersistentVolumeSource {
    constructor();
    constructor(spec: AzureFilePersistentVolumeSource);

	/**
     * ReadOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.
     * 
     */
    readOnly?: boolean

	/**
     * SecretName is the name of secret that contains Azure Storage Account Name and Key
     * 
     */
    secretName: string

	/**
     * SecretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod
     * 
     */
    secretNamespace?: string

	/**
     * ShareName is the azure Share Name
     * 
     */
    shareName: string
}