// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getSnapshotPolicies";
export * from "./getSnapshots";
export * from "./getStorages";
export * from "./getStoragesSet";
export * from "./snapshot";
export * from "./snapshotPolicy";
export * from "./snapshotPolicyAttachment";
export * from "./storage";
export * from "./storageAttachment";
export * from "./storageSet";
export * from "./storageSetAttachment";

// Import resources to register:
import { Snapshot } from "./snapshot";
import { SnapshotPolicy } from "./snapshotPolicy";
import { SnapshotPolicyAttachment } from "./snapshotPolicyAttachment";
import { Storage } from "./storage";
import { StorageAttachment } from "./storageAttachment";
import { StorageSet } from "./storageSet";
import { StorageSetAttachment } from "./storageSetAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Cbs/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy":
                return new SnapshotPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment":
                return new SnapshotPolicyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/storage:Storage":
                return new Storage(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/storageAttachment:StorageAttachment":
                return new StorageAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/storageSet:StorageSet":
                return new StorageSet(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/storageSetAttachment:StorageSetAttachment":
                return new StorageSetAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshot", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshotPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshotPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storage", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storageAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storageSet", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storageSetAttachment", _module)
