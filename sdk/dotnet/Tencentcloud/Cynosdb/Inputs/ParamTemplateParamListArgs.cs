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

    public sealed class ParamTemplateParamListArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Current value.
        /// </summary>
        [Input("currentValue")]
        public Input<string>? CurrentValue { get; set; }

        /// <summary>
        /// Parameter Name.
        /// </summary>
        [Input("paramName")]
        public Input<string>? ParamName { get; set; }

        public ParamTemplateParamListArgs()
        {
        }
    }
}
