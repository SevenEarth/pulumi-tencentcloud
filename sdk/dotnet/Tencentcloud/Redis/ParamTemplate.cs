// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Redis
{
    /// <summary>
    /// Provides a resource to create a redis parameter template
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
    ///         var paramTemplate = new Tencentcloud.Redis.ParamTemplate("paramTemplate", new Tencentcloud.Redis.ParamTemplateArgs
    ///         {
    ///             Description = "This is an example redis param template.",
    ///             ParamsOverrides = 
    ///             {
    ///                 new Tencentcloud.Redis.Inputs.ParamTemplateParamsOverrideArgs
    ///                 {
    ///                     Key = "timeout",
    ///                     Value = "7200",
    ///                 },
    ///             },
    ///             ProductType = 6,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Copy from another template
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Tencentcloud.Redis.ParamTemplate("foo", new Tencentcloud.Redis.ParamTemplateArgs
    ///         {
    ///             Description = "This is an example redis param template.",
    ///             ProductType = 6,
    ///             ParamsOverrides = 
    ///             {
    ///                 new Tencentcloud.Redis.Inputs.ParamTemplateParamsOverrideArgs
    ///                 {
    ///                     Key = "timeout",
    ///                     Value = "7200",
    ///                 },
    ///             },
    ///         });
    ///         var paramTemplate = new Tencentcloud.Redis.ParamTemplate("paramTemplate", new Tencentcloud.Redis.ParamTemplateArgs
    ///         {
    ///             Description = "This is an copied redis param template from tf-template.",
    ///             TemplateId = foo.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// redis param_template can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import tencentcloud:Redis/paramTemplate:ParamTemplate param_template param_template_id
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Redis/paramTemplate:ParamTemplate")]
    public partial class ParamTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Parameter template description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Parameter template name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Readonly full parameter list details.
        /// </summary>
        [Output("paramDetails")]
        public Output<ImmutableArray<Outputs.ParamTemplateParamDetail>> ParamDetails { get; private set; } = null!;

        /// <summary>
        /// Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        /// </summary>
        [Output("paramsOverrides")]
        public Output<ImmutableArray<Outputs.ParamTemplateParamsOverride>> ParamsOverrides { get; private set; } = null!;

        /// <summary>
        /// Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        /// </summary>
        [Output("productType")]
        public Output<int?> ProductType { get; private set; } = null!;

        /// <summary>
        /// Specify which existed template import from.
        /// </summary>
        [Output("templateId")]
        public Output<string?> TemplateId { get; private set; } = null!;


        /// <summary>
        /// Create a ParamTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ParamTemplate(string name, ParamTemplateArgs? args = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/paramTemplate:ParamTemplate", name, args ?? new ParamTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ParamTemplate(string name, Input<string> id, ParamTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Redis/paramTemplate:ParamTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ParamTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ParamTemplate Get(string name, Input<string> id, ParamTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new ParamTemplate(name, id, state, options);
        }
    }

    public sealed class ParamTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter template description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Parameter template name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("paramsOverrides")]
        private InputList<Inputs.ParamTemplateParamsOverrideArgs>? _paramsOverrides;

        /// <summary>
        /// Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        /// </summary>
        public InputList<Inputs.ParamTemplateParamsOverrideArgs> ParamsOverrides
        {
            get => _paramsOverrides ?? (_paramsOverrides = new InputList<Inputs.ParamTemplateParamsOverrideArgs>());
            set => _paramsOverrides = value;
        }

        /// <summary>
        /// Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        /// </summary>
        [Input("productType")]
        public Input<int>? ProductType { get; set; }

        /// <summary>
        /// Specify which existed template import from.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public ParamTemplateArgs()
        {
        }
    }

    public sealed class ParamTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Parameter template description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Parameter template name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("paramDetails")]
        private InputList<Inputs.ParamTemplateParamDetailGetArgs>? _paramDetails;

        /// <summary>
        /// Readonly full parameter list details.
        /// </summary>
        public InputList<Inputs.ParamTemplateParamDetailGetArgs> ParamDetails
        {
            get => _paramDetails ?? (_paramDetails = new InputList<Inputs.ParamTemplateParamDetailGetArgs>());
            set => _paramDetails = value;
        }

        [Input("paramsOverrides")]
        private InputList<Inputs.ParamTemplateParamsOverrideGetArgs>? _paramsOverrides;

        /// <summary>
        /// Specify override parameter list, NOTE: Do not remove override params once set, removing will not take effects to current value.
        /// </summary>
        public InputList<Inputs.ParamTemplateParamsOverrideGetArgs> ParamsOverrides
        {
            get => _paramsOverrides ?? (_paramsOverrides = new InputList<Inputs.ParamTemplateParamsOverrideGetArgs>());
            set => _paramsOverrides = value;
        }

        /// <summary>
        /// Specify product type. Valid values: 1 (Redis 2.8 Memory Edition in cluster architecture), 2 (Redis 2.8 Memory Edition in standard architecture), 3 (CKV 3.2 Memory Edition in standard architecture), 4 (CKV 3.2 Memory Edition in cluster architecture), 5 (Redis 2.8 Memory Edition in standalone architecture), 6 (Redis 4.0 Memory Edition in standard architecture), 7 (Redis 4.0 Memory Edition in cluster architecture), 8 (Redis 5.0 Memory Edition in standard architecture), 9 (Redis 5.0 Memory Edition in cluster architecture). If `template_id` is specified, this parameter can be left blank; otherwise, it is required.
        /// </summary>
        [Input("productType")]
        public Input<int>? ProductType { get; set; }

        /// <summary>
        /// Specify which existed template import from.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        public ParamTemplateState()
        {
        }
    }
}
