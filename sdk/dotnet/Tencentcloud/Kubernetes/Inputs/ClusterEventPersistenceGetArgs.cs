// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Inputs
{

    public sealed class ClusterEventPersistenceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specify weather the Event Persistence enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Specify id of existing CLS log set, or auto create a new set by leave it empty.
        /// </summary>
        [Input("logSetId")]
        public Input<string>? LogSetId { get; set; }

        /// <summary>
        /// Specify id of existing CLS log topic, or auto create a new topic by leave it empty.
        /// </summary>
        [Input("topicId")]
        public Input<string>? TopicId { get; set; }

        public ClusterEventPersistenceGetArgs()
        {
        }
    }
}