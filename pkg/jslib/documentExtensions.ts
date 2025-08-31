import * as anemos from "@ohayocorp/anemos"

declare module "@ohayocorp/anemos" {
    export interface Document {
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
