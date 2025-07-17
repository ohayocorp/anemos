import * as anemos from "@ohayocorp/anemos";

class App1Component extends anemos.Component {
    constructor() {
        super();
        this.addAction(anemos.steps.generateResources, this.generateResources);
    }

    generateResources = (context: anemos.BuildContext) => {
        context.addDocument(
            `app1-pod.yaml`,
            `
            apiVersion: v1
            kind: Pod
            metadata:
              name: app1
              labels:
                app: app1
            spec:
              containers:
                - name: app1
                  image: app1:latest
            `);
    };
}

class App2Component extends anemos.Component {
    constructor() {
        super();
        this.addAction(anemos.steps.generateResources, this.generateResources);
    }

    generateResources = (context: anemos.BuildContext) => {
        context.addDocument(
            `app2-pod.yaml`,
            `
            apiVersion: v1
            kind: Pod
            metadata:
              name: app2
              labels:
                app: app2
            spec:
              containers:
                - name: app2
                  image: app2:latest
            `);
    };
}

class SetNodeNameComponent extends anemos.Component {
    constructor() {
        super();
        this.addAction(anemos.steps.modify, this.modifyResources);
    }

    modifyResources = (context: anemos.BuildContext) => {
        for (const document of context.getAllDocuments()) {
            // Check if the document is a workload (e.g. Pod, Deployment, etc.) and skip if not.
            if (!document.isWorkload()) {
                continue;
            }

            // If the document does not have the label "app: app1", skip it.
            if (document.getLabel("app") !== "app1") {
                continue;
            }
            
            // Run app1 pods on node "worker-1".
            document.ensureWorkloadSpec().set("nodeName", "worker-1");
        }
    };
}

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

builder.addComponent(new SetNodeNameComponent());
builder.addComponent(new App1Component());
builder.addComponent(new App2Component());

builder.build();