// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tcr.Outputs
{

    [OutputType]
    public sealed class InstanceSecurityPolicy
    {
        /// <summary>
        /// The public network IP address of the access source.
        /// </summary>
        public readonly string? CidrBlock;
        /// <summary>
        /// Remarks of policy.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Index of policy.
        /// </summary>
        public readonly int? Index;
        /// <summary>
        /// Version of policy.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private InstanceSecurityPolicy(
            string? cidrBlock,

            string? description,

            int? index,

            string? version)
        {
            CidrBlock = cidrBlock;
            Description = description;
            Index = index;
            Version = version;
        }
    }
}
