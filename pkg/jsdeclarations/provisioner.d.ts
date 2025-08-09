import { DocumentGroup } from "./documentGroup";


export declare enum ProvisionerType {
    Apply,
    Wait
}

export declare class Provisioner {
    private constructor();

    type: ProvisionerType;
    documentGroup: DocumentGroup;

    runAfter(provisioner: Provisioner): void;
    runBefore(provisioner: Provisioner): void;
}