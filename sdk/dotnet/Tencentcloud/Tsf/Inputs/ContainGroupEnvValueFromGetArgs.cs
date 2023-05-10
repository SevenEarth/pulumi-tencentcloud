// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tsf.Inputs
{

    public sealed class ContainGroupEnvValueFromGetArgs : Pulumi.ResourceArgs
    {
        [Input("fieldReves")]
        private InputList<Inputs.ContainGroupEnvValueFromFieldRefGetArgs>? _fieldReves;

        /// <summary>
        /// FieldRef for k8s env.
        /// </summary>
        public InputList<Inputs.ContainGroupEnvValueFromFieldRefGetArgs> FieldReves
        {
            get => _fieldReves ?? (_fieldReves = new InputList<Inputs.ContainGroupEnvValueFromFieldRefGetArgs>());
            set => _fieldReves = value;
        }

        [Input("resourceFieldReves")]
        private InputList<Inputs.ContainGroupEnvValueFromResourceFieldRefGetArgs>? _resourceFieldReves;

        /// <summary>
        /// ResourceFieldRef of k8s env.
        /// </summary>
        public InputList<Inputs.ContainGroupEnvValueFromResourceFieldRefGetArgs> ResourceFieldReves
        {
            get => _resourceFieldReves ?? (_resourceFieldReves = new InputList<Inputs.ContainGroupEnvValueFromResourceFieldRefGetArgs>());
            set => _resourceFieldReves = value;
        }

        public ContainGroupEnvValueFromGetArgs()
        {
        }
    }
}