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
    /// Provides a resource to create a tmp tke template
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
    ///         var foo = new Tencentcloud.Monitor.TmpTkeTemplate("foo", new Tencentcloud.Monitor.TmpTkeTemplateArgs
    ///         {
    ///             Template = new Tencentcloud.Monitor.Inputs.TmpTkeTemplateTemplateArgs
    ///             {
    ///                 Name = "tf-template",
    ///                 Level = "cluster",
    ///                 Describe = "template",
    ///                 ServiceMonitors = 
    ///                 {
    ///                     new Tencentcloud.Monitor.Inputs.TmpTkeTemplateTemplateServiceMonitorArgs
    ///                     {
    ///                         Name = "tf-ServiceMonitor",
    ///                         Config = @"apiVersion: monitoring.coreos.com/v1
    /// kind: ServiceMonitor
    /// metadata:
    ///   name: example-service-monitor
    ///   namespace: monitoring
    ///   labels:
    ///     k8s-app: example-service
    /// spec:
    ///   selector:
    ///     matchLabels:
    ///       k8s-app: example-service
    ///   namespaceSelector:
    ///     matchNames:
    ///       - default
    ///   endpoints:
    ///   - port: http-metrics
    ///     interval: 30s
    ///     path: /metrics
    ///     scheme: http
    ///     bearerTokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token
    ///     tlsConfig:
    ///       insecureSkipVerify: true
    /// ",
    ///                     },
    ///                 },
    ///                 PodMonitors = 
    ///                 {
    ///                     new Tencentcloud.Monitor.Inputs.TmpTkeTemplateTemplatePodMonitorArgs
    ///                     {
    ///                         Name = "tf-PodMonitors",
    ///                         Config = @"apiVersion: monitoring.coreos.com/v1
    /// kind: PodMonitor
    /// metadata:
    ///   name: example-pod-monitor
    ///   namespace: monitoring
    ///   labels:
    ///     k8s-app: example-pod
    /// spec:
    ///   selector:
    ///     matchLabels:
    ///       k8s-app: example-pod
    ///   namespaceSelector:
    ///     matchNames:
    ///       - default
    ///   podMetricsEndpoints:
    ///   - port: http-metrics
    ///     interval: 30s
    ///     path: /metrics
    ///     scheme: http
    ///     bearerTokenFile: /var/run/secrets/kubernetes.io/serviceaccount/token
    ///     tlsConfig:
    ///       insecureSkipVerify: true
    /// ",
    ///                     },
    ///                     new Tencentcloud.Monitor.Inputs.TmpTkeTemplateTemplatePodMonitorArgs
    ///                     {
    ///                         Name = "tf-RawJobs",
    ///                         Config = @"scrape_configs:
    ///   - job_name: 'example-job'
    ///     scrape_interval: 30s
    ///     static_configs:
    ///       - targets: ['example-service.default.svc.cluster.local:8080']
    ///     metrics_path: /metrics
    ///     scheme: http
    ///     bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
    ///     tls_config:
    ///       insecure_skip_verify: true
    /// ",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TencentcloudResourceType("tencentcloud:Monitor/tmpTkeTemplate:TmpTkeTemplate")]
    public partial class TmpTkeTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Template settings.
        /// </summary>
        [Output("template")]
        public Output<Outputs.TmpTkeTemplateTemplate> Template { get; private set; } = null!;


        /// <summary>
        /// Create a TmpTkeTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TmpTkeTemplate(string name, TmpTkeTemplateArgs args, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeTemplate:TmpTkeTemplate", name, args ?? new TmpTkeTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TmpTkeTemplate(string name, Input<string> id, TmpTkeTemplateState? state = null, CustomResourceOptions? options = null)
            : base("tencentcloud:Monitor/tmpTkeTemplate:TmpTkeTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TmpTkeTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TmpTkeTemplate Get(string name, Input<string> id, TmpTkeTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new TmpTkeTemplate(name, id, state, options);
        }
    }

    public sealed class TmpTkeTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template settings.
        /// </summary>
        [Input("template", required: true)]
        public Input<Inputs.TmpTkeTemplateTemplateArgs> Template { get; set; } = null!;

        public TmpTkeTemplateArgs()
        {
        }
    }

    public sealed class TmpTkeTemplateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Template settings.
        /// </summary>
        [Input("template")]
        public Input<Inputs.TmpTkeTemplateTemplateGetArgs>? Template { get; set; }

        public TmpTkeTemplateState()
        {
        }
    }
}
