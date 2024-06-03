// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Pts.Inputs
{

    public sealed class ScenarioProtocolGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// File ID.
        /// </summary>
        [Input("fileId")]
        public Input<string>? FileId { get; set; }

        /// <summary>
        /// Protocol name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// File name.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// File type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Update time.
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public ScenarioProtocolGetArgs()
        {
        }
        public static new ScenarioProtocolGetArgs Empty => new ScenarioProtocolGetArgs();
    }
}
