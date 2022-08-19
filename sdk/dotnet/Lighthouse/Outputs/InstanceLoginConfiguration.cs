// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Lighthouse.Outputs
{

    [OutputType]
    public sealed class InstanceLoginConfiguration
    {
        /// <summary>
        /// whether auto generate password. if false, need set password.
        /// </summary>
        public readonly string AutoGeneratePassword;
        /// <summary>
        /// Login password.
        /// </summary>
        public readonly string? Password;

        [OutputConstructor]
        private InstanceLoginConfiguration(
            string autoGeneratePassword,

            string? password)
        {
            AutoGeneratePassword = autoGeneratePassword;
            Password = password;
        }
    }
}
