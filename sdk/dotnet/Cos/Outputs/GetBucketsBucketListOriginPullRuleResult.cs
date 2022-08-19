// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cos.Outputs
{

    [OutputType]
    public sealed class GetBucketsBucketListOriginPullRuleResult
    {
        public readonly ImmutableDictionary<string, object>? CustomHttpHeaders;
        public readonly ImmutableArray<string> FollowHttpHeaders;
        public readonly bool? FollowQueryString;
        public readonly bool? FollowRedirection;
        public readonly string Host;
        public readonly string? Prefix;
        public readonly int Priority;
        public readonly string? Protocol;
        public readonly bool? SyncBackToSource;

        [OutputConstructor]
        private GetBucketsBucketListOriginPullRuleResult(
            ImmutableDictionary<string, object>? customHttpHeaders,

            ImmutableArray<string> followHttpHeaders,

            bool? followQueryString,

            bool? followRedirection,

            string host,

            string? prefix,

            int priority,

            string? protocol,

            bool? syncBackToSource)
        {
            CustomHttpHeaders = customHttpHeaders;
            FollowHttpHeaders = followHttpHeaders;
            FollowQueryString = followQueryString;
            FollowRedirection = followRedirection;
            Host = host;
            Prefix = prefix;
            Priority = priority;
            Protocol = protocol;
            SyncBackToSource = syncBackToSource;
        }
    }
}
