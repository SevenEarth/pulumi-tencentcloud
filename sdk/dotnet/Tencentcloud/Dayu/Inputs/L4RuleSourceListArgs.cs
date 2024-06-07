// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dayu.Inputs
{

    public sealed class L4RuleSourceListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Source IP or domain, valid format of ip is like `1.1.1.1` and valid format of host source is like `abc.com`.
        /// </summary>
        [Input("source", required: true)]
        public Input<string> Source { get; set; } = null!;

        /// <summary>
        /// Weight of the source, the valid value ranges from 0 to 100.
        /// </summary>
        [Input("weight", required: true)]
        public Input<int> Weight { get; set; } = null!;

        public L4RuleSourceListArgs()
        {
        }
        public static new L4RuleSourceListArgs Empty => new L4RuleSourceListArgs();
    }
}
