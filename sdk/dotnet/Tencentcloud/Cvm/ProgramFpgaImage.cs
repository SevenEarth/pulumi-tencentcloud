// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Cvm
{
    [TencentcloudResourceType("tencentcloud:Cvm/programFpgaImage:ProgramFpgaImage")]
    public partial class ProgramFpgaImage : Pulumi.CustomResource
    {
        /// <summary>
        /// The DBDF number of the FPGA card on the instance, if left blank, the FPGA image will be burned to all FPGA cards owned
        /// by the instance by default.
        /// </summary>
        [Output("dbdFs")]
        public Output<ImmutableArray<string>> DbdFs { get; private set; } = null!;

        /// <summary>
        /// Trial run, will not perform the actual burning action, the default is False.
        /// </summary>
        [Output("dryRun")]
        public Output<bool?> DryRun { get; private set; } = null!;

        /// <summary>
        /// COS URL address of the FPGA image file.
        /// </summary>
        [Output("fpgaUrl")]
        public Output<string> FpgaUrl { get; private set; } = null!;

        /// <summary>
        /// The ID information of the instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a ProgramFpgaImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProgramFpgaImage(string name, ProgramFpgaImageArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/programFpgaImage:ProgramFpgaImage", name, args ?? new ProgramFpgaImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProgramFpgaImage(string name, Input<string> id, ProgramFpgaImageState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Cvm/programFpgaImage:ProgramFpgaImage", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/tencentcloudstack",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ProgramFpgaImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProgramFpgaImage Get(string name, Input<string> id, ProgramFpgaImageState? state = null, CustomResourceOptions? options = null)
        {
            return new ProgramFpgaImage(name, id, state, options);
        }
    }

    public sealed class ProgramFpgaImageArgs : Pulumi.ResourceArgs
    {
        [Input("dbdFs")]
        private InputList<string>? _dbdFs;

        /// <summary>
        /// The DBDF number of the FPGA card on the instance, if left blank, the FPGA image will be burned to all FPGA cards owned
        /// by the instance by default.
        /// </summary>
        public InputList<string> DbdFs
        {
            get => _dbdFs ?? (_dbdFs = new InputList<string>());
            set => _dbdFs = value;
        }

        /// <summary>
        /// Trial run, will not perform the actual burning action, the default is False.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// COS URL address of the FPGA image file.
        /// </summary>
        [Input("fpgaUrl", required: true)]
        public Input<string> FpgaUrl { get; set; } = null!;

        /// <summary>
        /// The ID information of the instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public ProgramFpgaImageArgs()
        {
        }
    }

    public sealed class ProgramFpgaImageState : Pulumi.ResourceArgs
    {
        [Input("dbdFs")]
        private InputList<string>? _dbdFs;

        /// <summary>
        /// The DBDF number of the FPGA card on the instance, if left blank, the FPGA image will be burned to all FPGA cards owned
        /// by the instance by default.
        /// </summary>
        public InputList<string> DbdFs
        {
            get => _dbdFs ?? (_dbdFs = new InputList<string>());
            set => _dbdFs = value;
        }

        /// <summary>
        /// Trial run, will not perform the actual burning action, the default is False.
        /// </summary>
        [Input("dryRun")]
        public Input<bool>? DryRun { get; set; }

        /// <summary>
        /// COS URL address of the FPGA image file.
        /// </summary>
        [Input("fpgaUrl")]
        public Input<string>? FpgaUrl { get; set; }

        /// <summary>
        /// The ID information of the instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public ProgramFpgaImageState()
        {
        }
    }
}
