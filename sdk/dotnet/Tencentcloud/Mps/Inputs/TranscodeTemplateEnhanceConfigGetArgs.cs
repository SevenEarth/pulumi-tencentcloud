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

    public sealed class TranscodeTemplateEnhanceConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Video Enhancement Configuration.Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("videoEnhance")]
        public Input<Inputs.TranscodeTemplateEnhanceConfigVideoEnhanceGetArgs>? VideoEnhance { get; set; }

        public TranscodeTemplateEnhanceConfigGetArgs()
        {
        }
        public static new TranscodeTemplateEnhanceConfigGetArgs Empty => new TranscodeTemplateEnhanceConfigGetArgs();
    }
}
