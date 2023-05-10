// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dts.Outputs
{

    [OutputType]
    public sealed class GetMigrateJobsListStepInfoStepInfoWarningResult
    {
        /// <summary>
        /// help document.
        /// </summary>
        public readonly string HelpDoc;
        /// <summary>
        /// message.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// solution.
        /// </summary>
        public readonly string Solution;

        [OutputConstructor]
        private GetMigrateJobsListStepInfoStepInfoWarningResult(
            string helpDoc,

            string message,

            string solution)
        {
            HelpDoc = helpDoc;
            Message = message;
            Solution = solution;
        }
    }
}
