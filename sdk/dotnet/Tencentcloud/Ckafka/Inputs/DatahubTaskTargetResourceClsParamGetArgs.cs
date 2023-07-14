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

    public sealed class DatahubTaskTargetResourceClsParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required when Decode Json is false.
        /// </summary>
        [Input("contentKey")]
        public Input<string>? ContentKey { get; set; }

        /// <summary>
        /// Whether the produced information is in json format.
        /// </summary>
        [Input("decodeJson", required: true)]
        public Input<bool> DecodeJson { get; set; } = null!;

        /// <summary>
        /// LogSet id.
        /// </summary>
        [Input("logSet")]
        public Input<string>? LogSet { get; set; }

        /// <summary>
        /// cls id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Specify the content of a field in the message as the time of the cls log. The format of the field content needs to be a second-level timestamp.
        /// </summary>
        [Input("timeField")]
        public Input<string>? TimeField { get; set; }

        public DatahubTaskTargetResourceClsParamGetArgs()
        {
        }
    }
}
