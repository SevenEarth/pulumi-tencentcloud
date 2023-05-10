// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ci.Inputs
{

    public sealed class MediaTranscodeTemplateTransConfigHlsEncryptArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to enable HLS encryption, support encryption when Container.Format is hls.
        /// </summary>
        [Input("isHlsEncrypt")]
        public Input<string>? IsHlsEncrypt { get; set; }

        /// <summary>
        /// HLS encrypted key, this parameter is only meaningful when IsHlsEncrypt is true.
        /// </summary>
        [Input("uriKey")]
        public Input<string>? UriKey { get; set; }

        public MediaTranscodeTemplateTransConfigHlsEncryptArgs()
        {
        }
    }
}
