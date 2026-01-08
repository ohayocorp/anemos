import { Component as AnemosComponent } from "@ohayocorp/anemos/component";
import { Builder } from "@ohayocorp/anemos/builder";
import { BuildContext } from "@ohayocorp/anemos/buildContext";
import * as steps from "@ohayocorp/anemos/steps";
import { Document } from "@ohayocorp/anemos/document";
import { Report } from "@ohayocorp/anemos/report";
import { Container, ResourceRequirements } from "@ohayocorp/anemos/k8s/core/v1";
import { Quantity } from "@ohayocorp/anemos/quantity";

export const componentType = "report-resource-requirements";

class ContainerEntry {
    name: string;
    isInitContainer: boolean;
    isSidecarContainer: boolean;
    resources: ResourceRequirements;

    constructor(name: string, isInitContainer: boolean, isSidecarContainer: boolean, resources: ResourceRequirements) {
        this.name = name;
        this.isInitContainer = isInitContainer;
        this.isSidecarContainer = isSidecarContainer;
        this.resources = resources;
    }
}

class Entry {
    document: Document;
    minReplicas: number;
    maxReplicas: number;
    isJob: boolean;
    containers: ContainerEntry[];

    constructor(document: Document, minReplicas: number, maxReplicas: number, isJob: boolean) {
        this.document = document;
        this.minReplicas = minReplicas;
        this.maxReplicas = maxReplicas;
        this.isJob = isJob;
        this.containers = [];
    }
}

export class Component extends AnemosComponent {
    constructor() {
        super();

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(steps.report, this.report);
    }

    private report = (context: BuildContext) => {
        const hpaDocuments = this.collectHPAs(context);
        const entries: Entry[] = [];

        for (const document of context.getAllDocuments()) {
            if (!document.asWorkload()) {
                continue;
            }

            const { min, max } = this.getReplicaCounts(document, hpaDocuments);

            const entry: Entry = {
                document: document,
                minReplicas: min,
                maxReplicas: max,
                isJob: document.isJob() || document.isCronJob(),
                containers: this.getContainerEntries(document),
            };

            entries.push(entry);
        }

        if (entries.length === 0) {
            return;
        }

        const report = new Report({ filePath: "resource-requirements.md" }, this.getReportContent(entries));
        context.addReport(report);
    }

    private getContainerEntries = (document: Document): ContainerEntry[] => {
        const containers = document.getContainers() ?? [];
        const initContainers = document.getInitContainers() ?? [];

        const containerEntries: ContainerEntry[] = [];

        for (const container of containers) {
            const containerEntry = this.getContainerEntry(container, false);
            if (containerEntry) {
                containerEntries.push(containerEntry);
            }
        }

        for (const container of initContainers) {
            const containerEntry = this.getContainerEntry(container, true);
            if (containerEntry) {
                containerEntries.push(containerEntry);
            }
        }

        return containerEntries;
    }

    private getContainerEntry = (container: Container, isInitContainer: boolean): ContainerEntry | undefined => {
        const resources = container.resources;
        if (!resources) {
            return undefined;
        }

        return {
            name: container.name,
            isInitContainer: isInitContainer,
            isSidecarContainer: isInitContainer && container.restartPolicy === "Always",
            resources: resources,
        };
    }

    private collectHPAs = (context: BuildContext): Map<string, Document> => {
        const hpaDocuments = new Map<string, Document>();

        for (const document of context.getAllDocuments()) {
            if (!document.asHorizontalPodAutoscaler()) {
                continue;
            }

            const namespace = document.metadata?.namespace;
            if (!namespace) {
                continue;
            }

            const scaleTargetRef = document.spec?.scaleTargetRef;
            if (!scaleTargetRef) {
                continue;
            }

            const apiVersion = scaleTargetRef.apiVersion;
            const kind = scaleTargetRef.kind;
            const name = scaleTargetRef.name;

            if (!apiVersion || !kind || !name) {
                continue;
            }

            hpaDocuments.set(`${apiVersion}/${kind}/${namespace}/${name}`, document);
        }

        return hpaDocuments;
    }

    private getReplicaCounts = (document: Document, hpas: Map<string, Document>): { min: number, max: number } => {
        if (document.asPod()) {
            return { min: 1, max: 1 };
        }

        if (document.asJob()) {
            const parallelism = document.spec?.parallelism;
            if (!parallelism) {
                return { min: 1, max: 1 };
            }

            return { min: 1, max: parallelism };
        }

        const hpa = hpas.get(`${document.apiVersion}/${document.kind}/${document.metadata?.namespace}/${document.metadata?.name}`);
        if (hpa) {
            const minReplicas = hpa.spec?.minReplicas ?? 1;
            const maxReplicas = hpa.spec?.maxReplicas ?? minReplicas;

            return { min: minReplicas, max: maxReplicas };
        }

        const replicas = document.spec?.replicas;
        if (replicas) {
            return { min: replicas, max: replicas };
        }

        return { min: 1, max: 1 };
    }

