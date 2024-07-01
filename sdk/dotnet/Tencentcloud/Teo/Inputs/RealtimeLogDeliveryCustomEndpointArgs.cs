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

    public sealed class RealtimeLogDeliveryCustomEndpointArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fill in a custom SecretId to generate an encrypted signature. This parameter is required if the source site requires authentication.
        /// </summary>
        [Input("accessId")]
        public Input<string>? AccessId { get; set; }

        /// <summary>
        /// Fill in the custom SecretKey to generate the encrypted signature. This parameter is required if the source site requires authentication.
        /// </summary>
        [Input("accessKey")]
        public Input<string>? AccessKey { get; set; }

        /// <summary>
        /// Data compression type, the possible values are: `gzip`: use gzip compression. If it is not filled in, compression is not enabled.
        /// </summary>
        [Input("compressType")]
        public Input<string>? CompressType { get; set; }

        [Input("headers")]
        private InputList<Inputs.RealtimeLogDeliveryCustomEndpointHeaderArgs>? _headers;

        /// <summary>
        /// The custom request header carried when delivering logs. If the header name you fill in is the default header carried by EdgeOne log push, such as Content-Type, then the header value you fill in will overwrite the default value. The header value references a single variable ${batchSize} to obtain the number of logs included in each POST request.
        /// </summary>
        public InputList<Inputs.RealtimeLogDeliveryCustomEndpointHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.RealtimeLogDeliveryCustomEndpointHeaderArgs>());
            set => _headers = value;
        }

        /// <summary>
        /// When sending logs via POST request, the application layer protocol type used can be: `http`: HTTP protocol; `https`: HTTPS protocol. If not filled in, the protocol type will be parsed according to the filled in URL address.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The custom HTTP interface address for real-time log delivery. Currently, only HTTP/HTTPS protocols are supported.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public RealtimeLogDeliveryCustomEndpointArgs()
        {
        }
        public static new RealtimeLogDeliveryCustomEndpointArgs Empty => new RealtimeLogDeliveryCustomEndpointArgs();
    }
}
