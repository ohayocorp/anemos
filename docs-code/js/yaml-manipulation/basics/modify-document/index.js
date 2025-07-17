const anemos = require("@ohayocorp/anemos");

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

const document = new anemos.Document("modified.yaml");

// Set simple values.
const root = document.getRoot();

root.set("apiVersion", "apps/v1");
root.set("kind", "Deployment");

// Set objects as values.
root.set("metadata", {
    name: "modified",
    namespace: "default",
});

// Create chain of objects. Here "spec", "template" and "spec" are all created and the last mapping
// is returned.
const templateSpec = root.ensureMapping(["spec", "template", "spec"]);

// Set an array as YAML sequence.
templateSpec.set("containers", [
    {
        name: "modified",
        image: "nginx:latest",
    },
]);

// Anemos has some convenience methods for Kubernetes objects. For example, `getContainer` returns the
// container of a workload with given name.
const container = document.getContainer("modified");
container?.set("imagePullPolicy", "Always");

builder.addDocument(document);
builder.build();