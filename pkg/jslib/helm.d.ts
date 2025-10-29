import { BuildContext } from "./buildContext";
import { DocumentGroup } from "./documentGroup";
import * as steps from "./steps";

declare module "./builder" {
    export interface Builder {
        /**
         * Creates a document group from the Helm chart using the given values on
         * {@link steps.generateResources} step. Chart identifier can be a local path or a URL.
         */
        addHelmChart(chartIdentifier: string, releaseName: string, values?: string): void;
    }
}

/**
 * Options for generating Helm charts.
 * @param releaseName The name of the Helm release.
 * @param namespace The namespace to use for the Helm release.
 * @param values Optional values file to use for the Helm chart.
 */
export class HelmOptions {
    constructor(releaseName: string, namespace: string, values?: string);

    /** The name of the Helm release. */
    releaseName: string;

    /** The namespace to use for the Helm release. */
    namespace: string;

    /** Optional values file to use for the Helm chart. */
    values?: string;
}

/**
 * Represents a Helm chart that can be used to generate Kubernetes documents.
 * It can be initialized with a file path or a byte array containing the chart data.
 * The {@link HelmChart.generate} method creates documents from the Helm chart using
 * the provided context and options.
 */
export class HelmChart {
    constructor(path: string);
    constructor(data: Uint8Array);

    /** Creates a document group from the Helm chart using the provided options. */
    generate(context: BuildContext, options: HelmOptions): DocumentGroup;
}