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

    public sealed class JobPluginArgs : Pulumi.ResourceArgs
    {
        [Input("fileId")]
        public Input<string>? FileId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Scene Type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public JobPluginArgs()
        {
        }
    }
}
