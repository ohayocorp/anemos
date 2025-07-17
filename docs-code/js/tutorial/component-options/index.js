const anemos = require("@ohayocorp/anemos");
const Component = require("./component");

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

builder.addComponent(new Component({
    image: "nginx:1.27",
    replicas: 3,
}));

builder.build();