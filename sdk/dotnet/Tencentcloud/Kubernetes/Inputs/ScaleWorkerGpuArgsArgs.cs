// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Kubernetes.Inputs
{

    public sealed class ScaleWorkerGpuArgsArgs : Pulumi.ResourceArgs
    {
        [Input("cuda")]
        private InputMap<object>? _cuda;

        /// <summary>
        /// CUDA  version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.
        /// </summary>
        public InputMap<object> Cuda
        {
            get => _cuda ?? (_cuda = new InputMap<object>());
            set => _cuda = value;
        }

        [Input("cudnn")]
        private InputMap<object>? _cudnn;

        /// <summary>
        /// cuDNN version. Format like: `{ version: String, name: String, doc_name: String, dev_name: String }`. `version`: cuDNN version; `name`: cuDNN name; `doc_name`: Doc name of cuDNN; `dev_name`: Dev name of cuDNN.
        /// </summary>
        public InputMap<object> Cudnn
        {
            get => _cudnn ?? (_cudnn = new InputMap<object>());
            set => _cudnn = value;
        }

        [Input("customDriver")]
        private InputMap<object>? _customDriver;

        /// <summary>
        /// Custom GPU driver. Format like: `{address: String}`. `address`: URL of custom GPU driver address.
        /// </summary>
        public InputMap<object> CustomDriver
        {
            get => _customDriver ?? (_customDriver = new InputMap<object>());
            set => _customDriver = value;
        }

        [Input("driver")]
        private InputMap<object>? _driver;

        /// <summary>
        /// GPU driver version. Format like: `{ version: String, name: String }`. `version`: Version of GPU driver or CUDA; `name`: Name of GPU driver or CUDA.
        /// </summary>
        public InputMap<object> Driver
        {
            get => _driver ?? (_driver = new InputMap<object>());
            set => _driver = value;
        }

        /// <summary>
        /// Whether to enable MIG.
        /// </summary>
        [Input("migEnable")]
        public Input<bool>? MigEnable { get; set; }

        public ScaleWorkerGpuArgsArgs()
        {
        }
    }
}
