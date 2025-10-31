import { Document } from "@ohayocorp/anemos/document";
import { Container, PodSpec } from "@ohayocorp/anemos/k8s/core/v1";
import { ObjectMeta } from "@ohayocorp/anemos/k8s/apimachinery/meta/v1";

// Keep the as*/is* keys that are explicitly allowed and all non-prefixed keys by omitting the remaining as*/is* keys.
type AsOrIsKeys<T> = Extract<keyof T, `as${string}` | `is${string}`>;
type ExcludeAsIsExcept<T, Allowed extends AsOrIsKeys<T>> =
    Omit<T, Exclude<AsOrIsKeys<T>, Allowed>>;

// Whitelist of workload-related as*/is* methods to retain.
type WorkloadRelatedAsIs = Extract<keyof Document,
    | 'asCronJob' | 'isCronJob'
    | 'asDaemonSet' | 'isDaemonSet'
    | 'asDeployment' | 'isDeployment'
    | 'asJob' | 'isJob'
    | 'asPod' | 'isPod'
    | 'asReplicaSet' | 'isReplicaSet'
    | 'asStatefulSet' | 'isStatefulSet'
>;

export declare class Workload implements ExcludeAsIsExcept<Document, WorkloadRelatedAsIs> {
    getWorkloadSpec(): PodSpec | undefined;
    ensureWorkloadSpec(): PodSpec;

    getWorkloadMetadata(): ObjectMeta | undefined;
    ensureWorkloadMetadata(): ObjectMeta;
    
    getWorkloadLabels(): Record<string, string> | undefined;
    ensureWorkloadLabels(): Record<string, string>;

    getWorkloadAnnotations(): Record<string, string> | undefined;
    ensureWorkloadAnnotations(): Record<string, string>;

    getContainers(): Array<Container> | undefined;
    getContainer(indexOrName: number | string): Container | undefined;

    getInitContainers(): Array<Container> | undefined;
    getInitContainer(indexOrName: number | string): Container | undefined;
}

Document.prototype.getWorkloadSpec = function (this: Document): PodSpec | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    if (this.isPod()) {
        return this.spec;
    }

    return this.spec?.template?.spec;
}

Document.prototype.ensureWorkloadSpec = function (this: Document): PodSpec {
    if (!this.asWorkload()) {
        throw new Error("Document is not a workload");
    }

    if (this.isPod()) {
        return this.spec ??= {};
    }

    const spec = this.spec ??= {};
    const template = spec.template ??= {};
    
    return template.spec ??= {};
}

Document.prototype.getWorkloadMetadata = function (this: Document): ObjectMeta | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    if (this.isPod()) {
        return this.metadata;
    }

    return this.spec?.template?.metadata;
}

Document.prototype.ensureWorkloadMetadata = function (this: Document): ObjectMeta {
    if (!this.asWorkload()) {
        throw new Error("Document is not a workload");
    }

    if (this.isPod()) {
        return this.metadata ??= {};
    }

    const spec = this.spec ??= {};
    const template = spec.template ??= {};
    
    return template.metadata ??= {};
}

Document.prototype.getWorkloadLabels = function (this: Document): Record<string, string> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    return this.getWorkloadMetadata()?.labels;
}

Document.prototype.ensureWorkloadLabels = function (this: Document): Record<string, string> {
    if (!this.asWorkload()) {
        throw new Error("Document is not a workload");
    }

    return this.ensureWorkloadMetadata().labels ??= {};
}

Document.prototype.getWorkloadAnnotations = function (this: Document): Record<string, string> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    return this.getWorkloadMetadata()?.labels;
}

Document.prototype.ensureWorkloadAnnotations = function (this: Document): Record<string, string> {
    if (!this.asWorkload()) {
        throw new Error("Document is not a workload");
    }

    return this.ensureWorkloadMetadata().annotations ??= {};
}

Document.prototype.getContainers = function (this: Document): Array<Container> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    return this.getWorkloadSpec()?.containers;
}

Document.prototype.getInitContainers = function (this: Document): Array<Container> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }
    
    return this.getWorkloadSpec()?.initContainers;
}

Document.prototype.getContainer = function (this: Document, indexOrName: number | string): Container | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    const containers = this.getContainers();
    if (!containers) {
        return undefined;
    }

    if (typeof indexOrName === "number") {
        return containers[indexOrName];
    }

    return containers.find(container => container.name === indexOrName);
};

Document.prototype.getInitContainer = function (this: Document, indexOrName: number | string): Container | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    const containers = this.getInitContainers();
    if (!containers) {
        return undefined;
    }

    if (typeof indexOrName === "number") {
        return containers[indexOrName];
    }

    return containers.find(container => container.name === indexOrName);
};
