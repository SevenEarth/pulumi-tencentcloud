// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cluster";
export * from "./getAutoScaleRecords";
export * from "./getCvmQuota";
export * from "./getInstance";
export * from "./getNodes";
export * from "./userManager";

// Import resources to register:
import { Cluster } from "./cluster";
import { UserManager } from "./userManager";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Emr/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "tencentcloud:Emr/userManager:UserManager":
                return new UserManager(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Emr/cluster", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Emr/userManager", _module)
