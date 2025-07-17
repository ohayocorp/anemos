const anemos = require("@ohayocorp/anemos");

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

builder.onStep(anemos.steps.generateResources, (context) => {
    // Download hello world Helm chart to the current script directory before running this script:
    // https://github.com/crowdsecurity/helm-charts/releases/download/helloworld-0.2.1/helloworld-0.2.1.tgz
    const chart = new anemos.HelmChart(`${anemos.file.currentScriptDirectory()}/helloworld-0.2.1.tgz`);
    const values = anemos.dedent(`
        replicaCount: 2
        image:
          repository: hashicorp/http-echo
          tag: 1.0.0
        `);

    const documentGroup = chart.generate(context, {
        releaseName: "hello-world",
        namespace: "default",
        values: values
    });

    for (const document of documentGroup.documents) {
        document.getLabels().remove("helm.sh/chart");
        document.setLabel("example", "value");
    }

    context.addDocumentGroup(documentGroup);
});

builder.build();