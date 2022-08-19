// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tencentcloud.ApiGateway.Outputs
{

    [OutputType]
    public sealed class GetCustomerDomainsListPathMappingResult
    {
        /// <summary>
        /// Release environment.
        /// </summary>
        public readonly string Environment;
        /// <summary>
        /// The domain mapping path.
        /// </summary>
        public readonly string Path;

        [OutputConstructor]
        private GetCustomerDomainsListPathMappingResult(
            string environment,

            string path)
        {
            Environment = environment;
            Path = path;
        }
    }
}
