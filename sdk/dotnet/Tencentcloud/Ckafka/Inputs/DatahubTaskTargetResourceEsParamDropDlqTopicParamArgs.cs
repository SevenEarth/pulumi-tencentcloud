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

    public sealed class DatahubTaskTargetResourceEsParamDropDlqTopicParamArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to perform compression when writing a topic, if it is not enabled, fill in none, if it is enabled, you can choose one of gzip, snappy, lz4 to fill in.
        /// </summary>
        [Input("compressionType")]
        public Input<string>? CompressionType { get; set; }

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
        /// The topic name of the topic sold separately.
        /// </summary>
        [Input("resource", required: true)]
        public Input<string> Resource { get; set; } = null!;

        /// <summary>
        /// It must be passed when the Offset type is timestamp, and the time stamp is passed, accurate to the second.
        /// </summary>
        [Input("startTime")]
        public Input<int>? StartTime { get; set; }

        /// <summary>
        /// TopicId.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        /// <summary>
        /// whether the used topic need to be automatically created (currently only supports SOURCE inflow tasks).
        /// </summary>
        [Input("useAutoCreateTopic")]
        public Input<bool>? UseAutoCreateTopic { get; set; }

        public DatahubTaskTargetResourceEsParamDropDlqTopicParamArgs()
        {
        }
    }
}
