import { Document } from "@ohayocorp/anemos/document";

export enum LabelNode {
    /** .metadata.labels */
    metadataLabels,

    /** .metadata.labels for Pods, .spec.template.metadata.labels for other workloads. */
    workloadLabels,

    /** .spec.selector.matchLabels for workloads other than Pods. */
    workloadSelector,

    /** .spec.selector for Services. */
    serviceSelector,

    /** .spec.selector for monitoring.coreos.com/v1 ServiceMonitors. */
    serviceMonitorSelector
}

declare module "@ohayocorp/anemos/document" {
    export interface Document {
        /**
         * Sets a label on the document.
         */
        setLabel(key: string, value: string): void;

        /** Sets the given key value pairs to the specified nodes on the document. */
        setLabels(labels: { [key: string]: string }, nodes?: LabelNode[]): void;
    }
}

Document.prototype.setLabel = function (this: Document, key: string, value: string): void {
    this.metadata ??= {};
    this.metadata.labels ??= {};
    this.metadata.labels[key] = value;
}

Document.prototype.setLabels = function (this: Document, labels: { [key: string]: string }, nodes?: LabelNode[]): void {
    nodes ??= [LabelNode.metadataLabels];

    const setters: LabelSetter[] = nodes.map(node => {
        switch (node) {
            case LabelNode.metadataLabels:
                return MetadataLabels.instance;
            case LabelNode.workloadLabels:
                return WorkloadLabels.instance;
            case LabelNode.workloadSelector:
                return WorkloadSelector.instance;
            case LabelNode.serviceSelector:
                return ServiceSelector.instance;
            case LabelNode.serviceMonitorSelector:
                return ServiceMonitorSelector.instance;
        }
    });

    for (const setter of setters) {
        setter.setLabels(this, labels);
    }
};

abstract class LabelSetter {
    setLabels(document: Document, labels: { [key: string]: string }): void {
        const node = this.getNode(document);
        if (!node) {
            return;
        }

        Object.entries(labels).forEach(([key, value]) => {
            node[key] = value;
        });
    }

    ensureObject(document: Document, parts: string[]): void {
        let current: any = document;
        for (const part of parts) {
            current[part] ??= {};
            current = current[part];
        }
    }

    abstract getNode(document: Document): any | undefined;
}

class MetadataLabels extends LabelSetter {
    static readonly instance = new MetadataLabels();

    getNode(document: Document): any | undefined {
        super.ensureObject(document, ["metadata", "labels"]);
        return document.metadata?.labels;
    }
}

class WorkloadLabels extends LabelSetter {
    static readonly instance = new WorkloadLabels();

    getNode(document: Document): any | undefined {
        if (!document.isWorkload()) {
            return undefined;
        }

        if (document.isPod()) {
            return MetadataLabels.instance.getNode(document);
        }

        super.ensureObject(document, ["spec", "template", "metadata", "labels"]);
        return document.spec.template.metadata.labels;
    }
}

class WorkloadSelector extends LabelSetter {
    static readonly instance = new WorkloadSelector();

    getNode(document: Document): any | undefined {
        if (!document.isWorkload()) {
            return undefined;
        }

        // Specifying pod selector for Jobs is not necessary most of the time. If a custom
        // selector is needed, users can set it manually.
        // https://kubernetes.io/docs/concepts/workloads/controllers/job/#specifying-your-own-pod-selector
        if (document.isJob()) {
            return undefined;
        }

        super.ensureObject(document, ["spec", "selector", "matchLabels"]);
        return document.spec.selector.matchLabels;
    }
}

class ServiceSelector extends LabelSetter {
    static readonly instance = new ServiceSelector();

    getNode(document: Document): any | undefined {
        if (!document.isService()) {
            return undefined;
        }

        super.ensureObject(document, ["spec", "selector"]);
        return document.spec.selector;
    }
}

class ServiceMonitorSelector extends LabelSetter {
    static readonly instance = new ServiceMonitorSelector();

    getNode(document: Document): any | undefined {
        if (document.kind !== "ServiceMonitor") {
            return undefined;
        }

        super.ensureObject(document, ["spec", "selector"]);
        return document.spec.selector;
    }
}