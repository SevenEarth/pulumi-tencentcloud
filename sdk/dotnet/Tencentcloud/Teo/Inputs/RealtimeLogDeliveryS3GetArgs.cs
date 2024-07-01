// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Inputs
{

    public sealed class RealtimeLogDeliveryS3GetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fill in a custom SecretId to generate an encrypted signature. This parameter is required if the source site requires authentication.
        /// </summary>
        [Input("accessId", required: true)]
        public Input<string> AccessId { get; set; } = null!;

        [Input("accessKey", required: true)]
        private Input<string>? _accessKey;

        /// <summary>
        /// Fill in the custom SecretKey to generate the encrypted signature. This parameter is required if the source site requires authentication.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Bucket name and log storage directory, for example: `your_bucket_name/EO-logs/`. If this directory does not exist in the bucket, it will be created automatically.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Data compression type, the possible values are: `gzip`: use gzip compression. If it is not filled in, compression is not enabled.
        /// </summary>
        [Input("compressType")]
        public Input<string>? CompressType { get; set; }

        /// <summary>
        /// URLs that do not include bucket names or paths, for example: `https://storage.googleapis.com`, `https://s3.ap-northeast-2.amazonaws.com`, `https://cos.ap-nanjing.myqcloud.com`.
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// The region where the bucket is located, for example: ap-northeast-2.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        public RealtimeLogDeliveryS3GetArgs()
        {
        }
        public static new RealtimeLogDeliveryS3GetArgs Empty => new RealtimeLogDeliveryS3GetArgs();
    }
}
