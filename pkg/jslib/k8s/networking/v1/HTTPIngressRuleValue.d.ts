// Auto generated code; DO NOT EDIT.
import { HTTPIngressPath } from "./HTTPIngressPath"

/**
 * HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
 */
export declare class HTTPIngressRuleValue {
    constructor();
    constructor(spec: Pick<HTTPIngressRuleValue, "paths">);

	/**
     * Paths is a collection of paths that map requests to backends.
     */
    paths: Array<HTTPIngressPath>

	/**
     * This declaration allows setting and getting custom properties on the document without TypeScript
     * compiler errors.
     */
    [customProperties: string]: any;
}