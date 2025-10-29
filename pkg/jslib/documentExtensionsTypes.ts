import { Document } from "@ohayocorp/anemos/document";
import { Workload } from "./documentExtensionsWorkload";
import { CustomResourceDefinition } from "./k8s/apiextensions/v1";
import { DaemonSet, Deployment, ReplicaSet, StatefulSet } from "./k8s/apps/v1";
import { HorizontalPodAutoscaler } from "./k8s/autoscaling/v2";
import { CronJob, Job } from "./k8s/batch/v1";
import { ConfigMap, Namespace, PersistentVolume, PersistentVolumeClaim, Pod, Secret, Service, ServiceAccount } from "./k8s/core/v1";
import { Ingress } from "./k8s/networking/v1";
import { ClusterRole, ClusterRoleBinding, Role, RoleBinding } from "./k8s/rbac/v1";

declare module "@ohayocorp/anemos/document" {
    export interface Document {
        /** Type guard for ClusterRole documents. */
        asClusterRole(): this is Document & ClusterRole;

        /** Type guard for ClusterRoleBinding documents. */
        asClusterRoleBinding(): this is Document & ClusterRoleBinding;

        /** Type guard for ConfigMap documents. */
        asConfigMap(): this is Document & ConfigMap;

        /** Type guard for CustomResourceDefinition documents. */
        asCRD(): this is Document & CustomResourceDefinition;

        /** Type guard for CronJob documents. */
        asCronJob(): this is Document & CronJob;

        /** Type guard for CustomResourceDefinition documents. */
        asCustomResourceDefinition(): this is Document & CustomResourceDefinition;

        /** Type guard for DaemonSet documents. */
        asDaemonSet(): this is Document & DaemonSet;

        /** Type guard for Deployment documents. */
        asDeployment(): this is Document & Deployment;

        /** Type guard for HorizontalPodAutoscaler documents. */
        asHorizontalPodAutoscaler(): this is Document & HorizontalPodAutoscaler;

        /** Type guard for Ingress documents. */
        asIngress(): this is Document & Ingress;

        /** Type guard for Job documents. */
        asJob(): this is Document & Job;

        /** Type guard for Namespace documents. */
        asNamespace(): this is Document & Namespace;

        /** Type guard for PersistentVolume documents. */
        asPersistentVolume(): this is Document & PersistentVolume;

        /** Type guard for PersistentVolumeClaim documents. */
        asPersistentVolumeClaim(): this is Document & PersistentVolumeClaim;

        /** Type guard for Pod documents. */
        asPod(): this is Document & Pod;

        /** Type guard for ReplicaSet documents. */
        asReplicaSet(): this is Document & ReplicaSet;

        /** Type guard for Role documents. */
        asRole(): this is Document & Role;

        /** Type guard for RoleBinding documents. */
        asRoleBinding(): this is Document & RoleBinding;

        /** Type guard for Secret documents. */
        asSecret(): this is Document & Secret;

        /** Type guard for Service documents. */
        asService(): this is Document & Service;

        /** Type guard for ServiceAccount documents. */
        asServiceAccount(): this is Document & ServiceAccount;

        /** Type guard for StatefulSet documents. */
        asStatefulSet(): this is Document & StatefulSet;

