import * as anemos from "@ohayocorp/anemos";
import { Container, PodSpec } from "./native/k8s/core/v1";
import { ObjectMeta } from "./native/k8s/apimachinery/meta/v1";

// Keep the as*/is* keys that are explicitly allowed and all non-prefixed keys by omitting the remaining as*/is* keys.
type AsOrIsKeys<T> = Extract<keyof T, `as${string}` | `is${string}`>;
type ExcludeAsIsExcept<T, Allowed extends AsOrIsKeys<T>> =
    Omit<T, Exclude<AsOrIsKeys<T>, Allowed>>;

// Whitelist of workload-related as*/is* methods to retain.
type WorkloadRelatedAsIs = Extract<keyof anemos.Document,
    | 'asCronJob' | 'isCronJob'
    | 'asDaemonSet' | 'isDaemonSet'
    | 'asDeployment' | 'isDeployment'
    | 'asJob' | 'isJob'
    | 'asPod' | 'isPod'
    | 'asReplicaSet' | 'isReplicaSet'
    | 'asStatefulSet' | 'isStatefulSet'
>;

export declare class Workload implements ExcludeAsIsExcept<anemos.Document, WorkloadRelatedAsIs> {
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

anemos.Document.prototype.getWorkloadSpec = function (this: anemos.Document): PodSpec | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    if (this.isPod()) {
        return this.spec;
    }

    return this.spec?.template?.spec;
}

anemos.Document.prototype.ensureWorkloadSpec = function (this: anemos.Document): PodSpec {
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

anemos.Document.prototype.getWorkloadMetadata = function (this: anemos.Document): ObjectMeta | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    if (this.isPod()) {
        return this.metadata;
    }

    return this.spec?.template?.metadata;
}

anemos.Document.prototype.ensureWorkloadMetadata = function (this: anemos.Document): ObjectMeta {
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

anemos.Document.prototype.getWorkloadLabels = function (this: anemos.Document): Record<string, string> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    return this.getWorkloadMetadata()?.labels;
}

anemos.Document.prototype.ensureWorkloadLabels = function (this: anemos.Document): Record<string, string> {
    if (!this.asWorkload()) {
        throw new Error("Document is not a workload");
    }

    return this.ensureWorkloadMetadata().labels ??= {};
}

anemos.Document.prototype.getWorkloadAnnotations = function (this: anemos.Document): Record<string, string> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    return this.getWorkloadMetadata()?.labels;
}

anemos.Document.prototype.ensureWorkloadAnnotations = function (this: anemos.Document): Record<string, string> {
    if (!this.asWorkload()) {
        throw new Error("Document is not a workload");
    }

    return this.ensureWorkloadMetadata().annotations ??= {};
}

anemos.Document.prototype.getContainers = function (this: anemos.Document): Array<Container> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }

    return this.getWorkloadSpec()?.containers;
}

anemos.Document.prototype.getInitContainers = function (this: anemos.Document): Array<Container> | undefined {
    if (!this.asWorkload()) {
        return undefined;
    }
    
    return this.getWorkloadSpec()?.initContainers;
}

anemos.Document.prototype.getContainer = function (this: anemos.Document, indexOrName: number | string): Container | undefined {
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

anemos.Document.prototype.getInitContainer = function (this: anemos.Document, indexOrName: number | string): Container | undefined {
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
