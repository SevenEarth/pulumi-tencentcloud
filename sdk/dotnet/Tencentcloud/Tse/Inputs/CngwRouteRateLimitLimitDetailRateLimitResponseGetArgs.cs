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

    public sealed class CngwRouteRateLimitLimitDetailRateLimitResponseGetArgs : Pulumi.ResourceArgs
    {
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputList<Inputs.CngwRouteRateLimitLimitDetailRateLimitResponseHeaderGetArgs>? _headers;
        public InputList<Inputs.CngwRouteRateLimitLimitDetailRateLimitResponseHeaderGetArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.CngwRouteRateLimitLimitDetailRateLimitResponseHeaderGetArgs>());
            set => _headers = value;
        }

        [Input("httpStatus")]
        public Input<int>? HttpStatus { get; set; }

        public CngwRouteRateLimitLimitDetailRateLimitResponseGetArgs()
        {
        }
    }
}
