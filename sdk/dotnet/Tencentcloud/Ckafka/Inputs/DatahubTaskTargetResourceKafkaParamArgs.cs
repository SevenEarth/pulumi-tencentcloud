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

    public sealed class DatahubTaskTargetResourceKafkaParamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to compress when writing to the Topic, if it is not enabled, fill in none, if it is enabled, fill in open.
        /// </summary>
        [Input("compressionType")]
        public Input<string>? CompressionType { get; set; }

        /// <summary>
        /// Enable the fault-tolerant instance and enable the dead-letter queue.
        /// </summary>
        [Input("enableToleration")]
        public Input<bool>? EnableToleration { get; set; }

        /// <summary>
        /// 1 source topic message is amplified into msg Multiple and written to the target topic (this parameter is currently only applicable to ckafka flowing into ckafka).
        /// </summary>
        [Input("msgMultiple")]
        public Input<int>? MsgMultiple { get; set; }

        /// <summary>
        /// Offset type, initial position earliest, latest position latest, time point position timestamp.
        /// </summary>
        [Input("offsetType")]
        public Input<string>? OffsetType { get; set; }

        /// <summary>
        /// Partition num.
        /// </summary>
        [Input("partitionNum")]
        public Input<int>? PartitionNum { get; set; }

        /// <summary>
        /// Qps limit.
        /// </summary>
        [Input("qpsLimit")]
        public Input<int>? QpsLimit { get; set; }

        /// <summary>
        /// resource id.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// resource id name.
        /// </summary>
        [Input("resourceName")]
        public Input<string>? ResourceName { get; set; }

        /// <summary>
        /// Whether it is a self-built cluster.
        /// </summary>
        [Input("selfBuilt", required: true)]
        public Input<bool> SelfBuilt { get; set; } = null!;

        /// <summary>
        /// It must be passed when the Offset type is timestamp, and the time stamp is passed, accurate to the second.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        [Input("tableMappings")]
        private InputList<Inputs.DatahubTaskTargetResourceKafkaParamTableMappingArgs>? _tableMappings;

        /// <summary>
        /// The route from Table to Topic must be passed when the Distribute to multiple topics switch is turned on.
        /// </summary>
        public InputList<Inputs.DatahubTaskTargetResourceKafkaParamTableMappingArgs> TableMappings
        {
            get => _tableMappings ?? (_tableMappings = new InputList<Inputs.DatahubTaskTargetResourceKafkaParamTableMappingArgs>());
            set => _tableMappings = value;
        }

        /// <summary>
        /// Topic name, multiple separated by,.
        /// </summary>
        [Input("topic")]
        public Input<string>? Topic { get; set; }

        /// <summary>
        /// Topic Id.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        /// <summary>
        /// whether the used topic need to be automatically created (currently only supports SOURCE inflow tasks, if you do not use to distribute to multiple topics, you need to fill in the topic name that needs to be automatically created in the Topic field).
        /// </summary>
        [Input("useAutoCreateTopic")]
        public Input<bool>? UseAutoCreateTopic { get; set; }

        /// <summary>
        /// Distribute to multiple topics switch, the default is false.
        /// </summary>
        [Input("useTableMapping")]
        public Input<bool>? UseTableMapping { get; set; }

        /// <summary>
        /// Zone ID.
        /// </summary>
        [Input("zoneId")]
        public Input<int>? ZoneId { get; set; }

        public DatahubTaskTargetResourceKafkaParamArgs()
        {
        }
    }
}
