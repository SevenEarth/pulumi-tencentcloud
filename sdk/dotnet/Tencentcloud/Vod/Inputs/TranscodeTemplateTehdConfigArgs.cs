// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Vod.Inputs
{

    public sealed class TranscodeTemplateTehdConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum bitrate, which is valid when `Type` is `TESHD`.If this parameter is left blank or 0 is entered, there will be no upper limit for bitrate.
        /// </summary>
        [Input("maxVideoBitrate")]
        public Input<int>? MaxVideoBitrate { get; set; }

        /// <summary>
        /// TESHD transcoding type. Valid values: TEHD-100, OFF (default).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public TranscodeTemplateTehdConfigArgs()
        {
        }
        public static new TranscodeTemplateTehdConfigArgs Empty => new TranscodeTemplateTehdConfigArgs();
    }
}
