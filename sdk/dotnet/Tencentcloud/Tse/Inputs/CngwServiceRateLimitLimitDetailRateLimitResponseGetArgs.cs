// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tse.Inputs
{

    public sealed class CngwServiceRateLimitLimitDetailRateLimitResponseGetArgs : Pulumi.ResourceArgs
    {
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputList<Inputs.CngwServiceRateLimitLimitDetailRateLimitResponseHeaderGetArgs>? _headers;
        public InputList<Inputs.CngwServiceRateLimitLimitDetailRateLimitResponseHeaderGetArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.CngwServiceRateLimitLimitDetailRateLimitResponseHeaderGetArgs>());
            set => _headers = value;
        }

        [Input("httpStatus")]
        public Input<int>? HttpStatus { get; set; }

        public CngwServiceRateLimitLimitDetailRateLimitResponseGetArgs()
        {
        }
    }
}