// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class GetDomainsDomainListRequestHeaderResult
    {
        /// <summary>
        /// Custom request header configuration rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDomainsDomainListRequestHeaderHeaderRuleResult> HeaderRules;
        /// <summary>
        /// Cache configuration switch.
        /// </summary>
        public readonly string Switch;

        [OutputConstructor]
        private GetDomainsDomainListRequestHeaderResult(
            ImmutableArray<Outputs.GetDomainsDomainListRequestHeaderHeaderRuleResult> headerRules,

            string @switch)
        {
            HeaderRules = headerRules;
            Switch = @switch;
        }
    }
}
