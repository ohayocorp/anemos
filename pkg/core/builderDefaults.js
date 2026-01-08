(context => {
    const anemos = require("@ohayocorp/anemos");
    const builder = context.builder;

    builder.deleteOutputDirectory();
    builder.reportDiagnostics();
    builder.writeDocuments();
    builder.writeReports();

    anemos.sortFields.add(builder);
    anemos.setDefaultProvisionerDependencies.add(builder);
    anemos.collectCRDs.add(builder);
    anemos.collectNamespaces.add(builder);

    anemos.diagnostics.addDefaultDiagnostics(builder);
    anemos.reports.addDefaultReports(builder);

    if (context.apply) {
        builder.apply({
            skipConfirmation: context.skipConfirmation,
            forceConflicts: context.forceConflicts,
            documentGroups: context.documentGroups
        });
    }
})(__anemos__context);