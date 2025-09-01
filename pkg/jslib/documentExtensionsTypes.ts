import * as anemos from "@ohayocorp/anemos";

import { CustomResourceDefinition } from "./native/k8s/apiextensions/v1";
import { DaemonSet, Deployment, ReplicaSet, StatefulSet } from "./native/k8s/apps/v1";
import { HorizontalPodAutoscaler } from "./native/k8s/autoscaling/v2";
import { CronJob, Job } from "./native/k8s/batch/v1";
import { ConfigMap, Namespace, PersistentVolume, PersistentVolumeClaim, Pod, Secret, Service, ServiceAccount } from "./native/k8s/core/v1";
import { Ingress } from "./native/k8s/networking/v1";
import { ClusterRole, ClusterRoleBinding, Role, RoleBinding } from "./native/k8s/rbac/v1";

declare module "@ohayocorp/anemos" {
    export interface Document {
        /** Type guard for ClusterRole documents. */
        asClusterRole(): this is anemos.Document & ClusterRole;

        /** Type guard for ClusterRoleBinding documents. */
        asClusterRoleBinding(): this is anemos.Document & ClusterRoleBinding;

        /** Type guard for ConfigMap documents. */
        asConfigMap(): this is anemos.Document & ConfigMap;

        /** Type guard for CustomResourceDefinition documents. */
        asCRD(): this is anemos.Document & CustomResourceDefinition;

        /** Type guard for CronJob documents. */
        asCronJob(): this is anemos.Document & CronJob;

        /** Type guard for CustomResourceDefinition documents. */
        asCustomResourceDefinition(): this is anemos.Document & CustomResourceDefinition;

        /** Type guard for DaemonSet documents. */
        asDaemonSet(): this is anemos.Document & DaemonSet;

        /** Type guard for Deployment documents. */
        asDeployment(): this is anemos.Document & Deployment;

        /** Type guard for HorizontalPodAutoscaler documents. */
        asHorizontalPodAutoscaler(): this is anemos.Document & HorizontalPodAutoscaler;

        /** Type guard for Ingress documents. */
        asIngress(): this is anemos.Document & Ingress;

        /** Type guard for Job documents. */
        asJob(): this is anemos.Document & Job;

        /** Type guard for Namespace documents. */
        asNamespace(): this is anemos.Document & Namespace;

        /** Type guard for PersistentVolume documents. */
        asPersistentVolume(): this is anemos.Document & PersistentVolume;

        /** Type guard for PersistentVolumeClaim documents. */
        asPersistentVolumeClaim(): this is anemos.Document & PersistentVolumeClaim;

        /** Type guard for Pod documents. */
        asPod(): this is anemos.Document & Pod;

        /** Type guard for ReplicaSet documents. */
        asReplicaSet(): this is anemos.Document & ReplicaSet;

        /** Type guard for Role documents. */
        asRole(): this is anemos.Document & Role;

        /** Type guard for RoleBinding documents. */
        asRoleBinding(): this is anemos.Document & RoleBinding;

        /** Type guard for Secret documents. */
        asSecret(): this is anemos.Document & Secret;

        /** Type guard for Service documents. */
        asService(): this is anemos.Document & Service;

        /** Type guard for ServiceAccount documents. */
        asServiceAccount(): this is anemos.Document & ServiceAccount;

        /** Type guard for StatefulSet documents. */
        asStatefulSet(): this is anemos.Document & StatefulSet;

        /** Type guard for workload documents. */
        asWorkload(): this is anemos.Document & anemos.Workload;

        /** Returns true if the document has the given apiVersion and kind. */
        isOfKind(apiVersion: string, kind: string): boolean;

        /** Returns true if the document is a ClusterRole. */
        isClusterRole(): boolean;

        /** Returns true if the document is a ClusterRoleBinding. */
        isClusterRoleBinding(): boolean;

        /** Returns true if the document is a ConfigMap. */
        isConfigMap(): boolean;

        /** Returns true if the document is a CustomResourceDefinition. */
        isCRD(): boolean;

        /** Returns true if the document is a CronJob. */
        isCronJob(): boolean;

