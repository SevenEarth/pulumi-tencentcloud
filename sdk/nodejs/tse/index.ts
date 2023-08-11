// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./cngwCanaryRule";
export * from "./cngwRoute";
export * from "./cngwRouteRateLimit";
export * from "./cngwService";
export * from "./cngwServiceRateLimit";
export * from "./getAccessAddress";
export * from "./getGatewayCanaryRules";
export * from "./getGatewayNodes";
export * from "./getGatewayRoutes";
export * from "./getGatewayServices";
export * from "./getNacosReplicas";
export * from "./getNacosServerInterfaces";
export * from "./getZookeeperReplicas";
export * from "./getZookeeperServerInterfaces";
export * from "./instance";

// Import resources to register:
import { CngwCanaryRule } from "./cngwCanaryRule";
import { CngwRoute } from "./cngwRoute";
import { CngwRouteRateLimit } from "./cngwRouteRateLimit";
import { CngwService } from "./cngwService";
import { CngwServiceRateLimit } from "./cngwServiceRateLimit";
import { Instance } from "./instance";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Tse/cngwCanaryRule:CngwCanaryRule":
                return new CngwCanaryRule(name, <any>undefined, { urn })
            case "tencentcloud:Tse/cngwRoute:CngwRoute":
                return new CngwRoute(name, <any>undefined, { urn })
            case "tencentcloud:Tse/cngwRouteRateLimit:CngwRouteRateLimit":
                return new CngwRouteRateLimit(name, <any>undefined, { urn })
            case "tencentcloud:Tse/cngwService:CngwService":
                return new CngwService(name, <any>undefined, { urn })
            case "tencentcloud:Tse/cngwServiceRateLimit:CngwServiceRateLimit":
                return new CngwServiceRateLimit(name, <any>undefined, { urn })
            case "tencentcloud:Tse/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Tse/cngwCanaryRule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tse/cngwRoute", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tse/cngwRouteRateLimit", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tse/cngwService", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tse/cngwServiceRateLimit", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tse/instance", _module)