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

    public sealed class DatahubTaskTransformsParamFieldChainSMTValueOperateJsonPathReplaceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Replacement value, Jsonpath expression or string.
        /// </summary>
        [Input("newValue", required: true)]
        public Input<string> NewValue { get; set; } = null!;

        /// <summary>
        /// Replaced value, Jsonpath expression.
        /// </summary>
        [Input("oldValue", required: true)]
        public Input<string> OldValue { get; set; } = null!;

        public DatahubTaskTransformsParamFieldChainSMTValueOperateJsonPathReplaceGetArgs()
        {
        }
    }
}
