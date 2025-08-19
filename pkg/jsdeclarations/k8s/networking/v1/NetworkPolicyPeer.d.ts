// Auto generated code; DO NOT EDIT.

import { LabelSelector } from "../../apimachinery/meta/v1"
import { IPBlock } from "./IPBlock"

/**
 * NetworkPolicyPeer describes a peer to allow traffic to/from. Only certain combinations of fields are allowed
 * 
 */
export declare class NetworkPolicyPeer {
    constructor();
    constructor(spec: NetworkPolicyPeer);

	/**
     * IpBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.
     * 
     */
    ipBlock?: IPBlock

	/**
     * NamespaceSelector selects namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.
     * 
     * If podSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the namespaces selected by namespaceSelector. Otherwise it selects all pods in the namespaces selected by namespaceSelector.
     * 
     */
    namespaceSelector?: LabelSelector

	/**
     * PodSelector is a label selector which selects pods. This field follows standard label selector semantics; if present but empty, it selects all pods.
     * 
     * If namespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the pods matching podSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the pods matching podSelector in the policy's own namespace.
     * 
     */
    podSelector?: LabelSelector
}