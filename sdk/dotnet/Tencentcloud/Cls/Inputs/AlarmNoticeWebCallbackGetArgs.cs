// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cls.Inputs
{

    public sealed class AlarmNoticeWebCallbackGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// abandoned.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// callback type, WeCom or Http.
        /// </summary>
        [Input("callbackType", required: true)]
        public Input<string> CallbackType { get; set; } = null!;

        [Input("headers")]
        private InputList<string>? _headers;

        /// <summary>
        /// abandoned.
        /// </summary>
        public InputList<string> Headers
        {
            get => _headers ?? (_headers = new InputList<string>());
            set => _headers = value;
        }

        /// <summary>
        /// index.
        /// </summary>
        [Input("index")]
        public Input<int>? Index { get; set; }

        /// <summary>
        /// Method, POST or PUT.
        /// </summary>
        [Input("method")]
        public Input<string>? Method { get; set; }

        /// <summary>
        /// callback url.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public AlarmNoticeWebCallbackGetArgs()
        {
        }
        public static new AlarmNoticeWebCallbackGetArgs Empty => new AlarmNoticeWebCallbackGetArgs();
    }
}
