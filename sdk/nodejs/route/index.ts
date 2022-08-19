// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./table";
export * from "./tableEntry";

// Import resources to register:
import { Table } from "./table";
import { TableEntry } from "./tableEntry";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Route/table:Table":
                return new Table(name, <any>undefined, { urn })
            case "tencentcloud:Route/tableEntry:TableEntry":
                return new TableEntry(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Route/table", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Route/tableEntry", _module)
