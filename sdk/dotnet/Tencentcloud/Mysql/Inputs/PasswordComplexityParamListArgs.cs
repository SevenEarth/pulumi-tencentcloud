// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Mysql.Inputs
{

    public sealed class PasswordComplexityParamListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter value.
        /// </summary>
        [Input("currentValue")]
        public Input<string>? CurrentValue { get; set; }

        /// <summary>
        /// Parameter name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public PasswordComplexityParamListArgs()
        {
        }
        public static new PasswordComplexityParamListArgs Empty => new PasswordComplexityParamListArgs();
    }
}
