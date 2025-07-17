const anemos = require("@ohayocorp/anemos");

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

// Create an empty document. The root of a document is set to an empty `Mapping` by default.
const empty = new anemos.Document("empty.yaml");
builder.addDocument(empty);

// Parse a document from a string.
// Template strings can be used to create multi-line strings. The common
// indentation is stripped from the string, so the contents can be indented freely.
// It is also possible to use interpolation to insert variables into the string.
const name = "mysecret";
const namespace = "default";

const parsed = anemos.parseDocument(
    `folder/parsed.yaml`,
    `
    apiVersion: v1
    kind: Secret
    metadata:
      name: "${name}"
      namespace: "${namespace}"
    type: Opaque
    `);

builder.addDocument(parsed);

// Create a document from an object. Anemos doesn't provide an API for Kubernetes
// object models, but it is possible to use other libraries such as:
//   https://www.npmjs.com/package/kubernetes-models
//   https://www.npmjs.com/package/kubernetes-types

const fromObject = new anemos.Document(
    `from-object.yaml`,
    {
        apiVersion: "v1",
        kind: "Secret",
        metadata: {
            name: name,
            namespace: namespace,
        },
        type: "Opaque",
    });

builder.addDocument(fromObject);

builder.build();