// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetApiAppApiResultBase64EncodedTriggerRuleResult
    {
        /// <summary>
        /// The backend service parameter name of the API. This parameter will be used only if the ServiceType is HTTP. The front-end and back-end parameter names can be different.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The value of the note.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private GetApiAppApiResultBase64EncodedTriggerRuleResult(
            string name,

            ImmutableArray<string> values)
        {
            Name = name;
            Values = values;
        }
    }
}
