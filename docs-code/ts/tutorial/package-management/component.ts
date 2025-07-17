import * as anemos from "@ohayocorp/anemos";
import { Options } from "./options";

export class Component extends anemos.Component {
    options: Options;

    constructor(options?: Options) {
        super();

        this.options = options ?? {};
        
        this.addAction(anemos.steps.sanitize, this.sanitize);
        this.addAction(anemos.steps.generateResources, this.generateResources);
    }

    sanitize = () => {
        // Assign default values to the options if they are not provided.
        const options = this.options;

        options.name ??= "example-app";
        options.namespace ??= "default";
        options.image ??= "nginx";
        options.replicas ??= 1;
    };

    generateResources = (context: anemos.BuildContext) => {
        const name = this.options.name;
        const namespace = this.options.namespace;
        const image = this.options.image;
        const replicas = this.options.replicas;

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