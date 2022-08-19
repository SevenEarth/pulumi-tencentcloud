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
    public sealed class DomainAuthentication
    {
        /// <summary>
        /// Authentication switching, available values: `on`, `off`.
        /// </summary>
        public readonly string? Switch;
        /// <summary>
        /// Timestamp hotlink protection mode A configuration.
        /// </summary>
        public readonly Outputs.DomainAuthenticationTypeA? TypeA;
        /// <summary>
        /// Timestamp hotlink protection mode B configuration. NOTE: according to upgrading of TencentCloud Platform, TypeB is unavailable for now.
        /// </summary>
        public readonly Outputs.DomainAuthenticationTypeB? TypeB;
        /// <summary>
        /// Timestamp hotlink protection mode C configuration.
        /// </summary>
        public readonly Outputs.DomainAuthenticationTypeC? TypeC;
        /// <summary>
        /// Timestamp hotlink protection mode D configuration.
        /// </summary>
        public readonly Outputs.DomainAuthenticationTypeD? TypeD;

        [OutputConstructor]
        private DomainAuthentication(
            string? @switch,

            Outputs.DomainAuthenticationTypeA? typeA,

            Outputs.DomainAuthenticationTypeB? typeB,

            Outputs.DomainAuthenticationTypeC? typeC,

            Outputs.DomainAuthenticationTypeD? typeD)
        {
            Switch = @switch;
            TypeA = typeA;
            TypeB = typeB;
            TypeC = typeC;
            TypeD = typeD;
        }
    }
}
