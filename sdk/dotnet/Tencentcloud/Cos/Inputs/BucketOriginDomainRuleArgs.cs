// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketOriginDomainRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify domain host.
        /// </summary>
        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        /// <summary>
        /// Domain status, default: `ENABLED`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Specify origin domain type, available values: `REST`, `WEBSITE`, `ACCELERATE`, default: `REST`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public BucketOriginDomainRuleArgs()
        {
        }
        public static new BucketOriginDomainRuleArgs Empty => new BucketOriginDomainRuleArgs();
    }
}
