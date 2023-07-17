// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Ckafka.Outputs
{

    [OutputType]
    public sealed class GetTopicSubscribeGroupGroupsInfoResult
    {
        /// <summary>
        /// Error code, normally 0.
        /// </summary>
        public readonly string ErrorCode;
        /// <summary>
        /// Kafka consumer group.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// This array contains information only if state is Stable and protocol_type is consumer.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetTopicSubscribeGroupGroupsInfoMemberResult> Members;
        /// <summary>
        /// Common consumer partition allocation algorithms are as follows (the default option for Kafka consumer SDK is range) range|roundrobin| sticky.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The protocol type selected by the consumption group is normally the consumer, but some systems use their own protocol, such as kafka-connect, which uses connect. Only the standard consumer protocol, this interface knows the format of the specific allocation method, and can analyze the specific partition allocation.
        /// </summary>
        public readonly string ProtocolType;
        /// <summary>
        /// Group state description (commonly Empty, Stable, and Dead states): Dead: The consumption group does not exist Empty: The consumption group does not currently have any consumer subscriptions PreparingRebalance: The consumption group is in the rebalance state CompletingRebalance: The consumption group is in the rebalance state Stable: Each consumer in the consumption group has joined and is in a stable state.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetTopicSubscribeGroupGroupsInfoResult(
            string errorCode,

            string group,

            ImmutableArray<Outputs.GetTopicSubscribeGroupGroupsInfoMemberResult> members,

            string protocol,

            string protocolType,

            string state)
        {
            ErrorCode = errorCode;
            Group = group;
            Members = members;
            Protocol = protocol;
            ProtocolType = protocolType;
            State = state;
        }
    }
}
