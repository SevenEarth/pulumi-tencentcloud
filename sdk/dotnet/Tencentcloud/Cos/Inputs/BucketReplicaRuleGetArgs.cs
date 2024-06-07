// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cos.Inputs
{

    public sealed class BucketReplicaRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination bucket identifier, format: `qcs::cos:&lt;region&gt;::&lt;bucketname-appid&gt;`. NOTE: destination bucket must enable versioning.
        /// </summary>
        [Input("destinationBucket", required: true)]
        public Input<string> DestinationBucket { get; set; } = null!;

        /// <summary>
        /// Storage class of destination, available values: `STANDARD`, `INTELLIGENT_TIERING`, `STANDARD_IA`. default is following current class of destination.
        /// </summary>
        [Input("destinationStorageClass")]
        public Input<string>? DestinationStorageClass { get; set; }

        /// <summary>
        /// Name of a specific rule.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Prefix matching policy. Policies cannot overlap; otherwise, an error will be returned. To match the root directory, leave this parameter empty.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        /// <summary>
        /// Status identifier, available values: `Enabled`, `Disabled`.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public BucketReplicaRuleGetArgs()
        {
        }
        public static new BucketReplicaRuleGetArgs Empty => new BucketReplicaRuleGetArgs();
    }
}
