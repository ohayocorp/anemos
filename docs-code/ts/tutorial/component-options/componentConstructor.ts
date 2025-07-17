import * as anemos from "@ohayocorp/anemos";
import { Options } from "./options";

export class Component extends anemos.Component {
    // highlight-next-line
    options: Options;

    // highlight-next-line
    constructor(options?: Options) {
        super();

        // highlight-next-line
        this.options = options ?? {};
        
        this.addAction(anemos.steps.generateResources, this.generateResources);
    }

    generateResources = (context: anemos.BuildContext) => {
        const name = "example-app";
        const namespace = "default";
        const image = "nginx";
        const replicas = 1;

        context.addDocument(
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

        context.addDocument(
            `service.yaml`,
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
    };
}