const anemos = require("@ohayocorp/anemos");

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

// highlight-next-line
const name = "example-app2";
const namespace = "default";
const image = "nginx";
const replicas = 1;

builder.addDocument(
    `deployment.yaml`,
    `
    apiVersion: apps/v1
    kind: Deployment
    metadata:
      name: ${name}
      namespace: ${namespace}
    spec:
      replicas: ${replicas}
      selector:
        matchLabels:
          app: ${name}
      template:
        metadata:
          labels:
            app: ${name}
        spec:
          containers:
            - name: app
              image: ${image}
              ports:
                - containerPort: 80
    `);

builder.addDocument(
    "service.yaml",
    {
        apiVersion: "v1",
        kind: "Service",
        metadata: {
            name: name,
            namespace: namespace,
        },
        spec: {
            selector: {
                app: name
            },
            ports: [
                {
                    protocol: "TCP",
                    port: 80,
                    targetPort: 80
                }
            ]
        }
    });

builder.build();