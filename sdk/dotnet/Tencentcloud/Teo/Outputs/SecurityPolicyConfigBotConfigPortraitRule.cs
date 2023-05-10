// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Teo.Outputs
{

    [OutputType]
    public sealed class SecurityPolicyConfigBotConfigPortraitRule
    {
        /// <summary>
        /// Rules to enable when action is `alg`. See details in data source `bot_portrait_rules`.
        /// </summary>
        public readonly ImmutableArray<int> AlgManagedIds;
        /// <summary>
        /// Rules to enable when action is `captcha`. See details in data source `bot_portrait_rules`.
        /// </summary>
        public readonly ImmutableArray<int> CapManagedIds;
        /// <summary>
        /// Rules to enable when action is `drop`. See details in data source `bot_portrait_rules`.
        /// </summary>
        public readonly ImmutableArray<int> DropManagedIds;
        /// <summary>
        /// Rules to enable when action is `monitor`. See details in data source `bot_portrait_rules`.
        /// </summary>
        public readonly ImmutableArray<int> MonManagedIds;
        public readonly int? RuleId;
        /// <summary>
        /// - `on`: Enable.- `off`: Disable.
        /// </summary>
        public readonly string? Switch;

        [OutputConstructor]
        private SecurityPolicyConfigBotConfigPortraitRule(
            ImmutableArray<int> algManagedIds,

            ImmutableArray<int> capManagedIds,

            ImmutableArray<int> dropManagedIds,

            ImmutableArray<int> monManagedIds,

            int? ruleId,

            string? @switch)
        {
            AlgManagedIds = algManagedIds;
            CapManagedIds = capManagedIds;
            DropManagedIds = dropManagedIds;
            MonManagedIds = monManagedIds;
            RuleId = ruleId;
            Switch = @switch;
        }
    }
}
