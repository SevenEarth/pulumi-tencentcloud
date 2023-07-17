// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Inputs
{

    public sealed class DatahubTaskTransformParamMapParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Type, DEFAULT default, DATE system default - timestamp, CUSTOMIZE custom, MAPPING mapping.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// value.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public DatahubTaskTransformParamMapParamGetArgs()
        {
        }
    }
}
