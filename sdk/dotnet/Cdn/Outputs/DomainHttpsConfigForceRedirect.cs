// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.Cdn.Outputs
{

    [OutputType]
    public sealed class DomainHttpsConfigForceRedirect
    {
        /// <summary>
        /// Forced redirect status code. Valid values are `301` and `302`. When `switch` setting `off`, this property does not need to be set or set to `302`. Default value is `302`.
        /// </summary>
        public readonly int? RedirectStatusCode;
        /// <summary>
        /// Forced redirect type. Valid values are `http` and `https`. `http` means a forced redirect from HTTPS to HTTP, `https` means a forced redirect from HTTP to HTTPS. When `switch` setting `off`, this property does not need to be set or set to `http`. Default value is `http`.
        /// </summary>
        public readonly string? RedirectType;
        /// <summary>
        /// Forced redirect configuration switch. Valid values are `on` and `off`. Default value is `off`.
        /// </summary>
        public readonly string? Switch;

        [OutputConstructor]
        private DomainHttpsConfigForceRedirect(
            int? redirectStatusCode,

            string? redirectType,

            string? @switch)
        {
            RedirectStatusCode = redirectStatusCode;
            RedirectType = redirectType;
            Switch = @switch;
        }
    }
}
