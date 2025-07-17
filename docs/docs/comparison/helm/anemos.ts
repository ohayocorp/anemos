// Compute if the server serviceaccount is enabled.
function serverServiceAccountEnabled(): boolean {
    return options.server.serviceAccount.create && (options.server.enabled || options.global.enabled);
}

function addServiceAccount(context: anemos.BuildContext): void {
    if (!serverServiceAccountEnabled()) {
        return;
    }

    const document = anemos.parseDocument(
        "server-serviceaccount.yaml",
        `
        apiVersion: v1
        kind: ServiceAccount
        metadata:
          name: ${options.name}
          namespace: ${options.namespace}
          labels:
            app.kubernetes.io/name: vault
            app.kubernetes.io/instance: ${options.name}
            app.kubernetes.io/managed-by: anemos
        `
    );

    // You can get labels from options and pass to the document. Alternatively users themselves can set
    // labels on generated documents.
    document.setLabels(options.labels);

    context.addDocument(document);
}