        /** Returns true if the document is a CustomResourceDefinition. */
        isCustomResourceDefinition(): boolean;

        /** Returns true if the document is a DaemonSet. */
        isDaemonSet(): boolean;

        /** Returns true if the document is a Deployment. */
        isDeployment(): boolean;

        /** Returns true if the document is a HorizontalPodAutoscaler. */
        isHorizontalPodAutoscaler(): boolean;

        /** Returns true if the document is an Ingress. */
        isIngress(): boolean;

        /** Returns true if the document is a Job. */
        isJob(): boolean;

        /** Returns true if the document is a Namespace. */
        isNamespace(): boolean;

        /** Returns true if the document is a PersistentVolume. */
        isPersistentVolume(): boolean;

        /** Returns true if the document is a PersistentVolumeClaim. */
        isPersistentVolumeClaim(): boolean;

        /** Returns true if the document is a Pod. */
        isPod(): boolean;

        /** Returns true if the document is a ReplicaSet. */
        isReplicaSet(): boolean;

        /** Returns true if the document is a Role. */
        isRole(): boolean;

        /** Returns true if the document is a RoleBinding. */
        isRoleBinding(): boolean;

        /** Returns true if the document is a Secret. */
        isSecret(): boolean;

        /** Returns true if the document is a Service. */
        isService(): boolean;

        /** Returns true if the document is a ServiceAccount. */
        isServiceAccount(): boolean;

        /** Returns true if the document is a StatefulSet. */
        isStatefulSet(): boolean;

        /** Returns true if the document is one of these: CronJob, DaemonSet, Deployment, Job, Pod, ReplicaSet, StatefulSet. */
        isWorkload(): boolean;
    }
}

anemos.Document.prototype.asClusterRole = function (this: anemos.Document): this is anemos.Document & ClusterRole {
    return this.isClusterRole();
}

anemos.Document.prototype.asClusterRoleBinding = function (this: anemos.Document): this is anemos.Document & ClusterRoleBinding {
    return this.isClusterRoleBinding();
}

anemos.Document.prototype.asConfigMap = function (this: anemos.Document): this is anemos.Document & ConfigMap {
    return this.isConfigMap();
}

anemos.Document.prototype.asCRD = function (this: anemos.Document): this is anemos.Document & CustomResourceDefinition {
    return this.isCRD();
}

anemos.Document.prototype.asCronJob = function (this: anemos.Document): this is anemos.Document & CronJob {
    return this.isCronJob();
}

anemos.Document.prototype.asCustomResourceDefinition = function (this: anemos.Document): this is anemos.Document & CustomResourceDefinition {
    return this.isCRD();
}

anemos.Document.prototype.asDaemonSet = function (this: anemos.Document): this is anemos.Document & DaemonSet {
    return this.isDaemonSet();
}

anemos.Document.prototype.asDeployment = function (this: anemos.Document): this is anemos.Document & Deployment {
    return this.isDeployment();
}

anemos.Document.prototype.asHorizontalPodAutoscaler = function (this: anemos.Document): this is anemos.Document & HorizontalPodAutoscaler {
    return this.isHorizontalPodAutoscaler();
}

anemos.Document.prototype.asIngress = function (this: anemos.Document): this is anemos.Document & Ingress {
    return this.isIngress();
}

anemos.Document.prototype.asJob = function (this: anemos.Document): this is anemos.Document & Job {
    return this.isJob();
}

anemos.Document.prototype.asNamespace = function (this: anemos.Document): this is anemos.Document & Namespace {
    return this.isNamespace();
}

anemos.Document.prototype.asPersistentVolume = function (this: anemos.Document): this is anemos.Document & PersistentVolume {
    return this.isPersistentVolume();
}

anemos.Document.prototype.asPersistentVolumeClaim = function (this: anemos.Document): this is anemos.Document & PersistentVolumeClaim {
    return this.isPersistentVolumeClaim();
}

anemos.Document.prototype.asPod = function (this: anemos.Document): this is anemos.Document & Pod {
    return this.isPod();
}

anemos.Document.prototype.asReplicaSet = function (this: anemos.Document): this is anemos.Document & ReplicaSet {
    return this.isReplicaSet();
}

