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

    public sealed class ParamTemplateParamListGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value of parameter.
        /// </summary>
        [Input("currentValue")]
        public Input<string>? CurrentValue { get; set; }

        /// <summary>
        /// The name of parameter.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ParamTemplateParamListGetArgs()
        {
        }
        public static new ParamTemplateParamListGetArgs Empty => new ParamTemplateParamListGetArgs();
    }
}
