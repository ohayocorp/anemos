import * as anemos from "@ohayocorp/anemos";
import { Container, PodSpec } from "./native/k8s/core/v1";

export interface Workload {
    getWorkloadSpec(): PodSpec | undefined;

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
