// Auto generated code; DO NOT EDIT.



/**
 * EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.
 * 
 */
export declare class EventSeries {
    constructor();
    constructor(spec: EventSeries);

	/**
     * Number of occurrences in this series up to the last heartbeat time
     * 
     */
    count?: number

	/**
     * Time of the last occurrence observed
     * 
     */
    lastObservedTime?: string
}