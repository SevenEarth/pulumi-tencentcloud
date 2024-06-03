// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { DiskBackupArgs, DiskBackupState } from "./diskBackup";
export type DiskBackup = import("./diskBackup").DiskBackup;
export const DiskBackup: typeof import("./diskBackup").DiskBackup = null as any;
utilities.lazyLoad(exports, ["DiskBackup"], () => require("./diskBackup"));

export { DiskBackupRollbackOperationArgs, DiskBackupRollbackOperationState } from "./diskBackupRollbackOperation";
export type DiskBackupRollbackOperation = import("./diskBackupRollbackOperation").DiskBackupRollbackOperation;
export const DiskBackupRollbackOperation: typeof import("./diskBackupRollbackOperation").DiskBackupRollbackOperation = null as any;
utilities.lazyLoad(exports, ["DiskBackupRollbackOperation"], () => require("./diskBackupRollbackOperation"));

export { GetSnapshotPoliciesArgs, GetSnapshotPoliciesResult, GetSnapshotPoliciesOutputArgs } from "./getSnapshotPolicies";
export const getSnapshotPolicies: typeof import("./getSnapshotPolicies").getSnapshotPolicies = null as any;
export const getSnapshotPoliciesOutput: typeof import("./getSnapshotPolicies").getSnapshotPoliciesOutput = null as any;
utilities.lazyLoad(exports, ["getSnapshotPolicies","getSnapshotPoliciesOutput"], () => require("./getSnapshotPolicies"));

export { GetSnapshotsArgs, GetSnapshotsResult, GetSnapshotsOutputArgs } from "./getSnapshots";
export const getSnapshots: typeof import("./getSnapshots").getSnapshots = null as any;
export const getSnapshotsOutput: typeof import("./getSnapshots").getSnapshotsOutput = null as any;
utilities.lazyLoad(exports, ["getSnapshots","getSnapshotsOutput"], () => require("./getSnapshots"));

export { GetStoragesArgs, GetStoragesResult, GetStoragesOutputArgs } from "./getStorages";
export const getStorages: typeof import("./getStorages").getStorages = null as any;
export const getStoragesOutput: typeof import("./getStorages").getStoragesOutput = null as any;
utilities.lazyLoad(exports, ["getStorages","getStoragesOutput"], () => require("./getStorages"));

export { GetStoragesSetArgs, GetStoragesSetResult, GetStoragesSetOutputArgs } from "./getStoragesSet";
export const getStoragesSet: typeof import("./getStoragesSet").getStoragesSet = null as any;
export const getStoragesSetOutput: typeof import("./getStoragesSet").getStoragesSetOutput = null as any;
utilities.lazyLoad(exports, ["getStoragesSet","getStoragesSetOutput"], () => require("./getStoragesSet"));

export { SnapshotArgs, SnapshotState } from "./snapshot";
export type Snapshot = import("./snapshot").Snapshot;
export const Snapshot: typeof import("./snapshot").Snapshot = null as any;
utilities.lazyLoad(exports, ["Snapshot"], () => require("./snapshot"));

export { SnapshotPolicyArgs, SnapshotPolicyState } from "./snapshotPolicy";
export type SnapshotPolicy = import("./snapshotPolicy").SnapshotPolicy;
export const SnapshotPolicy: typeof import("./snapshotPolicy").SnapshotPolicy = null as any;
utilities.lazyLoad(exports, ["SnapshotPolicy"], () => require("./snapshotPolicy"));

export { SnapshotPolicyAttachmentArgs, SnapshotPolicyAttachmentState } from "./snapshotPolicyAttachment";
export type SnapshotPolicyAttachment = import("./snapshotPolicyAttachment").SnapshotPolicyAttachment;
export const SnapshotPolicyAttachment: typeof import("./snapshotPolicyAttachment").SnapshotPolicyAttachment = null as any;
utilities.lazyLoad(exports, ["SnapshotPolicyAttachment"], () => require("./snapshotPolicyAttachment"));

export { SnapshotSharePermissionArgs, SnapshotSharePermissionState } from "./snapshotSharePermission";
export type SnapshotSharePermission = import("./snapshotSharePermission").SnapshotSharePermission;
export const SnapshotSharePermission: typeof import("./snapshotSharePermission").SnapshotSharePermission = null as any;
utilities.lazyLoad(exports, ["SnapshotSharePermission"], () => require("./snapshotSharePermission"));

export { StorageArgs, StorageState } from "./storage";
export type Storage = import("./storage").Storage;
export const Storage: typeof import("./storage").Storage = null as any;
utilities.lazyLoad(exports, ["Storage"], () => require("./storage"));

export { StorageAttachmentArgs, StorageAttachmentState } from "./storageAttachment";
export type StorageAttachment = import("./storageAttachment").StorageAttachment;
export const StorageAttachment: typeof import("./storageAttachment").StorageAttachment = null as any;
utilities.lazyLoad(exports, ["StorageAttachment"], () => require("./storageAttachment"));

export { StorageSetArgs, StorageSetState } from "./storageSet";
export type StorageSet = import("./storageSet").StorageSet;
export const StorageSet: typeof import("./storageSet").StorageSet = null as any;
utilities.lazyLoad(exports, ["StorageSet"], () => require("./storageSet"));

export { StorageSetAttachmentArgs, StorageSetAttachmentState } from "./storageSetAttachment";
export type StorageSetAttachment = import("./storageSetAttachment").StorageSetAttachment;
export const StorageSetAttachment: typeof import("./storageSetAttachment").StorageSetAttachment = null as any;
utilities.lazyLoad(exports, ["StorageSetAttachment"], () => require("./storageSetAttachment"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Cbs/diskBackup:DiskBackup":
                return new DiskBackup(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/diskBackupRollbackOperation:DiskBackupRollbackOperation":
                return new DiskBackupRollbackOperation(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshot:Snapshot":
                return new Snapshot(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshotPolicy:SnapshotPolicy":
                return new SnapshotPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshotPolicyAttachment:SnapshotPolicyAttachment":
                return new SnapshotPolicyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cbs/snapshotSharePermission:SnapshotSharePermission":
                return new SnapshotSharePermission(name, <any>undefined, { urn })
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
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/diskBackup", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/diskBackupRollbackOperation", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshot", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshotPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshotPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/snapshotSharePermission", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storage", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storageAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storageSet", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cbs/storageSetAttachment", _module)
