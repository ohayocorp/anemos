import * as anemos from "@ohayocorp/anemos"

export const componentType = "sort-fields";

export class Component extends anemos.Component {
    constructor() {
        super();

        this.setComponentType(componentType);
        this.setIdentifier(componentType);

        this.addAction(new anemos.Step("Sort Fields", [...anemos.steps.modify.numbers, 1]), this.modify);
    }

    modify = (context: anemos.BuildContext) => {
        for (const document of context.getAllDocuments()) {
            for (const descriptor of sortDescriptors) {
                if (!descriptor.canApply(document)) {
                    continue;
                }

                for (const sorter of descriptor.getSorters()) {
                    sorter.sort(document);
                }
            }
        }
    }
}

export function add(builder: anemos.Builder): Component {
    const component = new Component();
    builder.addComponent(component);

    return component;
}

declare module "@ohayocorp/anemos" {
    export interface Builder {
        /**
         * Sorts the fields of the documents in a way that is consistent with the general Kubernetes
         * conventions.
         */
        sortFields(): Component;
    }
}

anemos.Builder.prototype.sortFields = function (this: anemos.Builder): Component {
    return add(this);
}

abstract class Sorter {
    sort(document: anemos.Document): void {
        const object = this.getObjects(document);
        const objects = Array.isArray(object) ? object : [object];

        for (const object of objects) {
            this.sortInternal(object);
        }
    }

    abstract getObjects(document: anemos.Document): any | any[];
    abstract sortInternal(object: any): void;
}

interface SortDescriptor {
    canApply(document: anemos.Document): boolean;
    getSorters(): Sorter[];
}

class SpecSortDescriptor implements SortDescriptor {
    canApply(_: anemos.Document): boolean {
        return true;
    }

    getSorters(): Sorter[] {
        return [
            new SortFields((document) => document, "spec"),
        ];
    }
}

class WorkloadSpecSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isWorkload() && !document.isPod();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields((document) => document.spec, "replicas", "serviceName", "strategy", "selector", "template"),
            new SortByKey((document) => document.spec?.selector?.matchLabels),
            new SortByKey((document) => document.spec?.template?.metadata?.labels),
            new SortByKey((document) => document.spec?.template?.metadata?.annotations),
        ];
    }
}

class WorkloadTemplateSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isWorkload() && !document.isPod();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields((document) => document.spec?.template, "metadata", "spec"),
        ];
    }
}

class WorkloadTemplateSpecSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isWorkload();
    }

    getSorters(): Sorter[] {
        const getter = (document: anemos.Document) => document.isPod() ? document.spec : document.spec?.template?.spec;

        return [
            new SortFields(
                getter,
                "containers",
                "initContainers",

                "hostNetwork",
                "hostPID",
                "hostIPC",
                "shareProcessNamespace",
                "serviceAccountName",
                "automountServiceAccountToken",
                "securityContext",

                "hostAliases",
                "dnsPolicy",
                "dnsConfig",
                "hostname",
                "subdomain",

                "nodeName",
                "nodeSelector",
                "tolerations",
                "affinity",
                "topologySpreadConstraints",

                "schedulerName",
                "preemptionPolicy",
                "priorityClassName",
                "priority",

                "terminationGracePeriodSeconds",
                "restartPolicy",
                "enableServiceLinks",
                "imagePullSecrets",
            ),
            new SortFieldsTrailing(getter, "volumes")
        ];
    }
}

class ContainerSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isWorkload();
    }

    getSorters(): Sorter[] {
        const getter = (document: anemos.Document) => {
            const spec = document.isPod() ? document.spec : document.spec?.template?.spec;
            const containers = spec?.containers || [];
            const initContainers = spec?.initContainers || [];

            return [...containers, ...initContainers];
        };

        return [
            new SortFields(
                getter,
                "name",
                "image",
                "imagePullPolicy",
                "workingDir",
                "command",
                "args",
                "env",
                "envFrom",
                "ports",
                "lifecycle",
                "securityContext",
                "resources",
                "livenessProbe",
                "readinessProbe",
                "startupProbe",
            ),
            new SortFieldsTrailing(getter, "volumeMounts")
        ];
    }
}

class VolumeSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isWorkload();
    }

    getSorters(): Sorter[] {
        const getter = (document: anemos.Document) => {
            const spec = document.isPod() ? document.spec : document.spec?.template?.spec;
            return spec?.volumes || [];
        };

        return [
            new SortFields(getter, "name"),
        ];
    }
}

class VolumeMountSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isWorkload();
    }

    getSorters(): Sorter[] {
        const getter = (document: anemos.Document) => {
            const spec = document.isPod() ? document.spec : document.spec?.template?.spec;
            const containers = spec?.containers || [];
            const initContainers = spec?.initContainers || [];
            const allContainers = [...containers, ...initContainers];

            return allContainers.map(container => container.volumeMounts || []).flat();
        };

        return [
            new SortFields(getter, "name"),
        ];
    }
}

class ContainerPortSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isWorkload();
    }

    getSorters(): Sorter[] {
        const getter = (document: anemos.Document) => {
            const spec = document.isPod() ? document.spec : document.spec?.template?.spec;
            const containers = spec?.containers || [];
            const initContainers = spec?.initContainers || [];
            const allContainers = [...containers, ...initContainers];

            return allContainers.map(container => container.ports || []).flat();
        };

        return [
            new SortFields(getter, "name", "containerPort", "hostIP", "hostPort", "protocol"),
        ];
    }
}

class ConfigMapSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isConfigMap();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields(document => document, "immutable", "data", "binaryData"),
            new SortByKey(document => document.data),
            new SortByKey(document => document.binaryData),
        ];
    }
}

class SecretSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isSecret();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields(document => document, "type", "immutable", "stringData", "data"),
            new SortByKey(document => document.data),
            new SortByKey(document => document.binaryData),
        ];
    }
}

class ServiceSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isService();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields(document => document, "spec"),
            new SortByKey(document => document.spec?.selector),
            new SortFields(
                document => document.spec,
                "type",
                "ports",
                "selector",
                "clusterIP",
                "externalIPs",
                "externalName",
                "loadBalancerIP",
                "loadBalancerSourceRanges",
                "sessionAffinity",
                "externalTrafficPolicy",
                "healthCheckNodePort",
            ),
        ];
    }
}

class IngressSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isIngress();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields(document => document, "spec"),
            new SortFields(document => document.spec, "ingressClassName", "defaultBackend", "rules", "tls"),
            new SortFields(document => document.spec?.rules, "host", "http"),
            new SortFields(document => document.spec?.tls, "secretName", "hosts"),
        ];
    }
}

class RoleClusterRoleSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isRole() || document.isClusterRole();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields(document => document.rules, "apiGroups", "resources", "resourceNames", "nonResourceURLs", "verbs"),
        ];
    }
}

class RoleBindingClusterRoleBindingSortDescriptor implements SortDescriptor {
    canApply(document: anemos.Document): boolean {
        return document.isRoleBinding() || document.isClusterRoleBinding();
    }

    getSorters(): Sorter[] {
        return [
            new SortFields(document => document, "subjects", "roleRef"),
            new SortFields(document => document.subjects, "apiGroup", "kind", "name", "namespace"),
            new SortFields(document => document.roleRef, "apiGroup", "kind", "name"),
        ];
    }
}

class ObjectMetadataSortDescriptor implements SortDescriptor {
    canApply(_: anemos.Document): boolean {
        return true;
    }

    getSorters(): Sorter[] {
        return [
            new SortFields((document) => document.metadata, "name", "namespace", "labels", "annotations", "ownerReferences", "finalizers"),
            new SortByKey((document) => document.metadata?.labels),
            new SortByKey((document) => document.metadata?.annotations),
        ];
    }
}

class ObjectSortDescriptor implements SortDescriptor {
    canApply(_: anemos.Document): boolean {
        return true;
    }

    getSorters(): Sorter[] {
        return [
            new SortFields((document) => document, "apiVersion", "kind", "metadata"),
        ];
    }
}

class SortFields extends Sorter {
    private getter: (document: anemos.Document) => any | any[];
    private fields: string[];

    constructor(getter: (document: anemos.Document) => any | any[], ...fields: string[]) {
        super();

        this.getter = getter;
        this.fields = fields;
    }

