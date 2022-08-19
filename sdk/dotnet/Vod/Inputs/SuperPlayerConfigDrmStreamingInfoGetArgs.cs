// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Vod.Inputs
{

    public sealed class SuperPlayerConfigDrmStreamingInfoGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the adaptive dynamic streaming template whose protection type is `SimpleAES`.
        /// </summary>
        [Input("simpleAesDefinition")]
        public Input<string>? SimpleAesDefinition { get; set; }

        public SuperPlayerConfigDrmStreamingInfoGetArgs()
        {
        }
    }
}
