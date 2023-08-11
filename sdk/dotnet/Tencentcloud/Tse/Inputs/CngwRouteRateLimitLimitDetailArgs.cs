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

    public sealed class CngwRouteRateLimitLimitDetailArgs : Pulumi.ResourceArgs
    {
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("externalRedis")]
        public Input<Inputs.CngwRouteRateLimitLimitDetailExternalRedisArgs>? ExternalRedis { get; set; }

        [Input("header")]
        public Input<string>? Header { get; set; }

        [Input("hideClientHeaders", required: true)]
        public Input<bool> HideClientHeaders { get; set; } = null!;

        [Input("isDelay", required: true)]
        public Input<bool> IsDelay { get; set; } = null!;

        [Input("limitBy", required: true)]
        public Input<string> LimitBy { get; set; } = null!;

        [Input("lineUpTime")]
        public Input<int>? LineUpTime { get; set; }

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("qpsThresholds", required: true)]
        private InputList<Inputs.CngwRouteRateLimitLimitDetailQpsThresholdArgs>? _qpsThresholds;
        public InputList<Inputs.CngwRouteRateLimitLimitDetailQpsThresholdArgs> QpsThresholds
        {
            get => _qpsThresholds ?? (_qpsThresholds = new InputList<Inputs.CngwRouteRateLimitLimitDetailQpsThresholdArgs>());
            set => _qpsThresholds = value;
        }

        [Input("rateLimitResponse")]
        public Input<Inputs.CngwRouteRateLimitLimitDetailRateLimitResponseArgs>? RateLimitResponse { get; set; }

        [Input("rateLimitResponseUrl")]
        public Input<string>? RateLimitResponseUrl { get; set; }

        [Input("responseType", required: true)]
        public Input<string> ResponseType { get; set; } = null!;

        public CngwRouteRateLimitLimitDetailArgs()
        {
        }
    }
}