        /** Type guard for workload documents. */
        asWorkload(): this is Document & Workload;

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

Document.prototype.asClusterRole = function (this: Document): this is Document & ClusterRole {
    return this.isClusterRole();
}

Document.prototype.asClusterRoleBinding = function (this: Document): this is Document & ClusterRoleBinding {
    return this.isClusterRoleBinding();
}

Document.prototype.asConfigMap = function (this: Document): this is Document & ConfigMap {
    return this.isConfigMap();
}

Document.prototype.asCRD = function (this: Document): this is Document & CustomResourceDefinition {
    return this.isCRD();
}

Document.prototype.asCronJob = function (this: Document): this is Document & CronJob {
    return this.isCronJob();
}

Document.prototype.asCustomResourceDefinition = function (this: Document): this is Document & CustomResourceDefinition {
    return this.isCustomResourceDefinition();
}

Document.prototype.asDaemonSet = function (this: Document): this is Document & DaemonSet {
    return this.isDaemonSet();
}

Document.prototype.asDeployment = function (this: Document): this is Document & Deployment {
    return this.isDeployment();
}

Document.prototype.asHorizontalPodAutoscaler = function (this: Document): this is Document & HorizontalPodAutoscaler {
    return this.isHorizontalPodAutoscaler();
}

Document.prototype.asIngress = function (this: Document): this is Document & Ingress {
    return this.isIngress();
}

Document.prototype.asJob = function (this: Document): this is Document & Job {
    return this.isJob();
}

Document.prototype.asNamespace = function (this: Document): this is Document & Namespace {
    return this.isNamespace();
}

Document.prototype.asPersistentVolume = function (this: Document): this is Document & PersistentVolume {
    return this.isPersistentVolume();
}

Document.prototype.asPersistentVolumeClaim = function (this: Document): this is Document & PersistentVolumeClaim {
    return this.isPersistentVolumeClaim();
}

Document.prototype.asPod = function (this: Document): this is Document & Pod {
    return this.isPod();
}

Document.prototype.asReplicaSet = function (this: Document): this is Document & ReplicaSet {
    return this.isReplicaSet();
}

Document.prototype.asRole = function (this: Document): this is Document & Role {
    return this.isRole();
}

Document.prototype.asRoleBinding = function (this: Document): this is Document & RoleBinding {
    return this.isRoleBinding();
}

Document.prototype.asSecret = function (this: Document): this is Document & Secret {
    return this.isSecret();
}

Document.prototype.asService = function (this: Document): this is Document & Service {
    return this.isService();
}

Document.prototype.asServiceAccount = function (this: Document): this is Document & ServiceAccount {
    return this.isServiceAccount();
}

Document.prototype.asStatefulSet = function (this: Document): this is Document & StatefulSet {
    return this.isStatefulSet();
}

Document.prototype.asWorkload = function (this: Document): this is Document & Workload {
    return this.isWorkload();
}

Document.prototype.isOfKind = function (this: Document, apiVersion: string, kind: string): boolean {
    return this.apiVersion === apiVersion && this.kind === kind;
};

Document.prototype.isClusterRole = function (this: Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "ClusterRole");
};

Document.prototype.isClusterRoleBinding = function (this: Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "ClusterRoleBinding");
};

Document.prototype.isConfigMap = function (this: Document): boolean {
    return this.isOfKind("v1", "ConfigMap");
};

Document.prototype.isCRD = function (this: Document): boolean {
    return this.isCustomResourceDefinition();
};

Document.prototype.isCronJob = function (this: Document): boolean {
    return this.isOfKind("batch/v1", "CronJob");
};

Document.prototype.isCustomResourceDefinition = function (this: Document): boolean {
    return this.isOfKind("apiextensions.k8s.io/v1", "CustomResourceDefinition");
};

Document.prototype.isDaemonSet = function (this: Document): boolean {
    return this.isOfKind("apps/v1", "DaemonSet");
};

Document.prototype.isDeployment = function (this: Document): boolean {
    return this.isOfKind("apps/v1", "Deployment");
};

Document.prototype.isHorizontalPodAutoscaler = function (this: Document): boolean {
    return this.isOfKind("autoscaling/v2", "HorizontalPodAutoscaler");
};

Document.prototype.isIngress = function (this: Document): boolean {
    return this.isOfKind("networking.k8s.io/v1", "Ingress");
};

Document.prototype.isJob = function (this: Document): boolean {
    return this.isOfKind("batch/v1", "Job");
};

Document.prototype.isNamespace = function (this: Document): boolean {
    return this.isOfKind("v1", "Namespace");
};

Document.prototype.isPersistentVolume = function (this: Document): boolean {
    return this.isOfKind("v1", "PersistentVolume");
};

Document.prototype.isPersistentVolumeClaim = function (this: Document): boolean {
    return this.isOfKind("v1", "PersistentVolumeClaim");
};

Document.prototype.isPod = function (this: Document): boolean {
    return this.isOfKind("v1", "Pod");
};

Document.prototype.isReplicaSet = function (this: Document): boolean {
    return this.isOfKind("apps/v1", "ReplicaSet");
};

Document.prototype.isRole = function (this: Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "Role");
};

Document.prototype.isRoleBinding = function (this: Document): boolean {
    return this.isOfKind("rbac.authorization.k8s.io/v1", "RoleBinding");
};

Document.prototype.isSecret = function (this: Document): boolean {
    return this.isOfKind("v1", "Secret");
};

Document.prototype.isService = function (this: Document): boolean {
    return this.isOfKind("v1", "Service");
};

Document.prototype.isServiceAccount = function (this: Document): boolean {
    return this.isOfKind("v1", "ServiceAccount");
};

Document.prototype.isStatefulSet = function (this: Document): boolean {
    return this.isOfKind("apps/v1", "StatefulSet");
};

Document.prototype.isWorkload = function (this: Document): boolean {
    return this.isCronJob()
        || this.isDaemonSet()
        || this.isDeployment()
        || this.isJob()
        || this.isPod()
        || this.isReplicaSet()
        || this.isStatefulSet();
};
