/**
 * Predefined types for the Kubernetes distribution. These are used to determine the behavior of some components.
 */
type KubernetesDistribution = string;

export const unknown: KubernetesDistribution;
export const aks: KubernetesDistribution;
export const eks: KubernetesDistribution;
export const gke: KubernetesDistribution;
export const k3s: KubernetesDistribution;
export const kubeadm: KubernetesDistribution;
export const microk8s: KubernetesDistribution;
export const minikube: KubernetesDistribution;
export const openshift: KubernetesDistribution;