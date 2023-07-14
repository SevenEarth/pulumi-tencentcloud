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

    public sealed class DatahubTaskTargetResourceDtsParamGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dts consumer group Id.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// Dts consumer group passwd.
        /// </summary>
        [Input("groupPassword")]
        public Input<string>? GroupPassword { get; set; }

        /// <summary>
        /// Dts account.
        /// </summary>
        [Input("groupUser")]
        public Input<string>? GroupUser { get; set; }

        /// <summary>
        /// Dts connection ip.
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Dts connection port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Dts instance Id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// Dts topic.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        /// <summary>
        /// False to synchronize the original data, true to synchronize the parsed json format data, the default is true.
        /// </summary>
        [Input("tranSql")]
        public Input<bool>? TranSql { get; set; }

        public DatahubTaskTargetResourceDtsParamGetArgs()
        {
        }
    }
}