anemos.Document.prototype.asRole = function (this: anemos.Document): this is anemos.Document & Role {
    return this.isRole();
}

anemos.Document.prototype.asRoleBinding = function (this: anemos.Document): this is anemos.Document & RoleBinding {
    return this.isRoleBinding();
}

anemos.Document.prototype.asSecret = function (this: anemos.Document): this is anemos.Document & Secret {
    return this.isSecret();
}

anemos.Document.prototype.asService = function (this: anemos.Document): this is anemos.Document & Service {
    return this.isService();
}

anemos.Document.prototype.asServiceAccount = function (this: anemos.Document): this is anemos.Document & ServiceAccount {
    return this.isServiceAccount();
}

anemos.Document.prototype.asStatefulSet = function (this: anemos.Document): this is anemos.Document & StatefulSet {
    return this.isStatefulSet();
}

anemos.Document.prototype.asWorkload = function (this: anemos.Document): this is anemos.Document & anemos.Workload {
    return this.isWorkload();
}

anemos.Document.prototype.isOfKind = function (this: anemos.Document, apiVersion: string, kind: string): boolean {
    return this.apiVersion === apiVersion && this.kind === kind;
};

anemos.Document.prototype.isClusterRole = function (this: anemos.Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "ClusterRole");
};

anemos.Document.prototype.isClusterRoleBinding = function (this: anemos.Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "ClusterRoleBinding");
};

anemos.Document.prototype.isConfigMap = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "ConfigMap");
};

anemos.Document.prototype.isCRD = function (this: anemos.Document): boolean {
    return this.isCustomResourceDefinition();
};

anemos.Document.prototype.isCronJob = function (this: anemos.Document): boolean {
    return this.isOfKind("batch/v1", "CronJob");
};

anemos.Document.prototype.isCustomResourceDefinition = function (this: anemos.Document): boolean {
    return this.isOfKind("apiextensions.k8s.io/v1", "CustomResourceDefinition");
};

anemos.Document.prototype.isDaemonSet = function (this: anemos.Document): boolean {
    return this.isOfKind("apps/v1", "DaemonSet");
};

anemos.Document.prototype.isDeployment = function (this: anemos.Document): boolean {
    return this.isOfKind("apps/v1", "Deployment");
};

anemos.Document.prototype.isHorizontalPodAutoscaler = function (this: anemos.Document): boolean {
    return this.isOfKind("autoscaling/v2", "HorizontalPodAutoscaler");
};

anemos.Document.prototype.isIngress = function (this: anemos.Document): boolean {
    return this.isOfKind("networking.k8s.io/v1", "Ingress");
};

anemos.Document.prototype.isJob = function (this: anemos.Document): boolean {
    return this.isOfKind("batch/v1", "Job");
};

anemos.Document.prototype.isNamespace = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "Namespace");
};

anemos.Document.prototype.isPersistentVolume = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "PersistentVolume");
};

anemos.Document.prototype.isPersistentVolumeClaim = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "PersistentVolumeClaim");
};

anemos.Document.prototype.isPod = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "Pod");
};

anemos.Document.prototype.isReplicaSet = function (this: anemos.Document): boolean {
    return this.isOfKind("apps/v1", "ReplicaSet");
};

anemos.Document.prototype.isRole = function (this: anemos.Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "Role");
};

anemos.Document.prototype.isRoleBinding = function (this: anemos.Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "RoleBinding");
};

anemos.Document.prototype.isSecret = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "Secret");
};

anemos.Document.prototype.isService = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "Service");
};

anemos.Document.prototype.isServiceAccount = function (this: anemos.Document): boolean {
    return this.isOfKind("v1", "ServiceAccount");
};

anemos.Document.prototype.isStatefulSet = function (this: anemos.Document): boolean {
    return this.isOfKind("apps/v1", "StatefulSet");
};

anemos.Document.prototype.isWorkload = function (this: anemos.Document): boolean {
    return this.isCronJob()
        || this.isDaemonSet()
        || this.isDeployment()
        || this.isJob()
        || this.isPod()
        || this.isReplicaSet()
        || this.isStatefulSet();
};
