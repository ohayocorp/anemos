// Auto generated code; DO NOT EDIT.

import { ExpressionWarning } from "./ExpressionWarning"

/**
 * TypeChecking contains results of type checking the expressions in the ValidatingAdmissionPolicy
 * 
 */
export declare class TypeChecking {
    constructor();
    constructor(spec: TypeChecking);

	/**
     * The type checking warnings for each expression.
     * 
     */
    expressionWarnings?: Array<ExpressionWarning>
}