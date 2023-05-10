// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Tat
{
    /// <summary>
    /// Provides a resource to create a tat command
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
    ///         var command = new Tencentcloud.Tat.Command("command", new Tencentcloud.Tat.CommandArgs
    ///         {
    ///             CommandName = "ls",
    ///             CommandType = "SHELL",
    ///             Content = "bHM=",
    ///             Description = "xxx",
    ///             Tags = 
    ///             {
    ///                 new Tencentcloud.Tat.Inputs.CommandTagArgs
    ///                 {
    ///                     Key = "",
    ///                     Value = "",
    ///                 },
    ///             },
    ///             Timeout = 50,
    ///             Username = "root",
    ///             WorkingDirectory = "/root",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// tat command can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Tat/command:Command command cmd-6fydo27j
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Tat/command:Command")]
    public partial class Command : Pulumi.CustomResource
    {
        /// <summary>
        /// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
        /// </summary>
        [Output("commandName")]
        public Output<string> CommandName { get; private set; } = null!;

        /// <summary>
        /// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
        /// </summary>
        [Output("commandType")]
        public Output<string?> CommandType { get; private set; } = null!;

        /// <summary>
        /// Command. The maximum length of Base64 encoding is 64KB.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
        /// </summary>
        [Output("createdBy")]
        public Output<string> CreatedBy { get; private set; } = null!;

        /// <summary>
        /// Command creation time.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;amp;#39;varA&amp;amp;#39;: &amp;amp;#39;222&amp;amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        /// </summary>
        [Output("defaultParameters")]
        public Output<string?> DefaultParameters { get; private set; } = null!;

        /// <summary>
        /// Command description. The maximum length is 120 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
        /// </summary>
        [Output("enableParameter")]
        public Output<bool?> EnableParameter { get; private set; } = null!;

        /// <summary>
        /// Formatted description of the command. This parameter is an empty string for user commands and contains values for public commands.
        /// </summary>
        [Output("formattedDescription")]
        public Output<string> FormattedDescription { get; private set; } = null!;

        /// <summary>
        /// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
        /// </summary>
        [Output("outputCosBucketUrl")]
        public Output<string?> OutputCosBucketUrl { get; private set; } = null!;

        /// <summary>
        /// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
        /// </summary>
        [Output("outputCosKeyPrefix")]
        public Output<string?> OutputCosKeyPrefix { get; private set; } = null!;

        /// <summary>
        /// Tags bound to the command. At most 10 tags are allowed.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.CommandTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// Command update time.
        /// </summary>
        [Output("updatedTime")]
        public Output<string> UpdatedTime { get; private set; } = null!;

        /// <summary>
        /// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
        /// </summary>
        [Output("username")]
        public Output<string?> Username { get; private set; } = null!;

        /// <summary>
        /// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
        /// </summary>
        [Output("workingDirectory")]
        public Output<string?> WorkingDirectory { get; private set; } = null!;


        /// <summary>
        /// Create a Command resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Command(string name, CommandArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Tat/command:Command", name, args ?? new CommandArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Command(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Tat/command:Command", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Command resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Command Get(string name, Input<string> id, CommandState? state = null, CustomResourceOptions? options = null)
        {
            return new Command(name, id, state, options);
        }
    }

    public sealed class CommandArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
        /// </summary>
        [Input("commandName", required: true)]
        public Input<string> CommandName { get; set; } = null!;

        /// <summary>
        /// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
        /// </summary>
        [Input("commandType")]
        public Input<string>? CommandType { get; set; }

        /// <summary>
        /// Command. The maximum length of Base64 encoding is 64KB.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;amp;#39;varA&amp;amp;#39;: &amp;amp;#39;222&amp;amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        /// </summary>
        [Input("defaultParameters")]
        public Input<string>? DefaultParameters { get; set; }

        /// <summary>
        /// Command description. The maximum length is 120 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
        /// </summary>
        [Input("enableParameter")]
        public Input<bool>? EnableParameter { get; set; }

        /// <summary>
        /// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
        /// </summary>
        [Input("outputCosBucketUrl")]
        public Input<string>? OutputCosBucketUrl { get; set; }

        /// <summary>
        /// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
        /// </summary>
        [Input("outputCosKeyPrefix")]
        public Input<string>? OutputCosKeyPrefix { get; set; }

        [Input("tags")]
        private InputList<Inputs.CommandTagArgs>? _tags;

        /// <summary>
        /// Tags bound to the command. At most 10 tags are allowed.
        /// </summary>
        public InputList<Inputs.CommandTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CommandTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
        /// </summary>
        [Input("workingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public CommandArgs()
        {
        }
    }

    public sealed class CommandState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Command name. The name can be up to 60 bytes, and contain [a-z], [A-Z], [0-9] and [_-.].
        /// </summary>
        [Input("commandName")]
        public Input<string>? CommandName { get; set; }

        /// <summary>
        /// Command type. `SHELL` and `POWERSHELL` are supported. The default value is `SHELL`.
        /// </summary>
        [Input("commandType")]
        public Input<string>? CommandType { get; set; }

        /// <summary>
        /// Command. The maximum length of Base64 encoding is 64KB.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Command creator. `TAT` indicates a public command and `USER` indicates a personal command.
        /// </summary>
        [Input("createdBy")]
        public Input<string>? CreatedBy { get; set; }

        /// <summary>
        /// Command creation time.
        /// </summary>
        [Input("createdTime")]
        public Input<string>? CreatedTime { get; set; }

        /// <summary>
        /// The default value of the custom parameter value when it is enabled. The field type is JSON encoded string. For example, {&amp;amp;#39;varA&amp;amp;#39;: &amp;amp;#39;222&amp;amp;#39;}.`key` is the name of the custom parameter and value is the default value. Both `key` and `value` are strings.If no parameter value is provided in the `InvokeCommand` API, the default value is used.Up to 20 custom parameters are supported.The name of the custom parameter cannot exceed 64 characters and can contain [a-z], [A-Z], [0-9] and [-_].
        /// </summary>
        [Input("defaultParameters")]
        public Input<string>? DefaultParameters { get; set; }

        /// <summary>
        /// Command description. The maximum length is 120 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether to enable the custom parameter feature.This cannot be modified once created.Default value: `false`.
        /// </summary>
        [Input("enableParameter")]
        public Input<bool>? EnableParameter { get; set; }

        /// <summary>
        /// Formatted description of the command. This parameter is an empty string for user commands and contains values for public commands.
        /// </summary>
        [Input("formattedDescription")]
        public Input<string>? FormattedDescription { get; set; }

        /// <summary>
        /// The COS bucket URL for uploading logs. The URL must start with `https`, such as `https://BucketName-123454321.cos.ap-beijing.myqcloud.com`.
        /// </summary>
        [Input("outputCosBucketUrl")]
        public Input<string>? OutputCosBucketUrl { get; set; }

        /// <summary>
        /// The COS bucket directory where the logs are saved. Check below for the rules of the directory name.1. It must be a combination of number, letters, and visible characters. Up to 60 characters are allowed.2. Use a slash (/) to create a subdirectory.3. Consecutive dots (.) and slashes (/) are not allowed. It can not start with a slash (/).
        /// </summary>
        [Input("outputCosKeyPrefix")]
        public Input<string>? OutputCosKeyPrefix { get; set; }

        [Input("tags")]
        private InputList<Inputs.CommandTagGetArgs>? _tags;

        /// <summary>
        /// Tags bound to the command. At most 10 tags are allowed.
        /// </summary>
        public InputList<Inputs.CommandTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.CommandTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Command timeout period. Default value: 60 seconds. Value range: [1, 86400].
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// Command update time.
        /// </summary>
        [Input("updatedTime")]
        public Input<string>? UpdatedTime { get; set; }

        /// <summary>
        /// The username used to execute the command on the CVM or Lighthouse instance.The principle of least privilege is the best practice for permission management. We recommend you execute TAT commands as a general user. By default, the root user is used to execute commands on Linux and the System user is used on Windows.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// Command execution path. The default value is /root for `SHELL` commands and C:/Program Files/qcloudtat_agent/workdir for `POWERSHELL` commands.
        /// </summary>
        [Input("workingDirectory")]
        public Input<string>? WorkingDirectory { get; set; }

        public CommandState()
        {
        }
    }
}
