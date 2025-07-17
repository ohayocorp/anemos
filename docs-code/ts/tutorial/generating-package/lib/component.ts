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

    sanitize = (context: anemos.BuildContext) => {
        this.options.name ??= "example-package";
        this.options.namespace ??= "default";
        this.options.image ??= "nginx:latest";
    }

    generateResources = (context: anemos.BuildContext) => {
        const pod = new anemos.Document(
            `pod.yaml`,
            `
            apiVersion: v1
            kind: Pod
            metadata:
              name: ${this.options.name}
              namespace: ${this.options.namespace}
              labels:
                app: ${this.options.name}
            spec:
              containers:
              - name: nginx
                image: ${this.options.image}
            `);

        const documentGroup = new anemos.DocumentGroup(this.options.name!);
        documentGroup.addDocument(pod);

        context.addDocumentGroup(documentGroup);
    }
}