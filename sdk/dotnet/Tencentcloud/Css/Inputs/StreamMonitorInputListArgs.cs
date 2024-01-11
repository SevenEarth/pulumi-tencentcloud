// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Css.Inputs
{

    public sealed class StreamMonitorInputListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description content.limit 256 bytes.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Wait monitor input push path.limit 32 bytes.
        /// </summary>
        [Input("inputApp")]
        public Input<string>? InputApp { get; set; }

        /// <summary>
        /// Wait monitor input push domain.limit 128 bytes.
        /// </summary>
        [Input("inputDomain")]
        public Input<string>? InputDomain { get; set; }

        /// <summary>
        /// Wait monitor input stream name.limit 256 bytes.
        /// </summary>
        [Input("inputStreamName", required: true)]
        public Input<string> InputStreamName { get; set; } = null!;

        /// <summary>
        /// Wait monitor input stream push url.
        /// </summary>
        [Input("inputUrl")]
        public Input<string>? InputUrl { get; set; }

        public StreamMonitorInputListArgs()
        {
        }
    }
}