// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getAccessAddress";
export * from "./getGatewayNodes";
export * from "./getNacosReplicas";
export * from "./getNacosServerInterfaces";
export * from "./getZookeeperReplicas";
export * from "./getZookeeperServerInterfaces";
export * from "./instance";

// Import resources to register:
import { Instance } from "./instance";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Tse/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Tse/instance", _module)
