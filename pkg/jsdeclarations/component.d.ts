import { BuildContext } from "./buildContext";
import { Step } from "./step";

export declare class Action {
    /**
     * The step during which this action will be executed.
     */
    step: Step;

    /**
     * The callback function to be executed for this action.
     */
    callback: (context: BuildContext) => void;
}

/**
 * {@link Component} consists of a series of actions. The action are sorted by their
 * step and then executed by builder.
 */
export declare class Component {
    /**
     * Actions of this component. Don't modify this array directly, use {@link Component.addAction}.
     */
    actions: ReadonlyArray<Action>;

    /**
     * Adds a new action to the component.
     * @param step The step during which this action will be executed.
     * @param callback The callback function to be executed for this action.
     */
    addAction(step: Step, callback: (context: BuildContext) => void): void;

    /**
     * Gets the custom data associated with the given key.
     * @param key The key of the custom data.
     * @returns The custom data associated with the given key, or null if not found.
     * @description Custom data is a key-value pair that can be used to store additional information
     * about the component. It is not used by the builder, but can be used by other components.
     */
    getCustomData(key: string): any | null;

    /**
     * Sets the custom data associated with the given key.
     * @param key The key of the custom data.
     * @param value The value to set for the custom data.
     * @description Custom data is a key-value pair that can be used to store additional information
     * about the component. It is not used by the builder, but can be used by other components.
     */
    setCustomData(key: string, value: any): void;

    /**
     * Gets the metadata associated with the given key.
     * @param key The key of the metadata.
     * @returns The metadata associated with the given key, or null if not found.
     * @description Metadata is a string key-value pair that can be used to store additional information
     * about the component. It is not used by the builder, but can be used by other components.
     */
    getMetadata(key: string): string | null;

    /**
     * Sets the metadata associated with the given key.
     * @param key The key of the metadata.
     * @param value The value to set for the metadata.
     * @description Metadata is a string key-value pair that can be used to store additional information
     * about the component. It is not used by the builder, but can be used by other components.
     */
    setMetadata(key: string, value: string): void;
    
    /**
     * Gets the identifier of the component.
     * @returns The identifier of the component, or null if not set.
     * @description The identifier is a unique string that can be used to identify the component.
     * It is not used by the builder, but can be used by other components to identify the component.
     */
    getIdentifier(): string | null;

    /**
     * Sets the identifier of the component.
     * @param identifier The identifier to set for the component.
     * @description The identifier is a unique string that can be used to identify the component.
     * It is not used by the builder, but can be used by other components to identify the component.
     */
    setIdentifier(identifier: string): void;
    
    /**
     * Gets the type of the component.
     * @returns The type of the component, or null if not set.
     * @description The component type is a string that can be used to identify the type of component
     * that this component belongs to. It is not used by the builder, but can be used by other components.
     */
    getComponentType(): string | null;

    /**
     * Sets the component type of the component.
     * @param componentType The component type to set for the component.
     * @description The component type is a string that can be used to identify the type of component
     * that this component belongs to. It is not used by the builder, but can be used by other components.
     */
    setComponentType(componentType: string): void;
}