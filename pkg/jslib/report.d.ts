export declare class ReportMetadata {
    filePath: string;

    constructor(filePath: string);
}

export declare class Report {
    metadata: ReportMetadata;
    markdownContent: string;

    constructor(metadata: ReportMetadata, markdownContent: string);
}