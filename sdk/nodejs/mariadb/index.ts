// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./dedicatedclusterDbInstance";
export * from "./encryptAttributes";
export * from "./getAccounts";
export * from "./getDatabaseObjects";
export * from "./getDatabaseTable";
export * from "./getDatabases";
export * from "./getDbInstances";
export * from "./getSecurityGroups";
export * from "./hourDbInstance";
export * from "./instance";
export * from "./logFileRetentionPeriod";
export * from "./parameters";
export * from "./securityGroups";

// Import resources to register:
import { Account } from "./account";
import { DedicatedclusterDbInstance } from "./dedicatedclusterDbInstance";
import { EncryptAttributes } from "./encryptAttributes";
import { HourDbInstance } from "./hourDbInstance";
import { Instance } from "./instance";
import { LogFileRetentionPeriod } from "./logFileRetentionPeriod";
import { Parameters } from "./parameters";
import { SecurityGroups } from "./securityGroups";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Mariadb/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "tencentcloud:Mariadb/dedicatedclusterDbInstance:DedicatedclusterDbInstance":
                return new DedicatedclusterDbInstance(name, <any>undefined, { urn })
            case "tencentcloud:Mariadb/encryptAttributes:EncryptAttributes":
                return new EncryptAttributes(name, <any>undefined, { urn })
            case "tencentcloud:Mariadb/hourDbInstance:HourDbInstance":
                return new HourDbInstance(name, <any>undefined, { urn })
            case "tencentcloud:Mariadb/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "tencentcloud:Mariadb/logFileRetentionPeriod:LogFileRetentionPeriod":
                return new LogFileRetentionPeriod(name, <any>undefined, { urn })
            case "tencentcloud:Mariadb/parameters:Parameters":
                return new Parameters(name, <any>undefined, { urn })
            case "tencentcloud:Mariadb/securityGroups:SecurityGroups":
                return new SecurityGroups(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/account", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/dedicatedclusterDbInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/encryptAttributes", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/hourDbInstance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/instance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/logFileRetentionPeriod", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/parameters", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Mariadb/securityGroups", _module)
