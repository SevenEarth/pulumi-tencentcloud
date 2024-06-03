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

    public sealed class BucketInventoryDestinationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the bucket owner.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// Bucket name.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Server-side encryption for the inventory result.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.BucketInventoryDestinationEncryptionGetArgs>? Encryption { get; set; }

        /// <summary>
        /// Format of the inventory result. Valid value: CSV.
        /// </summary>
        [Input("format", required: true)]
        public Input<string> Format { get; set; } = null!;

        /// <summary>
        /// Prefix of the inventory result.
        /// </summary>
        [Input("prefix")]
        public Input<string>? Prefix { get; set; }

        public BucketInventoryDestinationGetArgs()
        {
        }
        public static new BucketInventoryDestinationGetArgs Empty => new BucketInventoryDestinationGetArgs();
    }
}
