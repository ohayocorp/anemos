# Anemos

Anemos is a CLI tool designed to simplify the management of Kubernetes manifests. It lets you define your
manifests using JavaScript and TypeScript, enabling you to leverage the full power of the language to generate
and manipulate Kubernetes YAML manifests.

Anemos uses JavaScript package management system. You can distribute your libraries as npm packages,
use other Anemos packages, or any other JavaScript library, and take advantage of the vast ecosystem of tools
available in the JavaScript community.

Anemos provides you an SDK for both template based and programmatic generation of Kubernetes manifests.
JavaScript's template literals allow you to write clean and readable templates and libraries such as
[kubernetes-models](https://www.npmjs.com/package/kubernetes-models) provide type safety and better
IDE experience. You can also mix and match templates and programmatic generation, use the best approach
for each part of your manifest.

Anemos also supports node based modification of YAML manifests. This powerful feature enables you
to modify any YAML manifest without waiting for the package maintainers to add support for your use case
or fix a bug in the package. It also allows you to modify manifests in bulk, making it easy to
apply changes across multiple manifests.

## Installation

Anemos is a single binary that can be downloaded from the [GitHub releases page](https://github.com/ohayocorp/anemos/releases).
Get the latest version and place it in your PATH and you are ready to go.

## Docs

You can visit [documentation](https://ohayocorp.com/anemos/docs) for more information on how to use Anemos,
including examples, tutorials, and API references. You can also browse the documentation locally by running `anemos docs`.

## Quick Start

Simplest way to get started with Anemos is to create an `index.js` file with the following content and run it with
`anemos build index.js`:

```javascript
const anemos = require("@ohayocorp/anemos");

const builder = new anemos.Builder("1.31", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

builder.addDocument(
    `pod.yaml`,
    `
    apiVersion: v1
    kind: Pod
    metadata:
      name: nginx
    spec:
      containers:
        - name: nginx
          image: nginx
          ports:
            - containerPort: 80
    `);

builder.build();
```

This will generate an `output` directory containing the Kubernetes manifests defined in `pod.yaml`.

### Using Anemos Packages

Following example will use the `@ohayocorp/anemos-hello-world` package to generate a simple deployment:

```javascript
const anemos = require("@ohayocorp/anemos");
const helloWorld = require("@ohayocorp/anemos-hello-world");

const builder = new anemos.Builder("1.32", anemos.KubernetesDistribution.Minikube, anemos.EnvironmentType.Development);

helloWorld.add(builder, {
    name: "custom-hello-world",
    autoScaling: {
        minReplicas: 1,
        maxReplicas: 3,
    },
    ingress: {
        host: "hello-world.local",
    },
});

builder.build();
```

Run following command to install the package and generate the manifests:

```bash
anemos package add @ohayocorp/anemos-hello-world
anemos build index.js
```

## Contributing

We welcome contributions to Anemos! If you have an idea for a new feature, a bug fix, or an improvement, please
open an issue or a pull request on our [GitHub repository](https://github.com/ohayocorp/anemos).

To build Anemos from source, clone the repository and run the following commands:

```bash
./download-bun.sh
go build ./cmd/anemos
```

If you wish to contribute to the documentation, you should build the docs-code and then
run the development server:

```bash
# This command requires anemos binary to be in your PATH
./build-docs-project.sh
cd docs
# Ensure you have NodeJS installed.
npm install
npm run start
```
