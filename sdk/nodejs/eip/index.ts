// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AddressTransformArgs, AddressTransformState } from "./addressTransform";
export type AddressTransform = import("./addressTransform").AddressTransform;
export const AddressTransform: typeof import("./addressTransform").AddressTransform = null as any;
utilities.lazyLoad(exports, ["AddressTransform"], () => require("./addressTransform"));

export { AssociationArgs, AssociationState } from "./association";
export type Association = import("./association").Association;
export const Association: typeof import("./association").Association = null as any;
utilities.lazyLoad(exports, ["Association"], () => require("./association"));

export { GetAddressQuotaArgs, GetAddressQuotaResult, GetAddressQuotaOutputArgs } from "./getAddressQuota";
export const getAddressQuota: typeof import("./getAddressQuota").getAddressQuota = null as any;
export const getAddressQuotaOutput: typeof import("./getAddressQuota").getAddressQuotaOutput = null as any;
utilities.lazyLoad(exports, ["getAddressQuota","getAddressQuotaOutput"], () => require("./getAddressQuota"));

export { GetNetworkAccountTypeArgs, GetNetworkAccountTypeResult, GetNetworkAccountTypeOutputArgs } from "./getNetworkAccountType";
export const getNetworkAccountType: typeof import("./getNetworkAccountType").getNetworkAccountType = null as any;
export const getNetworkAccountTypeOutput: typeof import("./getNetworkAccountType").getNetworkAccountTypeOutput = null as any;
utilities.lazyLoad(exports, ["getNetworkAccountType","getNetworkAccountTypeOutput"], () => require("./getNetworkAccountType"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { NormalAddressReturnArgs, NormalAddressReturnState } from "./normalAddressReturn";
export type NormalAddressReturn = import("./normalAddressReturn").NormalAddressReturn;
export const NormalAddressReturn: typeof import("./normalAddressReturn").NormalAddressReturn = null as any;
utilities.lazyLoad(exports, ["NormalAddressReturn"], () => require("./normalAddressReturn"));

export { PublicAddressAdjustArgs, PublicAddressAdjustState } from "./publicAddressAdjust";
export type PublicAddressAdjust = import("./publicAddressAdjust").PublicAddressAdjust;
export const PublicAddressAdjust: typeof import("./publicAddressAdjust").PublicAddressAdjust = null as any;
utilities.lazyLoad(exports, ["PublicAddressAdjust"], () => require("./publicAddressAdjust"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Eip/addressTransform:AddressTransform":
                return new AddressTransform(name, <any>undefined, { urn })
            case "tencentcloud:Eip/association:Association":
                return new Association(name, <any>undefined, { urn })
            case "tencentcloud:Eip/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "tencentcloud:Eip/normalAddressReturn:NormalAddressReturn":
                return new NormalAddressReturn(name, <any>undefined, { urn })
            case "tencentcloud:Eip/publicAddressAdjust:PublicAddressAdjust":
                return new PublicAddressAdjust(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Eip/addressTransform", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eip/association", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eip/instance", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eip/normalAddressReturn", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Eip/publicAddressAdjust", _module)
