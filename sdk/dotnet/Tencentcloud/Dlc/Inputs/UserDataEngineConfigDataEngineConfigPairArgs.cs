// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Dlc.Inputs
{

    public sealed class UserDataEngineConfigDataEngineConfigPairArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Config key.
        /// </summary>
        [Input("configItem", required: true)]
        public Input<string> ConfigItem { get; set; } = null!;

        /// <summary>
        /// Config value.
        /// </summary>
        [Input("configValue", required: true)]
        public Input<string> ConfigValue { get; set; } = null!;

        public UserDataEngineConfigDataEngineConfigPairArgs()
        {
        }
        public static new UserDataEngineConfigDataEngineConfigPairArgs Empty => new UserDataEngineConfigDataEngineConfigPairArgs();
    }
}
