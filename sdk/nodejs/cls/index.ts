// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./config";
export * from "./configAttachment";
export * from "./configExtra";
export * from "./cosShipper";
export * from "./index_";
export * from "./logset";
export * from "./machineGroup";
export * from "./topic";

// Import resources to register:
import { Config } from "./config";
import { ConfigAttachment } from "./configAttachment";
import { ConfigExtra } from "./configExtra";
import { CosShipper } from "./cosShipper";
import { Index } from "./index_";
import { Logset } from "./logset";
import { MachineGroup } from "./machineGroup";
import { Topic } from "./topic";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Cls/config:Config":
                return new Config(name, <any>undefined, { urn })
            case "tencentcloud:Cls/configAttachment:ConfigAttachment":
                return new ConfigAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cls/configExtra:ConfigExtra":
                return new ConfigExtra(name, <any>undefined, { urn })
            case "tencentcloud:Cls/cosShipper:CosShipper":
                return new CosShipper(name, <any>undefined, { urn })
            case "tencentcloud:Cls/index:Index":
                return new Index(name, <any>undefined, { urn })
            case "tencentcloud:Cls/logset:Logset":
                return new Logset(name, <any>undefined, { urn })
            case "tencentcloud:Cls/machineGroup:MachineGroup":
                return new MachineGroup(name, <any>undefined, { urn })
            case "tencentcloud:Cls/topic:Topic":
                return new Topic(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/config", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/configAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/configExtra", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/cosShipper", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/index", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/logset", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/machineGroup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cls/topic", _module)
