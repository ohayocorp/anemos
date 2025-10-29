// Auto generated code; DO NOT EDIT.
import { CustomResourceColumnDefinition } from "./CustomResourceColumnDefinition"
import { CustomResourceSubresources } from "./CustomResourceSubresources"
import { CustomResourceValidation } from "./CustomResourceValidation"
import { SelectableField } from "./SelectableField"

/**
 * CustomResourceDefinitionVersion describes a version for CRD.
 */
export declare class CustomResourceDefinitionVersion {
    constructor();
    constructor(spec: Pick<CustomResourceDefinitionVersion, "additionalPrinterColumns" | "deprecated" | "deprecationWarning" | "name" | "schema" | "selectableFields" | "served" | "storage" | "subresources">);

	/**
     * AdditionalPrinterColumns specifies additional columns returned in Table output. See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. If no columns are specified, a single column displaying the age of the custom resource is used.
     */
    additionalPrinterColumns?: Array<CustomResourceColumnDefinition>

	/**
     * Deprecated indicates this version of the custom resource API is deprecated. When set to true, API requests to this version receive a warning header in the server response. Defaults to false.
     */
    deprecated?: boolean

	/**
     * DeprecationWarning overrides the default warning returned to API clients. May only be set when `deprecated` is true. The default warning indicates this version is deprecated and recommends use of the newest served version of equal or greater stability, if one exists.
     */
    deprecationWarning?: string

	/**
     * Name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
     */
    name: string

	/**
     * Schema describes the schema used for validation, pruning, and defaulting of this version of the custom resource.
     */
    schema?: CustomResourceValidation

	/**
     * SelectableFields specifies paths to fields that may be used as field selectors. A maximum of 8 selectable fields are allowed. See https://kubernetes.io/docs/concepts/overview/working-with-objects/field-selectors
     */
    selectableFields?: Array<SelectableField>

	/**
     * Served is a flag enabling/disabling this version from being served via REST APIs
     */
    served: boolean

	/**
     * Storage indicates this version should be used when persisting custom resources to storage. There must be exactly one version with storage=true.
     */
    storage: boolean

	/**
     * Subresources specify what subresources this version of the defined custom resource have.
     */
    subresources?: CustomResourceSubresources

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}