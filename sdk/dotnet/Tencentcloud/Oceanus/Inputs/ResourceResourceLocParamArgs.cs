// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Oceanus.Inputs
{

    public sealed class ResourceResourceLocParamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource bucket.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Resource path.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Resource region, if not set, use resource region, note: this field may return null, indicating that no valid values can be obtained.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public ResourceResourceLocParamArgs()
        {
        }
        public static new ResourceResourceLocParamArgs Empty => new ResourceResourceLocParamArgs();
    }
}
