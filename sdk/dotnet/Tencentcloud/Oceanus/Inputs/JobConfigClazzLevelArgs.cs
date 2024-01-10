// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Oceanus.Inputs
{

    public sealed class JobConfigClazzLevelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Java class full pathNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("clazz", required: true)]
        public Input<string> Clazz { get; set; } = null!;

        /// <summary>
        /// Log level TRACE, DEBUG, INFO, WARN, ERRORNote: This field may return null, indicating that no valid value can be obtained.
        /// </summary>
        [Input("level", required: true)]
        public Input<string> Level { get; set; } = null!;

        public JobConfigClazzLevelArgs()
        {
        }
    }
}
