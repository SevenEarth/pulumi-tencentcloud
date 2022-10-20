// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Scf
{
    /// <summary>
    /// Provide a resource to create a SCF function.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Scf.Function("foo", new Tencentcloud.Scf.FunctionArgs
    ///         {
    ///             CosBucketName = "scf-code-1234567890",
    ///             CosBucketRegion = "ap-guangzhou",
    ///             CosObjectName = "code.zip",
    ///             Handler = "main.do_it",
    ///             Runtime = "Python3.6",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// Using CFS config
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Scf.Function("foo", new Tencentcloud.Scf.FunctionArgs
    ///         {
    ///             CfsConfigs = 
    ///             {
    ///                 new Tencentcloud.Scf.Inputs.FunctionCfsConfigArgs
    ///                 {
    ///                     CfsId = "cfs-xxxxxxxx",
    ///                     LocalMountDir = "/mnt",
    ///                     MountInsId = "cfs-xxxxxxxx",
    ///                     RemoteMountDir = "/",
    ///                     UserGroupId = "10000",
    ///                     UserId = "10000",
    ///                 },
    ///             },
    ///             Handler = "main.do_it",
    ///             Runtime = "Python3.6",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// SCF function can be imported, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Scf/function:Function test default+test
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Scf/function:Function")]
    public partial class Function : Pulumi.CustomResource
    {
        /// <summary>
        /// List of CFS configurations.
        /// </summary>
        [Output("cfsConfigs")]
        public Output<ImmutableArray<Outputs.FunctionCfsConfig>> CfsConfigs { get; private set; } = null!;

        /// <summary>
        /// cls logset id of the SCF function.
        /// </summary>
        [Output("clsLogsetId")]
        public Output<string> ClsLogsetId { get; private set; } = null!;

        /// <summary>
        /// cls topic id of the SCF function.
        /// </summary>
        [Output("clsTopicId")]
        public Output<string> ClsTopicId { get; private set; } = null!;

        /// <summary>
        /// SCF function code error message.
        /// </summary>
        [Output("codeError")]
        public Output<string> CodeError { get; private set; } = null!;

        /// <summary>
        /// SCF function code is correct.
        /// </summary>
        [Output("codeResult")]
        public Output<string> CodeResult { get; private set; } = null!;

        /// <summary>
        /// SCF function code size, unit is M.
        /// </summary>
        [Output("codeSize")]
        public Output<int> CodeSize { get; private set; } = null!;

        /// <summary>
        /// Cos bucket name of the SCF function, such as `cos-1234567890`, conflict with `zip_file`.
        /// </summary>
        [Output("cosBucketName")]
        public Output<string?> CosBucketName { get; private set; } = null!;

        /// <summary>
        /// Cos bucket region of the SCF function, conflict with `zip_file`.
        /// </summary>
        [Output("cosBucketRegion")]
        public Output<string?> CosBucketRegion { get; private set; } = null!;

        /// <summary>
        /// Cos object name of the SCF function, should have suffix `.zip` or `.jar`, conflict with `zip_file`.
        /// </summary>
        [Output("cosObjectName")]
        public Output<string?> CosObjectName { get; private set; } = null!;

        /// <summary>
        /// Description of the SCF function. Description supports English letters, numbers, spaces, commas, newlines, periods and Chinese, the maximum length is 1000.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether EIP is a fixed IP.
        /// </summary>
        [Output("eipFixed")]
        public Output<bool> EipFixed { get; private set; } = null!;

        /// <summary>
        /// SCF function EIP list.
        /// </summary>
        [Output("eips")]
        public Output<ImmutableArray<string>> Eips { get; private set; } = null!;

        /// <summary>
        /// Indicates whether EIP config set to `ENABLE` when `enable_public_net` was true.
        /// </summary>
        [Output("enableEipConfig")]
        public Output<bool?> EnableEipConfig { get; private set; } = null!;

        /// <summary>
        /// Indicates whether public net config enabled. NOTE: only `vpc_id` specified can disable public net config.
        /// </summary>
        [Output("enablePublicNet")]
        public Output<bool?> EnablePublicNet { get; private set; } = null!;

        /// <summary>
        /// Environment of the SCF function.
        /// </summary>
        [Output("environment")]
        public Output<ImmutableDictionary<string, object>?> Environment { get; private set; } = null!;

        /// <summary>
        /// SCF function code error code.
        /// </summary>
        [Output("errNo")]
        public Output<int> ErrNo { get; private set; } = null!;

        /// <summary>
        /// Handler of the SCF function. The format of name is `&lt;filename&gt;.&lt;method_name&gt;`, and it supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be `-` or `_`. Available length is 2-60.
        /// </summary>
        [Output("handler")]
        public Output<string> Handler { get; private set; } = null!;

        /// <summary>
        /// SCF function domain name.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        /// <summary>
        /// Image of the SCF function, conflict with ``.
        /// </summary>
        [Output("imageConfigs")]
        public Output<ImmutableArray<Outputs.FunctionImageConfig>> ImageConfigs { get; private set; } = null!;

        /// <summary>
        /// Whether to automatically install dependencies.
        /// </summary>
        [Output("installDependency")]
        public Output<bool> InstallDependency { get; private set; } = null!;

        /// <summary>
        /// Enable L5 for SCF function, default is `false`.
        /// </summary>
        [Output("l5Enable")]
        public Output<bool?> L5Enable { get; private set; } = null!;

        /// <summary>
        /// The list of association layers.
        /// </summary>
        [Output("layers")]
        public Output<ImmutableArray<Outputs.FunctionLayer>> Layers { get; private set; } = null!;

        /// <summary>
        /// Memory size of the SCF function, unit is MB. The default is `128`MB. The ladder is 128M.
        /// </summary>
        [Output("memSize")]
        public Output<int?> MemSize { get; private set; } = null!;

        /// <summary>
        /// Modify time of SCF function trigger.
        /// </summary>
        [Output("modifyTime")]
        public Output<string> ModifyTime { get; private set; } = null!;

        /// <summary>
        /// Name of the SCF function. Name supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be `-` or `_`. Available length is 2-60.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Namespace of the SCF function, default is `default`.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Role of the SCF function.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;

        /// <summary>
        /// Runtime of the SCF function, only supports `Python2.7`, `Python3.6`, `Nodejs6.10`, `Nodejs8.9`, `Nodejs10.15`, `PHP5`, `PHP7`, `Golang1`, and `Java8`.
        /// </summary>
        [Output("runtime")]
        public Output<string> Runtime { get; private set; } = null!;

        /// <summary>
        /// SCF function status.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// SCF status description.
        /// </summary>
        [Output("statusDesc")]
        public Output<string> StatusDesc { get; private set; } = null!;

        /// <summary>
        /// Subnet ID of the SCF function.
        /// </summary>
        [Output("subnetId")]
        public Output<string?> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Tags of the SCF function.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Timeout of the SCF function, unit is second. Default `3`. Available value is 1-900.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// SCF trigger details list. Each element contains the following attributes:
        /// </summary>
        [Output("triggerInfos")]
        public Output<ImmutableArray<Outputs.FunctionTriggerInfo>> TriggerInfos { get; private set; } = null!;

        /// <summary>
        /// Trigger list of the SCF function, note that if you modify the trigger list, all existing triggers will be deleted, and then create triggers in the new list. Each element contains the following attributes:
        /// </summary>
        [Output("triggers")]
        public Output<ImmutableArray<Outputs.FunctionTrigger>> Triggers { get; private set; } = null!;

        /// <summary>
        /// SCF function vip.
        /// </summary>
        [Output("vip")]
        public Output<string> Vip { get; private set; } = null!;

        /// <summary>
        /// VPC ID of the SCF function.
        /// </summary>
        [Output("vpcId")]
        public Output<string?> VpcId { get; private set; } = null!;

        /// <summary>
        /// Zip file of the SCF function, conflict with `cos_bucket_name`, `cos_object_name`, `cos_bucket_region`.
        /// </summary>
        [Output("zipFile")]
        public Output<string?> ZipFile { get; private set; } = null!;


        /// <summary>
        /// Create a Function resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Function(string name, FunctionArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/function:Function", name, args ?? new FunctionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Function(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Scf/function:Function", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Function resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Function Get(string name, Input<string> id, FunctionState? state = null, CustomResourceOptions? options = null)
        {
            return new Function(name, id, state, options);
        }
    }

    public sealed class FunctionArgs : Pulumi.ResourceArgs
    {
        [Input("cfsConfigs")]
        private InputList<Inputs.FunctionCfsConfigArgs>? _cfsConfigs;

        /// <summary>
        /// List of CFS configurations.
        /// </summary>
        public InputList<Inputs.FunctionCfsConfigArgs> CfsConfigs
        {
            get => _cfsConfigs ?? (_cfsConfigs = new InputList<Inputs.FunctionCfsConfigArgs>());
            set => _cfsConfigs = value;
        }

        /// <summary>
        /// cls logset id of the SCF function.
        /// </summary>
        [Input("clsLogsetId")]
        public Input<string>? ClsLogsetId { get; set; }

        /// <summary>
        /// cls topic id of the SCF function.
        /// </summary>
        [Input("clsTopicId")]
        public Input<string>? ClsTopicId { get; set; }

        /// <summary>
        /// Cos bucket name of the SCF function, such as `cos-1234567890`, conflict with `zip_file`.
        /// </summary>
        [Input("cosBucketName")]
        public Input<string>? CosBucketName { get; set; }

        /// <summary>
        /// Cos bucket region of the SCF function, conflict with `zip_file`.
        /// </summary>
        [Input("cosBucketRegion")]
        public Input<string>? CosBucketRegion { get; set; }

        /// <summary>
        /// Cos object name of the SCF function, should have suffix `.zip` or `.jar`, conflict with `zip_file`.
        /// </summary>
        [Input("cosObjectName")]
        public Input<string>? CosObjectName { get; set; }

        /// <summary>
        /// Description of the SCF function. Description supports English letters, numbers, spaces, commas, newlines, periods and Chinese, the maximum length is 1000.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates whether EIP config set to `ENABLE` when `enable_public_net` was true.
        /// </summary>
        [Input("enableEipConfig")]
        public Input<bool>? EnableEipConfig { get; set; }

        /// <summary>
        /// Indicates whether public net config enabled. NOTE: only `vpc_id` specified can disable public net config.
        /// </summary>
        [Input("enablePublicNet")]
        public Input<bool>? EnablePublicNet { get; set; }

        [Input("environment")]
        private InputMap<object>? _environment;

        /// <summary>
        /// Environment of the SCF function.
        /// </summary>
        public InputMap<object> Environment
        {
            get => _environment ?? (_environment = new InputMap<object>());
            set => _environment = value;
        }

        /// <summary>
        /// Handler of the SCF function. The format of name is `&lt;filename&gt;.&lt;method_name&gt;`, and it supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be `-` or `_`. Available length is 2-60.
        /// </summary>
        [Input("handler", required: true)]
        public Input<string> Handler { get; set; } = null!;

        [Input("imageConfigs")]
        private InputList<Inputs.FunctionImageConfigArgs>? _imageConfigs;

        /// <summary>
        /// Image of the SCF function, conflict with ``.
        /// </summary>
        public InputList<Inputs.FunctionImageConfigArgs> ImageConfigs
        {
            get => _imageConfigs ?? (_imageConfigs = new InputList<Inputs.FunctionImageConfigArgs>());
            set => _imageConfigs = value;
        }

        /// <summary>
        /// Enable L5 for SCF function, default is `false`.
        /// </summary>
        [Input("l5Enable")]
        public Input<bool>? L5Enable { get; set; }

        [Input("layers")]
        private InputList<Inputs.FunctionLayerArgs>? _layers;

        /// <summary>
        /// The list of association layers.
        /// </summary>
        public InputList<Inputs.FunctionLayerArgs> Layers
        {
            get => _layers ?? (_layers = new InputList<Inputs.FunctionLayerArgs>());
            set => _layers = value;
        }

        /// <summary>
        /// Memory size of the SCF function, unit is MB. The default is `128`MB. The ladder is 128M.
        /// </summary>
        [Input("memSize")]
        public Input<int>? MemSize { get; set; }

        /// <summary>
        /// Name of the SCF function. Name supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be `-` or `_`. Available length is 2-60.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Namespace of the SCF function, default is `default`.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Role of the SCF function.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Runtime of the SCF function, only supports `Python2.7`, `Python3.6`, `Nodejs6.10`, `Nodejs8.9`, `Nodejs10.15`, `PHP5`, `PHP7`, `Golang1`, and `Java8`.
        /// </summary>
        [Input("runtime", required: true)]
        public Input<string> Runtime { get; set; } = null!;

        /// <summary>
        /// Subnet ID of the SCF function.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the SCF function.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Timeout of the SCF function, unit is second. Default `3`. Available value is 1-900.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("triggers")]
        private InputList<Inputs.FunctionTriggerArgs>? _triggers;

        /// <summary>
        /// Trigger list of the SCF function, note that if you modify the trigger list, all existing triggers will be deleted, and then create triggers in the new list. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.FunctionTriggerArgs> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<Inputs.FunctionTriggerArgs>());
            set => _triggers = value;
        }

        /// <summary>
        /// VPC ID of the SCF function.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Zip file of the SCF function, conflict with `cos_bucket_name`, `cos_object_name`, `cos_bucket_region`.
        /// </summary>
        [Input("zipFile")]
        public Input<string>? ZipFile { get; set; }

        public FunctionArgs()
        {
        }
    }

    public sealed class FunctionState : Pulumi.ResourceArgs
    {
        [Input("cfsConfigs")]
        private InputList<Inputs.FunctionCfsConfigGetArgs>? _cfsConfigs;

        /// <summary>
        /// List of CFS configurations.
        /// </summary>
        public InputList<Inputs.FunctionCfsConfigGetArgs> CfsConfigs
        {
            get => _cfsConfigs ?? (_cfsConfigs = new InputList<Inputs.FunctionCfsConfigGetArgs>());
            set => _cfsConfigs = value;
        }

        /// <summary>
        /// cls logset id of the SCF function.
        /// </summary>
        [Input("clsLogsetId")]
        public Input<string>? ClsLogsetId { get; set; }

        /// <summary>
        /// cls topic id of the SCF function.
        /// </summary>
        [Input("clsTopicId")]
        public Input<string>? ClsTopicId { get; set; }

        /// <summary>
        /// SCF function code error message.
        /// </summary>
        [Input("codeError")]
        public Input<string>? CodeError { get; set; }

        /// <summary>
        /// SCF function code is correct.
        /// </summary>
        [Input("codeResult")]
        public Input<string>? CodeResult { get; set; }

        /// <summary>
        /// SCF function code size, unit is M.
        /// </summary>
        [Input("codeSize")]
        public Input<int>? CodeSize { get; set; }

        /// <summary>
        /// Cos bucket name of the SCF function, such as `cos-1234567890`, conflict with `zip_file`.
        /// </summary>
        [Input("cosBucketName")]
        public Input<string>? CosBucketName { get; set; }

        /// <summary>
        /// Cos bucket region of the SCF function, conflict with `zip_file`.
        /// </summary>
        [Input("cosBucketRegion")]
        public Input<string>? CosBucketRegion { get; set; }

        /// <summary>
        /// Cos object name of the SCF function, should have suffix `.zip` or `.jar`, conflict with `zip_file`.
        /// </summary>
        [Input("cosObjectName")]
        public Input<string>? CosObjectName { get; set; }

        /// <summary>
        /// Description of the SCF function. Description supports English letters, numbers, spaces, commas, newlines, periods and Chinese, the maximum length is 1000.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether EIP is a fixed IP.
        /// </summary>
        [Input("eipFixed")]
        public Input<bool>? EipFixed { get; set; }

        [Input("eips")]
        private InputList<string>? _eips;

        /// <summary>
        /// SCF function EIP list.
        /// </summary>
        public InputList<string> Eips
        {
            get => _eips ?? (_eips = new InputList<string>());
            set => _eips = value;
        }

        /// <summary>
        /// Indicates whether EIP config set to `ENABLE` when `enable_public_net` was true.
        /// </summary>
        [Input("enableEipConfig")]
        public Input<bool>? EnableEipConfig { get; set; }

        /// <summary>
        /// Indicates whether public net config enabled. NOTE: only `vpc_id` specified can disable public net config.
        /// </summary>
        [Input("enablePublicNet")]
        public Input<bool>? EnablePublicNet { get; set; }

        [Input("environment")]
        private InputMap<object>? _environment;

        /// <summary>
        /// Environment of the SCF function.
        /// </summary>
        public InputMap<object> Environment
        {
            get => _environment ?? (_environment = new InputMap<object>());
            set => _environment = value;
        }

        /// <summary>
        /// SCF function code error code.
        /// </summary>
        [Input("errNo")]
        public Input<int>? ErrNo { get; set; }

        /// <summary>
        /// Handler of the SCF function. The format of name is `&lt;filename&gt;.&lt;method_name&gt;`, and it supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be `-` or `_`. Available length is 2-60.
        /// </summary>
        [Input("handler")]
        public Input<string>? Handler { get; set; }

        /// <summary>
        /// SCF function domain name.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        [Input("imageConfigs")]
        private InputList<Inputs.FunctionImageConfigGetArgs>? _imageConfigs;

        /// <summary>
        /// Image of the SCF function, conflict with ``.
        /// </summary>
        public InputList<Inputs.FunctionImageConfigGetArgs> ImageConfigs
        {
            get => _imageConfigs ?? (_imageConfigs = new InputList<Inputs.FunctionImageConfigGetArgs>());
            set => _imageConfigs = value;
        }

        /// <summary>
        /// Whether to automatically install dependencies.
        /// </summary>
        [Input("installDependency")]
        public Input<bool>? InstallDependency { get; set; }

        /// <summary>
        /// Enable L5 for SCF function, default is `false`.
        /// </summary>
        [Input("l5Enable")]
        public Input<bool>? L5Enable { get; set; }

        [Input("layers")]
        private InputList<Inputs.FunctionLayerGetArgs>? _layers;

        /// <summary>
        /// The list of association layers.
        /// </summary>
        public InputList<Inputs.FunctionLayerGetArgs> Layers
        {
            get => _layers ?? (_layers = new InputList<Inputs.FunctionLayerGetArgs>());
            set => _layers = value;
        }

        /// <summary>
        /// Memory size of the SCF function, unit is MB. The default is `128`MB. The ladder is 128M.
        /// </summary>
        [Input("memSize")]
        public Input<int>? MemSize { get; set; }

        /// <summary>
        /// Modify time of SCF function trigger.
        /// </summary>
        [Input("modifyTime")]
        public Input<string>? ModifyTime { get; set; }

        /// <summary>
        /// Name of the SCF function. Name supports 26 English letters, numbers, connectors, and underscores, it should start with a letter. The last character cannot be `-` or `_`. Available length is 2-60.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Namespace of the SCF function, default is `default`.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Role of the SCF function.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Runtime of the SCF function, only supports `Python2.7`, `Python3.6`, `Nodejs6.10`, `Nodejs8.9`, `Nodejs10.15`, `PHP5`, `PHP7`, `Golang1`, and `Java8`.
        /// </summary>
        [Input("runtime")]
        public Input<string>? Runtime { get; set; }

        /// <summary>
        /// SCF function status.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// SCF status description.
        /// </summary>
        [Input("statusDesc")]
        public Input<string>? StatusDesc { get; set; }

        /// <summary>
        /// Subnet ID of the SCF function.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// Tags of the SCF function.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        /// <summary>
        /// Timeout of the SCF function, unit is second. Default `3`. Available value is 1-900.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        [Input("triggerInfos")]
        private InputList<Inputs.FunctionTriggerInfoGetArgs>? _triggerInfos;

        /// <summary>
        /// SCF trigger details list. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.FunctionTriggerInfoGetArgs> TriggerInfos
        {
            get => _triggerInfos ?? (_triggerInfos = new InputList<Inputs.FunctionTriggerInfoGetArgs>());
            set => _triggerInfos = value;
        }

        [Input("triggers")]
        private InputList<Inputs.FunctionTriggerGetArgs>? _triggers;

        /// <summary>
        /// Trigger list of the SCF function, note that if you modify the trigger list, all existing triggers will be deleted, and then create triggers in the new list. Each element contains the following attributes:
        /// </summary>
        public InputList<Inputs.FunctionTriggerGetArgs> Triggers
        {
            get => _triggers ?? (_triggers = new InputList<Inputs.FunctionTriggerGetArgs>());
            set => _triggers = value;
        }

        /// <summary>
        /// SCF function vip.
        /// </summary>
        [Input("vip")]
        public Input<string>? Vip { get; set; }

        /// <summary>
        /// VPC ID of the SCF function.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// Zip file of the SCF function, conflict with `cos_bucket_name`, `cos_object_name`, `cos_bucket_region`.
        /// </summary>
        [Input("zipFile")]
        public Input<string>? ZipFile { get; set; }

        public FunctionState()
        {
        }
    }
}
