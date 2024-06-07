// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Eb.Inputs
{

    public sealed class EventConnectorConnectionDescriptionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// apigw parameter,Note: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("apiGwParams")]
        public Input<Inputs.EventConnectorConnectionDescriptionApiGwParamsArgs>? ApiGwParams { get; set; }

        /// <summary>
        /// ckafka parameter, note: this field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("ckafkaParams")]
        public Input<Inputs.EventConnectorConnectionDescriptionCkafkaParamsArgs>? CkafkaParams { get; set; }

        /// <summary>
        /// Resource qcs six-segment style, more reference [resource six-segment style](https://cloud.tencent.com/document/product/598/10606).
        /// </summary>
        [Input("resourceDescription", required: true)]
        public Input<string> ResourceDescription { get; set; } = null!;

        public EventConnectorConnectionDescriptionArgs()
        {
        }
        public static new EventConnectorConnectionDescriptionArgs Empty => new EventConnectorConnectionDescriptionArgs();
    }
}
