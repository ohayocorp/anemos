import * as anemos from "@ohayocorp/anemos";

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

builder.addDocument(
    `pod.yaml`,
    `
    apiVersion: v1
    kind: Pod
    metadata:
      name: nginx
      namespace: default
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:latest
          ports:
            - containerPort: 80
    `);

builder.build();