    getObjects(document: anemos.Document): any | any[] {
        return this.getter(document);
    }

    sortInternal(object: any): void {
        reorderObjectPropertiesWithFields(object, this.fields, true);
    }
}

class SortFieldsTrailing extends Sorter {
    private getter: (document: anemos.Document) => any | any[];
    private fields: string[];

    constructor(getter: (document: anemos.Document) => any | any[], ...fields: string[]) {
        super();

        this.getter = getter;
        this.fields = fields;
    }

    getObjects(document: anemos.Document): any | any[] {
        return this.getter(document);
    }

    sortInternal(object: any): void {
        reorderObjectPropertiesWithFields(object, this.fields, false);
    }
}

class SortByKey extends Sorter {
    private getter: (document: anemos.Document) => any | any[];

    constructor(getter: (document: anemos.Document) => any | any[]) {
        super();
        this.getter = getter;
    }

    getObjects(document: anemos.Document): any | any[] {
        return this.getter(document);
    }

    sortInternal(object: any): void {
        reorderObjectPropertiesByKey(object);
    }
}

function reorderObjectPropertiesWithFields(object: any, fields: string[], fieldsFirst: boolean): void {
    // Only operate on plain objects and skip arrays or non-objects.
    if (!object || typeof object !== "object" || Array.isArray(object)) {
        return;
    }

    const keys = Object.keys(object);
    if (keys.length <= 1) {
        return;
    }

    // Map each prioritized field to its index for quick lookup.
    const index = new Map<string, number>();
    fields.forEach((f, i) => index.set(f, i));

    // Split keys into those in fields (to be ordered by index) and the rest (stable order).
    const inFields: string[] = [];
    const rest: string[] = [];

    for (const k of keys) {
        if (index.has(k)) {
            inFields.push(k);
        } else {
            rest.push(k);
        }
    }

    inFields.sort((a, b) => (index.get(a)! - index.get(b)!));

    const ordered = fieldsFirst
        ? [...inFields, ...rest]
        : [...rest, ...inFields];

    reorderObjectProperties(object, new Map(ordered.map((k, i) => [k, i])));
}

function reorderObjectPropertiesByKey(object: any): void {
    // Only operate on plain objects and skip arrays or non-objects.
    if (!object || typeof object !== "object" || Array.isArray(object)) {
        return;
    }

    const keys = Object.keys(object);
    if (keys.length <= 1) {
        return;
    }

    const ordered = [...keys].sort();

    reorderObjectProperties(object, new Map(ordered.map((k, i) => [k, i])));
}

function reorderObjectProperties(object: any, indexes: Map<string, number>): void {
    const keys = Object.keys(object);
    if (keys.length <= 1) {
        return;
    }

    const ordered = [...keys].sort((a, b) => indexes.get(a)! - indexes.get(b)!);

    // If the order is already correct, do nothing.
    let changed = false;
    for (let i = 0; i < keys.length; i++) {
        if (keys[i] !== ordered[i]) {
            changed = true;
            break;
        }
    }
    if (!changed) {
        return;
    }

    // Recreate property insertion order by deleting and re-adding in the desired order.
    const values: Record<string, any> = {};
    for (const k of ordered) {
        values[k] = object[k];
    }
    for (const k of keys) {
        delete object[k];
    }
    for (const k of ordered) {
        object[k] = values[k];
    }
}

const sortDescriptors: SortDescriptor[] = [
    new SpecSortDescriptor(),
    new WorkloadSpecSortDescriptor(),
    new WorkloadTemplateSortDescriptor(),
    new WorkloadTemplateSpecSortDescriptor(),
    new ContainerSortDescriptor(),
    new VolumeSortDescriptor(),
    new VolumeMountSortDescriptor(),
    new ContainerPortSortDescriptor(),
    new ConfigMapSortDescriptor(),
    new SecretSortDescriptor(),
    new ServiceSortDescriptor(),
    new IngressSortDescriptor(),
    new RoleClusterRoleSortDescriptor(),
    new RoleBindingClusterRoleBindingSortDescriptor(),
    new ObjectMetadataSortDescriptor(),
    new ObjectSortDescriptor(),
];
