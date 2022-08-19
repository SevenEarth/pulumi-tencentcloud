// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse.Inputs
{

    public sealed class InstanceLoginConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// whether auto generate password. if false, need set password.
        /// </summary>
        [Input("autoGeneratePassword", required: true)]
        public Input<string> AutoGeneratePassword { get; set; } = null!;

        /// <summary>
        /// Login password.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        public InstanceLoginConfigurationGetArgs()
        {
        }
    }
}
