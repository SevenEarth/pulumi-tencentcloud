// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tat.Inputs
{

    public sealed class InvokerScheduleSettingsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The next execution time of the invoker. This field is required if Policy is ONCE.
        /// </summary>
        [Input("invokeTime")]
        public Input<string>? InvokeTime { get; set; }

        /// <summary>
        /// Execution policy: `ONCE`: Execute once; `RECURRENCE`: Execute repeatedly.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        /// <summary>
        /// Trigger the crontab expression. This field is required if `Policy` is `RECURRENCE`. The crontab expression is parsed in UTC+8.
        /// </summary>
        [Input("recurrence")]
        public Input<string>? Recurrence { get; set; }

        public InvokerScheduleSettingsGetArgs()
        {
        }
    }
}
