// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./certificate";
export * from "./freeCertificate";
export * from "./getCertificates";
export * from "./payCertificate";

// Import resources to register:
import { Certificate } from "./certificate";
import { FreeCertificate } from "./freeCertificate";
import { PayCertificate } from "./payCertificate";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Ssl/certificate:Certificate":
                return new Certificate(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/freeCertificate:FreeCertificate":
                return new FreeCertificate(name, <any>undefined, { urn })
            case "tencentcloud:Ssl/payCertificate:PayCertificate":
                return new PayCertificate(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/certificate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/freeCertificate", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Ssl/payCertificate", _module)
