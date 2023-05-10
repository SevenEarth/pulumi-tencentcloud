// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace TencentCloudIAC.PulumiPackage.Tencentcloud.Monitor
{
    /// <summary>
    /// Provides a resource to create a tke tmpRecordRule
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
    ///         var foo = new Tencentcloud.Monitor.TmpTkeRecordRuleYaml("foo", new Tencentcloud.Monitor.TmpTkeRecordRuleYamlArgs
    ///         {
    ///             Content = @"    apiVersion: monitoring.coreos.com/v1
    ///     kind: PrometheusRule
    ///     metadata:
    ///       name: example-record
    ///     spec:
    ///       groups:
    ///         - name: kube-apiserver.rules
    ///           rules:
    ///             - expr: sum(metrics_test)
    ///               labels:
    ///                 verb: read
    ///               record: 'apiserver_request:burnrate1d'
    /// 
    /// ",
    ///             InstanceId = "",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpTkeRecordRuleYaml:TmpTkeRecordRuleYaml")]
    public partial class TmpTkeRecordRuleYaml : Pulumi.CustomResource
    {
        /// <summary>
        /// An ID identify the cluster, like cls-xxxxxx.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// Contents of record rules in yaml format.
        /// </summary>
        [Output("content")]
        public Output<string> Content { get; private set; } = null!;

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Name of the instance.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Used for the argument, if the configuration comes to the template, the template id.
        /// </summary>
        [Output("templateId")]
        public Output<string> TemplateId { get; private set; } = null!;

        /// <summary>
        /// Last modified time of record rule.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a TmpTkeRecordRuleYaml resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpTkeRecordRuleYaml(string name, TmpTkeRecordRuleYamlArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeRecordRuleYaml:TmpTkeRecordRuleYaml", name, args ?? new TmpTkeRecordRuleYamlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpTkeRecordRuleYaml(string name, Input<string> id, TmpTkeRecordRuleYamlState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeRecordRuleYaml:TmpTkeRecordRuleYaml", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpTkeRecordRuleYaml resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpTkeRecordRuleYaml Get(string name, Input<string> id, TmpTkeRecordRuleYamlState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpTkeRecordRuleYaml(name, id, state, options);
        }
    }

    public sealed class TmpTkeRecordRuleYamlArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Contents of record rules in yaml format.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public TmpTkeRecordRuleYamlArgs()
        {
        }
    }

    public sealed class TmpTkeRecordRuleYamlState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An ID identify the cluster, like cls-xxxxxx.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// Contents of record rules in yaml format.
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// Instance Id.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name of the instance.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Used for the argument, if the configuration comes to the template, the template id.
        /// </summary>
        [Input("templateId")]
        public Input<string>? TemplateId { get; set; }

        /// <summary>
        /// Last modified time of record rule.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public TmpTkeRecordRuleYamlState()
        {
        }
    }
}
