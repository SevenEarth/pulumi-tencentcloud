// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./ccHttpPolicy";
export * from "./ccHttpsPolicy";
export * from "./ccPolicyV2";
export * from "./ddosPolicy";
export * from "./ddosPolicyAttachment";
export * from "./ddosPolicyCase";
export * from "./ddosPolicyV2";
export * from "./eip";
export * from "./getCcHttpPolicies";
export * from "./getCcHttpsPolicies";
export * from "./getDdosPolicies";
export * from "./getDdosPolicyAttachments";
export * from "./getDdosPolicyCases";
export * from "./getEip";
export * from "./getL4Rules";
export * from "./getL4RulesV2";
export * from "./getL7Rules";
export * from "./getL7RulesV2";
export * from "./l4rule";
export * from "./l4ruleV2";
export * from "./l7rule";
export * from "./l7ruleV2";

// Import resources to register:
import { CcHttpPolicy } from "./ccHttpPolicy";
import { CcHttpsPolicy } from "./ccHttpsPolicy";
import { CcPolicyV2 } from "./ccPolicyV2";
import { DdosPolicy } from "./ddosPolicy";
import { DdosPolicyAttachment } from "./ddosPolicyAttachment";
import { DdosPolicyCase } from "./ddosPolicyCase";
import { DdosPolicyV2 } from "./ddosPolicyV2";
import { Eip } from "./eip";
import { L4Rule } from "./l4rule";
import { L4RuleV2 } from "./l4ruleV2";
import { L7Rule } from "./l7rule";
import { L7RuleV2 } from "./l7ruleV2";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Dayu/ccHttpPolicy:CcHttpPolicy":
                return new CcHttpPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ccHttpsPolicy:CcHttpsPolicy":
                return new CcHttpsPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ccPolicyV2:CcPolicyV2":
                return new CcPolicyV2(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicy:DdosPolicy":
                return new DdosPolicy(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicyAttachment:DdosPolicyAttachment":
                return new DdosPolicyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicyCase:DdosPolicyCase":
                return new DdosPolicyCase(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/ddosPolicyV2:DdosPolicyV2":
                return new DdosPolicyV2(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/eip:Eip":
                return new Eip(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l4Rule:L4Rule":
                return new L4Rule(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l4RuleV2:L4RuleV2":
                return new L4RuleV2(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l7Rule:L7Rule":
                return new L7Rule(name, <any>undefined, { urn })
            case "tencentcloud:Dayu/l7RuleV2:L7RuleV2":
                return new L7RuleV2(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ccHttpPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ccHttpsPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ccPolicyV2", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicyCase", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/ddosPolicyV2", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/eip", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l4Rule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l4RuleV2", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l7Rule", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Dayu/l7RuleV2", _module)
