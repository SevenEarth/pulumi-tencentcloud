// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Waf
{
    /// <summary>
    /// Provides a resource to create a waf saas instance
    /// 
    /// &gt; **NOTE:** Region only supports `ap-guangzhou` and `ap-seoul`.
    /// 
    /// ## Example Usage
    /// 
    /// ### Create a basic waf premium saas instance
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
    ///     var example = new Tencentcloud.Waf.SaasInstance("example", new()
    ///     {
    ///         GoodsCategory = "premium_saas",
    ///         InstanceName = "tf-example-saas-waf",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Create a complete waf ultimate_saas instance
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
    ///     var example = new Tencentcloud.Waf.SaasInstance("example", new()
    ///     {
    ///         ApiSecurity = 1,
    ///         AutoRenewFlag = 1,
    ///         BotManagement = 1,
    ///         ElasticMode = 1,
    ///         GoodsCategory = "ultimate_saas",
    ///         InstanceName = "tf-example-saas-waf",
    ///         RealRegion = "gz",
    ///         TimeSpan = 1,
    ///         TimeUnit = "m",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Set waf ultimate_saas instance qps limit
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
    ///     var example = new Tencentcloud.Waf.SaasInstance("example", new()
    ///     {
    ///         ApiSecurity = 1,
    ///         AutoRenewFlag = 1,
    ///         BotManagement = 1,
    ///         ElasticMode = 1,
    ///         GoodsCategory = "ultimate_saas",
    ///         InstanceName = "tf-example-saas-waf",
    ///         QpsLimit = 200000,
    ///         RealRegion = "gz",
    ///         TimeSpan = 1,
    ///         TimeUnit = "m",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Waf/saasInstance:SaasInstance")]
    public partial class SaasInstance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to purchase API Security, 1: yes, 0: no. Default is 0.
        /// </summary>
        [Output("apiSecurity")]
        public Output<int?> ApiSecurity { get; private set; } = null!;

        /// <summary>
        /// Auto renew flag, 1: enable, 0: disable.
        /// </summary>
        [Output("autoRenewFlag")]
        public Output<int?> AutoRenewFlag { get; private set; } = null!;

        /// <summary>
        /// waf instance start time.
        /// </summary>
        [Output("beginTime")]
        public Output<string> BeginTime { get; private set; } = null!;

        /// <summary>
        /// Whether to purchase Bot management, 1: yes, 0: no. Default is 0.
        /// </summary>
        [Output("botManagement")]
        public Output<int?> BotManagement { get; private set; } = null!;

        /// <summary>
        /// waf instance edition, clb or saas.
        /// </summary>
        [Output("edition")]
        public Output<string> Edition { get; private set; } = null!;

        /// <summary>
        /// Is elastic billing enabled, 1: enable, 0: disable.
        /// </summary>
        [Output("elasticMode")]
        public Output<int?> ElasticMode { get; private set; } = null!;

        /// <summary>
        /// Billing order parameters. support premium_saas, enterprise_saas, ultimate_saas.
        /// </summary>
        [Output("goodsCategory")]
        public Output<string> GoodsCategory { get; private set; } = null!;

        /// <summary>
        /// waf instance id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Waf instance name.
        /// </summary>
        [Output("instanceName")]
        public Output<string> InstanceName { get; private set; } = null!;

        /// <summary>
        /// QPS Limit, Minimum setting 10000. Only `elastic_mode` is 1, can be set.
        /// </summary>
        [Output("qpsLimit")]
        public Output<int> QpsLimit { get; private set; } = null!;

        /// <summary>
        /// region. If Region is `ap-guangzhou`, support: gz, sh, bj, cd (Means: GuangZhou, ShangHai, BeiJing, ChengDu); If Region is `ap-seoul`, support: hk, sg, th, kr, in, de, ca, use, sao, usw, jkt (Means: HongKong, Singapore, Bandkok, Seoul, Mumbai, Frankfurt, Toronto, Virginia, SaoPaulo, SiliconValley, Jakarta).
        /// </summary>
        [Output("realRegion")]
        public Output<string?> RealRegion { get; private set; } = null!;

        /// <summary>
        /// waf instance status.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// Time interval.
        /// </summary>
        [Output("timeSpan")]
        public Output<int?> TimeSpan { get; private set; } = null!;

        /// <summary>
        /// Time unit, support d, m, y. d: day, m: month, y: year.
        /// </summary>
        [Output("timeUnit")]
        public Output<string?> TimeUnit { get; private set; } = null!;

        /// <summary>
        /// waf instance valid time.
        /// </summary>
        [Output("validTime")]
        public Output<string> ValidTime { get; private set; } = null!;


        /// <summary>
        /// Create a SaasInstance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SaasInstance(string name, SaasInstanceArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/saasInstance:SaasInstance", name, args ?? new SaasInstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SaasInstance(string name, Input<string> id, SaasInstanceState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Waf/saasInstance:SaasInstance", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SaasInstance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SaasInstance Get(string name, Input<string> id, SaasInstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new SaasInstance(name, id, state, options);
        }
    }

    public sealed class SaasInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to purchase API Security, 1: yes, 0: no. Default is 0.
        /// </summary>
        [Input("apiSecurity")]
        public Input<int>? ApiSecurity { get; set; }

        /// <summary>
        /// Auto renew flag, 1: enable, 0: disable.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// Whether to purchase Bot management, 1: yes, 0: no. Default is 0.
        /// </summary>
        [Input("botManagement")]
        public Input<int>? BotManagement { get; set; }

        /// <summary>
        /// Is elastic billing enabled, 1: enable, 0: disable.
        /// </summary>
        [Input("elasticMode")]
        public Input<int>? ElasticMode { get; set; }

        /// <summary>
        /// Billing order parameters. support premium_saas, enterprise_saas, ultimate_saas.
        /// </summary>
        [Input("goodsCategory", required: true)]
        public Input<string> GoodsCategory { get; set; } = null!;

        /// <summary>
        /// Waf instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// QPS Limit, Minimum setting 10000. Only `elastic_mode` is 1, can be set.
        /// </summary>
        [Input("qpsLimit")]
        public Input<int>? QpsLimit { get; set; }

        /// <summary>
        /// region. If Region is `ap-guangzhou`, support: gz, sh, bj, cd (Means: GuangZhou, ShangHai, BeiJing, ChengDu); If Region is `ap-seoul`, support: hk, sg, th, kr, in, de, ca, use, sao, usw, jkt (Means: HongKong, Singapore, Bandkok, Seoul, Mumbai, Frankfurt, Toronto, Virginia, SaoPaulo, SiliconValley, Jakarta).
        /// </summary>
        [Input("realRegion")]
        public Input<string>? RealRegion { get; set; }

        /// <summary>
        /// Time interval.
        /// </summary>
        [Input("timeSpan")]
        public Input<int>? TimeSpan { get; set; }

        /// <summary>
        /// Time unit, support d, m, y. d: day, m: month, y: year.
        /// </summary>
        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        public SaasInstanceArgs()
        {
        }
        public static new SaasInstanceArgs Empty => new SaasInstanceArgs();
    }

    public sealed class SaasInstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to purchase API Security, 1: yes, 0: no. Default is 0.
        /// </summary>
        [Input("apiSecurity")]
        public Input<int>? ApiSecurity { get; set; }

        /// <summary>
        /// Auto renew flag, 1: enable, 0: disable.
        /// </summary>
        [Input("autoRenewFlag")]
        public Input<int>? AutoRenewFlag { get; set; }

        /// <summary>
        /// waf instance start time.
        /// </summary>
        [Input("beginTime")]
        public Input<string>? BeginTime { get; set; }

        /// <summary>
        /// Whether to purchase Bot management, 1: yes, 0: no. Default is 0.
        /// </summary>
        [Input("botManagement")]
        public Input<int>? BotManagement { get; set; }

        /// <summary>
        /// waf instance edition, clb or saas.
        /// </summary>
        [Input("edition")]
        public Input<string>? Edition { get; set; }

        /// <summary>
        /// Is elastic billing enabled, 1: enable, 0: disable.
        /// </summary>
        [Input("elasticMode")]
        public Input<int>? ElasticMode { get; set; }

        /// <summary>
        /// Billing order parameters. support premium_saas, enterprise_saas, ultimate_saas.
        /// </summary>
        [Input("goodsCategory")]
        public Input<string>? GoodsCategory { get; set; }

        /// <summary>
        /// waf instance id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Waf instance name.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// QPS Limit, Minimum setting 10000. Only `elastic_mode` is 1, can be set.
        /// </summary>
        [Input("qpsLimit")]
        public Input<int>? QpsLimit { get; set; }

        /// <summary>
        /// region. If Region is `ap-guangzhou`, support: gz, sh, bj, cd (Means: GuangZhou, ShangHai, BeiJing, ChengDu); If Region is `ap-seoul`, support: hk, sg, th, kr, in, de, ca, use, sao, usw, jkt (Means: HongKong, Singapore, Bandkok, Seoul, Mumbai, Frankfurt, Toronto, Virginia, SaoPaulo, SiliconValley, Jakarta).
        /// </summary>
        [Input("realRegion")]
        public Input<string>? RealRegion { get; set; }

        /// <summary>
        /// waf instance status.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// Time interval.
        /// </summary>
        [Input("timeSpan")]
        public Input<int>? TimeSpan { get; set; }

        /// <summary>
        /// Time unit, support d, m, y. d: day, m: month, y: year.
        /// </summary>
        [Input("timeUnit")]
        public Input<string>? TimeUnit { get; set; }

        /// <summary>
        /// waf instance valid time.
        /// </summary>
        [Input("validTime")]
        public Input<string>? ValidTime { get; set; }

        public SaasInstanceState()
        {
        }
        public static new SaasInstanceState Empty => new SaasInstanceState();
    }
}
