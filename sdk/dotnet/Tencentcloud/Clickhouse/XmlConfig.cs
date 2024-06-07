// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Clickhouse
{
    /// <summary>
    /// Provides a resource to create a clickhouse xml_config
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tencentcloud = TencentCloudIAC.PulumiPackage.Tencentcloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var xmlConfig = new Tencentcloud.Clickhouse.XmlConfig("xmlConfig", new()
    ///     {
    ///         InstanceId = "cdwch-datuhk3z",
    ///         ModifyConfContext = new Tencentcloud.Clickhouse.Inputs.XmlConfigModifyConfContextArgs
    ///         {
    ///             FileName = "metrika.xml",
    ///             FilePath = "/etc/clickhouse-server",
    ///             NewConfValue = "PD94bWwgdmVyc2lvbj0iMS4wIiBlbmNvZGluZz0iVVRGLTgiPz4KPHlhbmRleD4KICAgIDx6b29rZWVwZXItc2VydmVycz4KICAgIDwvem9va2VlcGVyLXNlcnZlcnM+CjwveWFuZGV4Pgo=",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// clickhouse xml_config can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import tencentcloud:Clickhouse/xmlConfig:XmlConfig xml_config cdwch-datuhk3z#metrika.xml
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Clickhouse/xmlConfig:XmlConfig")]
    public partial class XmlConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Configuration file modification information.
        /// </summary>
        [Output("modifyConfContext")]
        public Output<Outputs.XmlConfigModifyConfContext> ModifyConfContext { get; private set; } = null!;


        /// <summary>
        /// Create a XmlConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public XmlConfig(string name, XmlConfigArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/xmlConfig:XmlConfig", name, args ?? new XmlConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private XmlConfig(string name, Input<string> id, XmlConfigState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Clickhouse/xmlConfig:XmlConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing XmlConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static XmlConfig Get(string name, Input<string> id, XmlConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new XmlConfig(name, id, state, options);
        }
    }

    public sealed class XmlConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// Configuration file modification information.
        /// </summary>
        [Input("modifyConfContext", required: true)]
        public Input<Inputs.XmlConfigModifyConfContextArgs> ModifyConfContext { get; set; } = null!;

        public XmlConfigArgs()
        {
        }
        public static new XmlConfigArgs Empty => new XmlConfigArgs();
    }

    public sealed class XmlConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Configuration file modification information.
        /// </summary>
        [Input("modifyConfContext")]
        public Input<Inputs.XmlConfigModifyConfContextGetArgs>? ModifyConfContext { get; set; }

        public XmlConfigState()
        {
        }
        public static new XmlConfigState Empty => new XmlConfigState();
    }
}
