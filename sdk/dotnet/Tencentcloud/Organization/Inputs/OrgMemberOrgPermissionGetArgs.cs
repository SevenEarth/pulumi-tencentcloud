// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Organization.Inputs
{

    public sealed class OrgMemberOrgPermissionGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Permissions ID.
        /// </summary>
        [Input("id")]
        public Input<int>? Id { get; set; }

        /// <summary>
        /// Member name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public OrgMemberOrgPermissionGetArgs()
        {
        }
    }
}
