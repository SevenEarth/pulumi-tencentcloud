// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dcdb.Inputs
{

    public sealed class DbInstanceInitParamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of parameter.
        /// </summary>
        [Input("param", required: true)]
        public Input<string> Param { get; set; } = null!;

        /// <summary>
        /// The value of parameter.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public DbInstanceInitParamArgs()
        {
        }
        public static new DbInstanceInitParamArgs Empty => new DbInstanceInitParamArgs();
    }
}
