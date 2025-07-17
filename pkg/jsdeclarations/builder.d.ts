import { Component } from "./component";
import { BuildContext } from "./buildContext";
import { BuilderOptions, EnvironmentType, KubernetesDistribution, Version } from "./builderOptions";
import { Document } from "./document";
import { Mapping } from "./mapping";
import { AdditionalFile } from "./documentGroup";
import { Step, steps } from "./step";

export declare class Builder {
    constructor(options: BuilderOptions);
    constructor(version: string | Version, distribution: KubernetesDistribution, environmentType: EnvironmentType);

    /**
     * All components that will be run by this builder. Don't modify this array directly, use {@link Builder.addComponent}
     * or {@link Builder.removeComponent} instead.
     */
    components: ReadonlyArray<Component>;

    /** Common options that are used by the builder components. */
    options: BuilderOptions;

    /** Adds given component to the list of components. */
    addComponent(component: Component): void;

    /** Removes given component from the list of components. */
    removeComponent(component: Component): void;

    /** Creates a new component with the given action and adds it to the list of components. */
    onStep(step: Step, callback: (context: BuildContext) => void): void;

    /** Adds the given document to a {@link DocumentGroup} named "" during the {@link steps.generateResources} step. */
    addDocument(document: Document): void;

    /**
     * Adds a new document to a {@link DocumentGroup} named "" during the {@link steps.generateResources}
     * step by parsing given YAML string as a {@link Document}.
     */
    addDocument(path: string, yamlContent: string): void;

    /**
     * Adds a new document to a {@link DocumentGroup} named "" during the {@link steps.generateResources}
     * step by converting given object to a {@link Document}.
     */
    addDocument(path: string, root: Mapping | object): void;

    /** Adds the given additional file to a {@link DocumentGroup} named "" during the {@link steps.generateResources} step. */
    addAdditionalFile(additionalFile: AdditionalFile): void;

    /** Runs all the components that were added to the builder. */
    build(): void;
}