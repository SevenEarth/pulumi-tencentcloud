// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.ApiGateway
{
    /// <summary>
    /// Provides a resource to create a apiGateway plugin_attachment
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var examplePlugin = new Tencentcloud.ApiGateway.Plugin("examplePlugin", new()
    ///     {
    ///         PluginName = "tf-example",
    ///         PluginType = "IPControl",
    ///         PluginData = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["type"] = "white_list",
    ///             ["blocks"] = "1.1.1.1",
    ///         }),
    ///         Description = "desc.",
    ///     });
    /// 
    ///     var exampleService = new Tencentcloud.ApiGateway.Service("exampleService", new()
    ///     {
    ///         ServiceName = "tf_example_service",
    ///         Protocol = "http&amp;https",
    ///         ServiceDesc = "your nice service",
    ///         NetTypes = new[]
    ///         {
    ///             "INNER",
    ///             "OUTER",
    ///         },
    ///         IpVersion = "IPv4",
    ///     });
    /// 
    ///     var exampleApi = new Tencentcloud.ApiGateway.Api("exampleApi", new()
    ///     {
    ///         ServiceId = exampleService.Id,
    ///         ApiName = "tf_example_api",
    ///         ApiDesc = "desc.",
    ///         AuthType = "APP",
    ///         Protocol = "HTTP",
    ///         EnableCors = true,
    ///         RequestConfigPath = "/user/info",
    ///         RequestConfigMethod = "GET",
    ///         RequestParameters = new[]
    ///         {
    ///             new Tencentcloud.ApiGateway.Inputs.ApiRequestParameterArgs
    ///             {
    ///                 Name = "name",
    ///                 Position = "QUERY",
    ///                 Type = "string",
    ///                 Desc = "desc.",
    ///                 DefaultValue = "terraform",
    ///                 Required = true,
    ///             },
    ///         },
    ///         ServiceConfigType = "HTTP",
    ///         ServiceConfigTimeout = 15,
    ///         ServiceConfigUrl = "https://www.qq.com",
    ///         ServiceConfigPath = "/user",
    ///         ServiceConfigMethod = "GET",
    ///         ResponseType = "HTML",
    ///         ResponseSuccessExample = "success",
    ///         ResponseFailExample = "fail",
    ///         ResponseErrorCodes = new[]
    ///         {
    ///             new Tencentcloud.ApiGateway.Inputs.ApiResponseErrorCodeArgs
    ///             {
    ///                 Code = 400,
    ///                 Msg = "system error msg.",
    ///                 Desc = "system error desc.",
    ///                 ConvertedCode = 407,
    ///                 NeedConvert = true,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleServiceRelease = new Tencentcloud.ApiGateway.ServiceRelease("exampleServiceRelease", new()
    ///     {
    ///         ServiceId = exampleApi.ServiceId,
    ///         EnvironmentName = "release",
    ///         ReleaseDesc = "desc.",
    ///     });
    /// 
    ///     var examplePluginAttachment = new Tencentcloud.ApiGateway.PluginAttachment("examplePluginAttachment", new()
    ///     {
    ///         PluginId = examplePlugin.Id,
    ///         ServiceId = exampleServiceRelease.ServiceId,
    ///         ApiId = exampleApi.Id,
    ///         EnvironmentName = "release",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// apiGateway plugin_attachment can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:ApiGateway/pluginAttachment:PluginAttachment example plugin-hnqntalp#service-q3f533ja#release#api-62ud9woa
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:ApiGateway/pluginAttachment:PluginAttachment")]
    public partial class PluginAttachment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Id of API.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// Name of Environment.
        /// </summary>
        [Output("environmentName")]
        public Output<string> EnvironmentName { get; private set; } = null!;

        /// <summary>
        /// Id of Plugin.
        /// </summary>
        [Output("pluginId")]
        public Output<string> PluginId { get; private set; } = null!;

        /// <summary>
        /// Id of Service.
        /// </summary>
        [Output("serviceId")]
        public Output<string> ServiceId { get; private set; } = null!;


        /// <summary>
        /// Create a PluginAttachment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PluginAttachment(string name, PluginAttachmentArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/pluginAttachment:PluginAttachment", name, args ?? new PluginAttachmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PluginAttachment(string name, Input<string> id, PluginAttachmentState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:ApiGateway/pluginAttachment:PluginAttachment", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PluginAttachment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PluginAttachment Get(string name, Input<string> id, PluginAttachmentState? state = null, CustomResourceOptions? options = null)
        {
            return new PluginAttachment(name, id, state, options);
        }
    }

    public sealed class PluginAttachmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of API.
        /// </summary>
        [Input("apiId", required: true)]
        public Input<string> ApiId { get; set; } = null!;

        /// <summary>
        /// Name of Environment.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Id of Plugin.
        /// </summary>
        [Input("pluginId", required: true)]
        public Input<string> PluginId { get; set; } = null!;

        /// <summary>
        /// Id of Service.
        /// </summary>
        [Input("serviceId", required: true)]
        public Input<string> ServiceId { get; set; } = null!;

        public PluginAttachmentArgs()
        {
        }
        public static new PluginAttachmentArgs Empty => new PluginAttachmentArgs();
    }

    public sealed class PluginAttachmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of API.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// Name of Environment.
        /// </summary>
        [Input("environmentName")]
        public Input<string>? EnvironmentName { get; set; }

        /// <summary>
        /// Id of Plugin.
        /// </summary>
        [Input("pluginId")]
        public Input<string>? PluginId { get; set; }

        /// <summary>
        /// Id of Service.
        /// </summary>
        [Input("serviceId")]
        public Input<string>? ServiceId { get; set; }

        public PluginAttachmentState()
        {
        }
        public static new PluginAttachmentState Empty => new PluginAttachmentState();
    }
}
