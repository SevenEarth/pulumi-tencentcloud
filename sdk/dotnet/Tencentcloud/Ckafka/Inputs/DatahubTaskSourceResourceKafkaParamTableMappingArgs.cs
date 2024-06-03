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

    public sealed class DatahubTaskSourceResourceKafkaParamTableMappingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// database name.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// table name,use, to separate.
        /// </summary>
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        /// <summary>
        /// Topic name.
        /// </summary>
        [Input("topic", required: true)]
        public Input<string> Topic { get; set; } = null!;

        /// <summary>
        /// Topic ID.
        /// </summary>
        [Input("topicId", required: true)]
        public Input<string> TopicId { get; set; } = null!;

        public DatahubTaskSourceResourceKafkaParamTableMappingArgs()
        {
        }
        public static new DatahubTaskSourceResourceKafkaParamTableMappingArgs Empty => new DatahubTaskSourceResourceKafkaParamTableMappingArgs();
    }
}
