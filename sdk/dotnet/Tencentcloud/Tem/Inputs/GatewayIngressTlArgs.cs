// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tem.Inputs
{

    public sealed class GatewayIngressTlArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// certificate ID.
        /// </summary>
        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("hosts", required: true)]
        private InputList<string>? _hosts;

        /// <summary>
        /// host names.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// secret name, if you use a certificate, you don't need to fill in this field.
        /// </summary>
        [Input("secretName")]
        public Input<string>? SecretName { get; set; }

        public GatewayIngressTlArgs()
        {
        }
    }
}
