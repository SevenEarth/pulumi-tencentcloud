// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./function";
export * from "./getFunctions";
export * from "./getLogs";
export * from "./getNamespaces";
export * from "./layer";
export * from "./namespace";

// Import resources to register:
import { Function } from "./function";
import { Layer } from "./layer";
import { Namespace } from "./namespace";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Scf/function:Function":
                return new Function(name, <any>undefined, { urn })
            case "tencentcloud:Scf/layer:Layer":
                return new Layer(name, <any>undefined, { urn })
            case "tencentcloud:Scf/namespace:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Scf/function", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Scf/layer", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Scf/namespace", _module)
