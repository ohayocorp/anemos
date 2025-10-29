// Auto generated code; DO NOT EDIT.
import { JSONSchemaProps } from "./JSONSchemaProps"

/**
 * CustomResourceValidation is a list of validation methods for CustomResources.
 */
export declare class CustomResourceValidation {
    constructor();
    constructor(spec: Pick<CustomResourceValidation, "openAPIV3Schema">);

	/**
     * OpenAPIV3Schema is the OpenAPI v3 schema to use for validation and pruning.
     */
    openAPIV3Schema?: JSONSchemaProps

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}