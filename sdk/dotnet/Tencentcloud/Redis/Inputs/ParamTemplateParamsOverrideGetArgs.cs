// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis.Inputs
{

    public sealed class ParamTemplateParamsOverrideGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter key e.g. `timeout`, check https://www.tencentcloud.com/document/product/239/39796 for more reference.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// Parameter value, check https://www.tencentcloud.com/document/product/239/39796 for more reference.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ParamTemplateParamsOverrideGetArgs()
        {
        }
    }
}