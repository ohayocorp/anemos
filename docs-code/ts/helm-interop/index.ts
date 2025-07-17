import * as anemos from "@ohayocorp/anemos";

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

builder.onStep(anemos.steps.generateResources, (context: anemos.BuildContext) => {
    // Download hello world Helm chart to the directory where the index.js file, that is created
    // by the tsc compiler, is located before running this script:
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
        document.getLabels()!.remove("helm.sh/chart");
        document.setLabel("example", "value");
    }

    context.addDocumentGroup(documentGroup);
});

builder.build();