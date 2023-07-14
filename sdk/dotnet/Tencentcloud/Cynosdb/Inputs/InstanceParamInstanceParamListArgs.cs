// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cynosdb.Inputs
{

    public sealed class InstanceParamInstanceParamListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current value of parameter.
        /// </summary>
        [Input("currentValue", required: true)]
        public Input<string> CurrentValue { get; set; } = null!;

        /// <summary>
        /// Parameter Name.
        /// </summary>
        [Input("paramName", required: true)]
        public Input<string> ParamName { get; set; } = null!;

        public InstanceParamInstanceParamListArgs()
        {
        }
    }
}
