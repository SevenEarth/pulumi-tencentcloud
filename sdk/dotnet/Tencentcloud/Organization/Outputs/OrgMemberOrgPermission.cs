// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Organization.Outputs
{

    [OutputType]
    public sealed class OrgMemberOrgPermission
    {
        /// <summary>
        /// Permissions ID.
        /// </summary>
        public readonly int? Id;
        /// <summary>
        /// Member name.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private OrgMemberOrgPermission(
            int? id,

            string? name)
        {
            Id = id;
            Name = name;
        }
    }
}
