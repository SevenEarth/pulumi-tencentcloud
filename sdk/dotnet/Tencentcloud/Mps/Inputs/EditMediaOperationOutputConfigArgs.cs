// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mps.Inputs
{

    public sealed class EditMediaOperationOutputConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Format. Valid values: `mp4` (default), `hls`, `mov`, `flv`, `avi`.
        /// </summary>
        [Input("container")]
        public Input<string>? Container { get; set; }

        /// <summary>
        /// The editing mode. Valid values are `normal` and `fast`. The default is `normal`, which indicates precise editing.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public EditMediaOperationOutputConfigArgs()
        {
        }
    }
}
