const anemos = require("@ohayocorp/anemos");
const app = require("../../lib");

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

app.add(builder, {
    name: "test-example-package",
    namespace: "test-namespace",
    image: "test-image:v1",
});

builder.build();