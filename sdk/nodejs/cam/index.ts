// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getGroupMemberships";
export * from "./getGroupPolicyAttachments";
export * from "./getGroups";
export * from "./getPolicies";
export * from "./getRolePolicyAttachments";
export * from "./getRoles";
export * from "./getSamlProviders";
export * from "./getUserPolicyAttachments";
export * from "./getUsers";
export * from "./group";
export * from "./groupMembership";
export * from "./groupPolicyAttachment";
export * from "./oidcSso";
export * from "./policy";
export * from "./role";
export * from "./rolePolicyAttachment";
export * from "./roleSso";
export * from "./samlProvider";
export * from "./user";
export * from "./userPolicyAttachment";

// Import resources to register:
import { Group } from "./group";
import { GroupMembership } from "./groupMembership";
import { GroupPolicyAttachment } from "./groupPolicyAttachment";
import { OidcSso } from "./oidcSso";
import { Policy } from "./policy";
import { Role } from "./role";
import { RolePolicyAttachment } from "./rolePolicyAttachment";
import { RoleSso } from "./roleSso";
import { SamlProvider } from "./samlProvider";
import { User } from "./user";
import { UserPolicyAttachment } from "./userPolicyAttachment";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Cam/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "tencentcloud:Cam/groupMembership:GroupMembership":
                return new GroupMembership(name, <any>undefined, { urn })
            case "tencentcloud:Cam/groupPolicyAttachment:GroupPolicyAttachment":
                return new GroupPolicyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cam/oidcSso:OidcSso":
                return new OidcSso(name, <any>undefined, { urn })
            case "tencentcloud:Cam/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "tencentcloud:Cam/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "tencentcloud:Cam/rolePolicyAttachment:RolePolicyAttachment":
                return new RolePolicyAttachment(name, <any>undefined, { urn })
            case "tencentcloud:Cam/roleSso:RoleSso":
                return new RoleSso(name, <any>undefined, { urn })
            case "tencentcloud:Cam/samlProvider:SamlProvider":
                return new SamlProvider(name, <any>undefined, { urn })
            case "tencentcloud:Cam/user:User":
                return new User(name, <any>undefined, { urn })
            case "tencentcloud:Cam/userPolicyAttachment:UserPolicyAttachment":
                return new UserPolicyAttachment(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/group", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/groupMembership", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/groupPolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/oidcSso", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/policy", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/role", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/rolePolicyAttachment", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/roleSso", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/samlProvider", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/user", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Cam/userPolicyAttachment", _module)
