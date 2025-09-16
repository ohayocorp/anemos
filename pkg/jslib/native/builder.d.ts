import { Component } from "./component";
import { BuildContext } from "./buildContext";
import { BuilderOptions, EnvironmentType, KubernetesDistribution, Version } from "./builderOptions";
import { Document, NewDocumentOptions } from "./document";
import { DocumentGroup, AdditionalFile } from "./documentGroup";
import { Step, steps } from "./step";

export declare class Builder {
    constructor();
    constructor(options: BuilderOptions);
    constructor(version: string | Version, distribution: KubernetesDistribution, environmentType: EnvironmentType);

    /**
     * All components that will be run by this builder. Don't add to this array directly, use {@link Builder.addComponent} instead.
     */
    components: Array<Component>;

    /** Common options that are used by the builder components. */
    options: BuilderOptions;

    /** Runs all the components that were added to the builder. */
    build(): void;

    /** Adds given component to the list of components. */
    addComponent(component: Component): void;

    /** Removes given component from the list of components. */
    removeComponent(component: Component): void;

    /**
     * Adds a component that creates a document group with the given name during {@link steps.generateResources}.
     * Document group doesn't contain any documents, it serves as a placeholder for provision dependencies.
     * @param name
     */
    addProvisionCheckpoint(name: string): Component;

    /**
     * Parses the given YAML string as a {@link Document} and adds it to a {@link DocumentGroup} named ""
     * during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the name "" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(document: string): void;

    /**
     * Adds the given document to a {@link DocumentGroup} named "" during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the name "" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(document: Document): void;

    /**
     * Adds a new document using the provided options during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name as `options.documentGroup` and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addDocument(options: NewDocumentOptions): void;

    /**
     * Adds the given additional file to a {@link DocumentGroup} named "" during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the name "" and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addAdditionalFile(additionalFile: AdditionalFile): void;

    /**
     * Adds the given additional file to a {@link DocumentGroup} with the given name during the {@link steps.generateResources} step.
     * 
     * Checks for an existing {@link DocumentGroup} with the same name and adds the document to it if it exists.
     * Creates a new {@link DocumentGroup} if it doesn't exist.
     */
    addAdditionalFile(documentGroupPath: string, additionalFile: AdditionalFile): void;

    /** Creates a new component with the given action and adds it to the list of components. */
    onStep(step: Step, callback: (context: BuildContext) => void): Component;

    /**
     * Creates a new component with the given action that will be run during {@link steps.populateKubernetesResources}
     * and adds it to the list of components.
     */
    onPopulateKubernetesResources(callback: (context: BuildContext) => void): Component;

    /**
     * Creates a new component with the given action that will be run during {@link steps.sanitize}
     * and adds it to the list of components.
     */
    onSanitize(callback: (context: BuildContext) => void): Component;

    /**
     * Creates a new component with the given action that will be run during {@link steps.generateResources}
     * and adds it to the list of components.
     */
    onGenerateResources(callback: (context: BuildContext) => void): Component;

    /**
     * Creates a new component with the given action that will be run during {@link steps.generateResourcesBasedOnOtherResources}
     * and adds it to the list of components.
     */
    onGenerateResourcesBasedOnOtherResources(callback: (context: BuildContext) => void): Component;

    /**
     * Creates a new component with the given action that will be run during {@link steps.modify}
     * and adds it to the list of components.
     */
    onModify(callback: (context: BuildContext) => void): Component;

    /**
     * Creates a new component with the given action that will be run during {@link steps.specifyProvisionerDependencies}
     * and adds it to the list of components.
     */
    onSpecifyProvisionerDependencies(callback: (context: BuildContext) => void): Component;
}