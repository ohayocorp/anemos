---
slug: /
sidebar_position: 1
---

# Overview

Anemos is a CLI tool designed to simplify the management of Kubernetes manifests and
auxiliary resources. It lets you manage these resources using JavaScript or TypeScript,
providing a powerful and flexible way to define your infrastructure as code.

Other tools typically focus on management of a single package. But, Kubernetes
environments often consist of many applications, each with its own set of dependencies and
configuration. These configurations are scattered across multiple places, making it
difficult to manage and maintain them.

Anemos takes a different approach and allows you to create and manage your infrastructure as a whole.
It lets you:

- Generate Kubernetes manifests using a programming language, supporting both template-based
  and programmatic approaches.
- Integrate and configure multiple packages in a single codebase.
- Handle multiple environments (e.g., development, staging, production) in the same project.
- Define and manage service requirements and dependencies effectively.
- Modify Kubernetes manifests individually or in bulk after they are generated.

[Install](installation) the Anemos CLI to get started and follow the
[tutorial](category/tutorial) to explore its core concepts and features.

## Enterprise Version

At [Ohayocorp](https://ohayocorp.com/) we know that architecting and deploying ecosystems of Kubernetes applications
across multiple environments and distributions can be complex and time-consuming. It requires deep knowledge of each
application, its dependencies, and the specific configurations needed for different environments.

That's why we offer an [enterprise version](https://ohayocorp.com/anemos-enterprise) of Anemos that further enhances the
capabilities of the open-source version and includes advanced features designed to simplify and speed up the development
and deployment of Kubernetes applications.

- Install many popular open-source tools to enhance your workflow. They are
  pre-configured for different environments (development, production, etc.)
  and different Kubernetes distributions.
- Anemos Enterprise brings tight integration between the installed open source tools,
  allowing you to use them to their full potential. It also makes it easy
  to integrate your own applications into the platform.
- On-premises & air-gapped customer environments have strict security and compliance
  requirements, without access to the public internet or even to the target Kubernetes
  clusters. Anemos Enterprise treats these environments as first-class citizens and
  provides many features to ease your workflow.
- Perform linting and security analysis on your Kubernetes manifests before applying them to your cluster.
  Generate reports for hardware requirements, requested permissions, list of container images, and more.
- Don't waste time configuring each component for each environment. Use Anemos Enterprise's
  bulk modification features, such as trusting CA certificates on pods or managing ingress TLS
  definitions, to make changes across all your components and environments at once.
- Easily generate load testing manifests and run them against your Kubernetes cluster.
  Anemos Enterprise provides powerful features to define load testing scenarios
  and configurations, run them in different environments and gather results for analysis.

These are just a few examples of the advanced features available in Anemos Enterprise. For more information, visit [Anemos Enterprise](https://ohayocorp.com/anemos-enterprise) or contact us at [anemos-enterprise@ohayocorp.com](mailto:anemos-enterprise@ohayocorp.com).