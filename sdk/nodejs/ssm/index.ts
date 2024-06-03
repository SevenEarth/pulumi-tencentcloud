// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { GetProductsArgs, GetProductsResult, GetProductsOutputArgs } from "./getProducts";
export const getProducts: typeof import("./getProducts").getProducts = null as any;
export const getProductsOutput: typeof import("./getProducts").getProductsOutput = null as any;
utilities.lazyLoad(exports, ["getProducts","getProductsOutput"], () => require("./getProducts"));

export { GetRotationDetailArgs, GetRotationDetailResult, GetRotationDetailOutputArgs } from "./getRotationDetail";
export const getRotationDetail: typeof import("./getRotationDetail").getRotationDetail = null as any;
export const getRotationDetailOutput: typeof import("./getRotationDetail").getRotationDetailOutput = null as any;
utilities.lazyLoad(exports, ["getRotationDetail","getRotationDetailOutput"], () => require("./getRotationDetail"));

export { GetRotationHistoryArgs, GetRotationHistoryResult, GetRotationHistoryOutputArgs } from "./getRotationHistory";
export const getRotationHistory: typeof import("./getRotationHistory").getRotationHistory = null as any;
export const getRotationHistoryOutput: typeof import("./getRotationHistory").getRotationHistoryOutput = null as any;
utilities.lazyLoad(exports, ["getRotationHistory","getRotationHistoryOutput"], () => require("./getRotationHistory"));

export { GetSecretVersionsArgs, GetSecretVersionsResult, GetSecretVersionsOutputArgs } from "./getSecretVersions";
export const getSecretVersions: typeof import("./getSecretVersions").getSecretVersions = null as any;
export const getSecretVersionsOutput: typeof import("./getSecretVersions").getSecretVersionsOutput = null as any;
utilities.lazyLoad(exports, ["getSecretVersions","getSecretVersionsOutput"], () => require("./getSecretVersions"));

export { GetSecretsArgs, GetSecretsResult, GetSecretsOutputArgs } from "./getSecrets";
export const getSecrets: typeof import("./getSecrets").getSecrets = null as any;
export const getSecretsOutput: typeof import("./getSecrets").getSecretsOutput = null as any;
utilities.lazyLoad(exports, ["getSecrets","getSecretsOutput"], () => require("./getSecrets"));

export { GetServiceStatusArgs, GetServiceStatusResult, GetServiceStatusOutputArgs } from "./getServiceStatus";
export const getServiceStatus: typeof import("./getServiceStatus").getServiceStatus = null as any;
export const getServiceStatusOutput: typeof import("./getServiceStatus").getServiceStatusOutput = null as any;
utilities.lazyLoad(exports, ["getServiceStatus","getServiceStatusOutput"], () => require("./getServiceStatus"));

export { GetSshKeyPairValueArgs, GetSshKeyPairValueResult, GetSshKeyPairValueOutputArgs } from "./getSshKeyPairValue";
export const getSshKeyPairValue: typeof import("./getSshKeyPairValue").getSshKeyPairValue = null as any;
export const getSshKeyPairValueOutput: typeof import("./getSshKeyPairValue").getSshKeyPairValueOutput = null as any;
utilities.lazyLoad(exports, ["getSshKeyPairValue","getSshKeyPairValueOutput"], () => require("./getSshKeyPairValue"));

export { ProductSecretArgs, ProductSecretState } from "./productSecret";
export type ProductSecret = import("./productSecret").ProductSecret;
export const ProductSecret: typeof import("./productSecret").ProductSecret = null as any;
utilities.lazyLoad(exports, ["ProductSecret"], () => require("./productSecret"));

export { RotateProductSecretArgs, RotateProductSecretState } from "./rotateProductSecret";
export type RotateProductSecret = import("./rotateProductSecret").RotateProductSecret;
export const RotateProductSecret: typeof import("./rotateProductSecret").RotateProductSecret = null as any;
utilities.lazyLoad(exports, ["RotateProductSecret"], () => require("./rotateProductSecret"));

export { SecretArgs, SecretState } from "./secret";
export type Secret = import("./secret").Secret;
export const Secret: typeof import("./secret").Secret = null as any;
utilities.lazyLoad(exports, ["Secret"], () => require("./secret"));

export { SecretVersionArgs, SecretVersionState } from "./secretVersion";
export type SecretVersion = import("./secretVersion").SecretVersion;
export const SecretVersion: typeof import("./secretVersion").SecretVersion = null as any;
utilities.lazyLoad(exports, ["SecretVersion"], () => require("./secretVersion"));

export { SshKeyPairSecretArgs, SshKeyPairSecretState } from "./sshKeyPairSecret";
export type SshKeyPairSecret = import("./sshKeyPairSecret").SshKeyPairSecret;
export const SshKeyPairSecret: typeof import("./sshKeyPairSecret").SshKeyPairSecret = null as any;
utilities.lazyLoad(exports, ["SshKeyPairSecret"], () => require("./sshKeyPairSecret"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Ssm/productSecret:ProductSecret":
                return new ProductSecret(name, <any>undefined, { urn })
            case "tencentcloud:Ssm/rotateProductSecret:RotateProductSecret":
                return new RotateProductSecret(name, <any>undefined, { urn })
            case "tencentcloud:Ssm/secret:Secret":
                return new Secret(name, <any>undefined, { urn })
            case "tencentcloud:Ssm/secretVersion:SecretVersion":
                return new SecretVersion(name, <any>undefined, { urn })
            case "tencentcloud:Ssm/sshKeyPairSecret:SshKeyPairSecret":
                return new SshKeyPairSecret(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/productSecret", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/rotateProductSecret", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/secret", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/secretVersion", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssm/sshKeyPairSecret", _module)
