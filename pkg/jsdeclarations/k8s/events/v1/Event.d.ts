// Auto generated code; DO NOT EDIT.

import { ObjectMeta } from "../../apimachinery/meta/v1"
import { EventSource, ObjectReference } from "../../core/v1"
import { EventSeries } from "./EventSeries"

/**
 * Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system. Events have a limited retention time and triggers and messages may evolve with time.  Event consumers should not rely on the timing of an event with a given Reason reflecting a consistent underlying trigger, or the continued existence of events with that Reason.  Events should be treated as informative, best-effort, supplemental data.
 * 
 */
export declare class Event {
    constructor();
    constructor(spec: Omit<Event, "apiVersion" | "kind">);

	/**
     * Action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field cannot be empty for new Events and it can have at most 128 characters.
     * 
     */
    action?: string

	/**
     * APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
     * 
     */
    apiVersion?: string

	/**
     * DeprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
     * 
     */
    deprecatedCount?: number

	/**
     * DeprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
     * 
     */
    deprecatedFirstTimestamp?: string

	/**
     * DeprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
     * 
     */
    deprecatedLastTimestamp?: string

	/**
     * DeprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
     * 
     */
    deprecatedSource?: EventSource

	/**
     * EventTime is the time when this Event was first observed. It is required.
     * 
     */
    eventTime: string

	/**
     * Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
     * 
     */
    kind?: string

	/**
     * Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
     * 
     */
    metadata?: ObjectMeta

	/**
     * Note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
     * 
     */
    note?: string

	/**
     * Reason is why the action was taken. It is human-readable. This field cannot be empty for new Events and it can have at most 128 characters.
     * 
     */
    reason?: string

	/**
     * Regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
     * 
     */
    regarding?: ObjectReference

	/**
     * Related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
     * 
     */
    related?: ObjectReference

	/**
     * ReportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
     * 
     */
    reportingController?: string

	/**
     * ReportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
     * 
     */
    reportingInstance?: string

	/**
     * Series is data about the Event series this event represents or nil if it's a singleton Event.
     * 
     */
    series?: EventSeries

	/**
     * Type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable. This field cannot be empty for new Events.
     * 
     */
    type?: string
}