    private getReportContent = (entries: Entry[]): string => {
        var content = `
            # Resource Requirements

            The following tables show the summary of the resource requirements for all workloads in the manifests. The tables include
            the values for minimum and maximum replica counts. This information can be used to determine the hardware requirements.

            The table also includes the highest resource requests. This information can be used to determine the node size requirements
            so that the container with the highest resource requirements can be scheduled on a node.

            Init containers run before the main containers are started. Therefore, only the init containers or the main containers are running
            at any given time. The only exception is the init containers that are sidecar containers, in which case they run alongside the main
            containers. See [Kubernetes documentation](https://kubernetes.io/docs/concepts/workloads/pods/sidecar-containers) for more information
            on sidecar containers.
            
            The resource requirement calculation takes into account that the init containers and the main containers are not running at the same time.
            Therefore, the results are not simply the sum of the init containers and the main containers, but the maximum value between the two for
            each workload.

            Lastly, jobs are mostly short-lived workloads that run to completion and not all the jobs run at the same time. When calculating the
            hardware requirements, it is possible to add only a fraction of the resources required by the jobs to the total resource requirements.
            `.dedent();

        content += "\n\n";

        content += this.getReportTable(entries, "Without Jobs", false, true);
        content += this.getReportTable(entries, "With Jobs", true, true);
        content += this.getReportTable(entries, "Jobs Only", true, false);

        return content;
    }

    private getReportTable = (entries: Entry[], description: string, includeJobs: boolean, includeNonJobs: boolean): string => {
        var content = `
            ### ${description}
            
            | Resource Type | Min Requests | Max Requests | Min Limits | Max Limits | Highest Request |
            | ------------- | ------------ | ------------ | ---------- | ---------- | --------------- |
            `.dedent();

        const requestsMinStats = this.getStatistic(entries, includeJobs, includeNonJobs, "requests", this.sumMinReplicas);
        const requestsMaxStats = this.getStatistic(entries, includeJobs, includeNonJobs, "requests", this.sumMaxReplicas);
        const limitsMinStats = this.getStatistic(entries, includeJobs, includeNonJobs, "limits", this.sumMinReplicas);
        const limitsMaxStats = this.getStatistic(entries, includeJobs, includeNonJobs, "limits", this.sumMaxReplicas);
        const highestRequestStats = this.getStatistic(entries, includeJobs, includeNonJobs, "requests", this.max);

        const resourceTypes: string[] = Array.from(new Set<string>([
            ...Object.keys(requestsMinStats),
            ...Object.keys(limitsMinStats)
        ]));

        resourceTypes.sort();

        for (const resourceType of resourceTypes) {
            content += `| ${resourceType} | ${requestsMinStats[resourceType]} | ${requestsMaxStats[resourceType]} | ${limitsMinStats[resourceType]} | ${limitsMaxStats[resourceType]} | ${highestRequestStats[resourceType]} |\n`;
        }

        content += "\n";

        return content;
    }

    private getStatistic = (entries: Entry[], includeJobs: boolean, includeNonJobs: boolean, resourceType: "requests" | "limits", statisticFunction: (aggregate: Quantity, current: Quantity, entry: Entry) => Quantity): Record<string, Quantity> => {
        const containerResources = new Map<string, Quantity>();
        const initContainerResources = new Map<string, Quantity>();
        const sidecarContainerResources = new Map<string, Quantity>();

        for (const entry of entries) {
            if (entry.isJob && !includeJobs) {
                continue;
            }

            if (!entry.isJob && !includeNonJobs) {
                continue;
            }

            // Aggregate the resources for each container.
            for (const container of entry.containers) {
                const resources = container.resources[resourceType];
                if (!resources) {
                    continue;
                }

                // Select the correct map based on the container type.
                const resourcesMap = container.isInitContainer
                    ? initContainerResources
                    : container.isSidecarContainer
                        ? sidecarContainerResources
                        : containerResources;

                for (const [resource, value] of Object.entries(resources)) {
                    const quantity = new Quantity(value.toString());

                    var aggregate = resourcesMap.get(resource) ?? new Quantity("0");
                    aggregate = statisticFunction(aggregate, quantity, entry);

                    resourcesMap.set(resource, aggregate);
                }
            }
        }

        const allUniqueKeys = new Set([
            ...containerResources.keys(),
            ...initContainerResources.keys(),
            ...sidecarContainerResources.keys()
        ]);

        const result: Record<string, Quantity> = {};

        for (const key of allUniqueKeys) {
            const containerResource = containerResources.get(key) ?? new Quantity("0");
            const initContainerResource = initContainerResources.get(key) ?? new Quantity("0");
            const sidecarContainerResource = sidecarContainerResources.get(key) ?? new Quantity("0");

            var value = result[key] ?? new Quantity("0");

            // Sidecar containers run for the entire lifecycle of the pod.
            value = value.add(sidecarContainerResource ?? new Quantity("0"));

            // Select the maximum value between the init containers and the main containers
            // since they are not running at the same time.
            if (containerResource.compare(initContainerResource) > 0) {
                value = value.add(containerResource);
            } else {
                value = value.add(initContainerResource);
            }

            result[key] = value;
        }

        return result;
    }

    private sumMinReplicas = (aggregate: Quantity, current: Quantity, entry: Entry): Quantity => {
        return aggregate.add(current.multiply(entry.minReplicas));
    }

    private sumMaxReplicas = (aggregate: Quantity, current: Quantity, entry: Entry): Quantity => {
        return aggregate.add(current.multiply(entry.maxReplicas));
    }

    private max = (aggregate: Quantity, current: Quantity): Quantity => {
        if (aggregate.compare(current) < 0) {
            return current;
        }

        return aggregate;
    }
}

export function add(builder: Builder): Component {
    const component = new Component();
    builder.addComponent(component);

    return component;
}