// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor.Inputs
{

    public sealed class TmpTkeClusterAgentAgentsExternalLabelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicator name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Index value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TmpTkeClusterAgentAgentsExternalLabelArgs()
        {
        }
        public static new TmpTkeClusterAgentAgentsExternalLabelArgs Empty => new TmpTkeClusterAgentAgentsExternalLabelArgs();
    }
}
