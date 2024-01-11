// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a dlc dataEngine
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tencentcloud from "@pulumi/tencentcloud";
 *
 * const dataEngine = new tencentcloud.Dlc.DataEngine("data_engine", {
 *     autoResume: false,
 *     autoSuspend: false,
 *     cidrBlock: "10.255.0.0/16",
 *     clusterType: "spark_cu",
 *     crontabResumeSuspend: 0,
 *     dataEngineName: "testSpark",
 *     defaultDataEngine: false,
 *     engineExecType: "BATCH",
 *     engineType: "spark",
 *     maxClusters: 1,
 *     message: "test spark1",
 *     minClusters: 1,
 *     mode: 1,
 *     payMode: 0,
 *     size: 16,
 *     timeSpan: 1,
 *     timeUnit: "h",
 * });
 * ```
 *
 * ## Import
 *
 * dlc data_engine can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import tencentcloud:Dlc/dataEngine:DataEngine data_engine data_engine_id
 * ```
 */
export class DataEngine extends pulumi.CustomResource {
    /**
     * Get an existing DataEngine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataEngineState, opts?: pulumi.CustomResourceOptions): DataEngine {
        return new DataEngine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tencentcloud:Dlc/dataEngine:DataEngine';

    /**
     * Returns true if the given object is an instance of DataEngine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataEngine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataEngine.__pulumiType;
    }

    /**
     * Engine auto renew, only support 0: Default, 1: AutoRenewON, 2: AutoRenewOFF.
     */
    public readonly autoRenew!: pulumi.Output<number | undefined>;
    /**
     * Whether to automatically start the cluster, prepay not support.
     */
    public readonly autoResume!: pulumi.Output<boolean>;
    /**
     * Whether to automatically suspend the cluster, prepay not support.
     */
    public readonly autoSuspend!: pulumi.Output<boolean | undefined>;
    /**
     * Cluster automatic suspension time, default 10 minutes.
     */
    public readonly autoSuspendTime!: pulumi.Output<number>;
    /**
     * Engine VPC network segment, just like 192.0.2.1/24.
     */
    public readonly cidrBlock!: pulumi.Output<string | undefined>;
    /**
     * Engine cluster type, only support: spark_cu/presto_cu.
     */
    public readonly clusterType!: pulumi.Output<string>;
    /**
     * Engine crontab resume or suspend strategy, only support: 0: Wait(default), 1: Kill.
     */
    public readonly crontabResumeSuspend!: pulumi.Output<number | undefined>;
    /**
     * Engine auto suspend strategy, when AutoSuspend is true, CrontabResumeSuspend must stop.
     */
    public readonly crontabResumeSuspendStrategy!: pulumi.Output<outputs.Dlc.DataEngineCrontabResumeSuspendStrategy>;
    /**
     * Cluster advanced configuration.
     */
    public readonly dataEngineConfigPairs!: pulumi.Output<outputs.Dlc.DataEngineDataEngineConfigPair[] | undefined>;
    /**
     * Engine name.
     */
    public readonly dataEngineName!: pulumi.Output<string>;
    /**
     * Whether it is the default virtual cluster.
     */
    public readonly defaultDataEngine!: pulumi.Output<boolean | undefined>;
    /**
     * For spark Batch ExecType, yearly and monthly cluster elastic limit.
     */
    public readonly elasticLimit!: pulumi.Output<number | undefined>;
    /**
     * For spark Batch ExecType, yearly and monthly cluster whether to enable elasticity.
     */
    public readonly elasticSwitch!: pulumi.Output<boolean | undefined>;
    /**
     * Engine exec type, only support SQL(default) or BATCH.
     */
    public readonly engineExecType!: pulumi.Output<string | undefined>;
    /**
     * Engine type, only support: spark/presto.
     */
    public readonly engineType!: pulumi.Output<string>;
    /**
     * Cluster image version name. Such as SuperSQL-P 1.1; SuperSQL-S 3.2, etc., do not upload, and create a cluster with the latest mirror version by default.
     */
    public readonly imageVersionName!: pulumi.Output<string>;
    /**
     * Primary cluster name, specified when creating a disaster recovery cluster.
     */
    public readonly mainClusterName!: pulumi.Output<string | undefined>;
    /**
     * Engine max cluster size, MaxClusters less than or equal to 10 and MaxClusters bigger than MinClusters.
     */
    public readonly maxClusters!: pulumi.Output<number | undefined>;
    /**
     * Maximum number of concurrent tasks in a single cluster, default 5.
     */
    public readonly maxConcurrency!: pulumi.Output<number>;
    /**
     * Engine description information.
     */
    public readonly message!: pulumi.Output<string | undefined>;
    /**
     * Engine min size, greater than or equal to 1 and MaxClusters bigger than MinClusters.
     */
    public readonly minClusters!: pulumi.Output<number | undefined>;
    /**
     * Engine mode, only support 1: ByAmount, 2: YearlyAndMonthly.
     */
    public readonly mode!: pulumi.Output<number>;
    /**
     * Engine pay mode type, only support 0: postPay, 1: prePay(default).
     */
    public readonly payMode!: pulumi.Output<number | undefined>;
    /**
     * Engine resource type not match, only support: Standard_CU/Memory_CU(only BATCH ExecType).
     */
    public readonly resourceType!: pulumi.Output<string>;
    /**
     * For spark Batch ExecType, cluster session resource configuration template.
     */
    public readonly sessionResourceTemplate!: pulumi.Output<outputs.Dlc.DataEngineSessionResourceTemplate>;
    /**
     * Cluster size. Required when updating.
     */
    public readonly size!: pulumi.Output<number | undefined>;
    /**
     * Engine TimeSpan, prePay: minimum of 1, representing one month of purchasing resources, with a maximum of 120, default 3600, postPay: fixed fee of 3600.
     */
    public readonly timeSpan!: pulumi.Output<number | undefined>;
    /**
     * Engine TimeUnit, prePay: use m(default), postPay: use h.
     */
    public readonly timeUnit!: pulumi.Output<string | undefined>;
    /**
     * Tolerable queuing time, default 0. scaling may be triggered when tasks are queued for longer than the tolerable time. if this parameter is 0, it means that capacity expansion may be triggered immediately once a task is queued.
     */
    public readonly tolerableQueueTime!: pulumi.Output<number | undefined>;

    /**
     * Create a DataEngine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataEngineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataEngineArgs | DataEngineState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataEngineState | undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoResume"] = state ? state.autoResume : undefined;
            resourceInputs["autoSuspend"] = state ? state.autoSuspend : undefined;
            resourceInputs["autoSuspendTime"] = state ? state.autoSuspendTime : undefined;
            resourceInputs["cidrBlock"] = state ? state.cidrBlock : undefined;
            resourceInputs["clusterType"] = state ? state.clusterType : undefined;
            resourceInputs["crontabResumeSuspend"] = state ? state.crontabResumeSuspend : undefined;
            resourceInputs["crontabResumeSuspendStrategy"] = state ? state.crontabResumeSuspendStrategy : undefined;
            resourceInputs["dataEngineConfigPairs"] = state ? state.dataEngineConfigPairs : undefined;
            resourceInputs["dataEngineName"] = state ? state.dataEngineName : undefined;
            resourceInputs["defaultDataEngine"] = state ? state.defaultDataEngine : undefined;
            resourceInputs["elasticLimit"] = state ? state.elasticLimit : undefined;
            resourceInputs["elasticSwitch"] = state ? state.elasticSwitch : undefined;
            resourceInputs["engineExecType"] = state ? state.engineExecType : undefined;
            resourceInputs["engineType"] = state ? state.engineType : undefined;
            resourceInputs["imageVersionName"] = state ? state.imageVersionName : undefined;
            resourceInputs["mainClusterName"] = state ? state.mainClusterName : undefined;
            resourceInputs["maxClusters"] = state ? state.maxClusters : undefined;
            resourceInputs["maxConcurrency"] = state ? state.maxConcurrency : undefined;
            resourceInputs["message"] = state ? state.message : undefined;
            resourceInputs["minClusters"] = state ? state.minClusters : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["payMode"] = state ? state.payMode : undefined;
            resourceInputs["resourceType"] = state ? state.resourceType : undefined;
            resourceInputs["sessionResourceTemplate"] = state ? state.sessionResourceTemplate : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["timeSpan"] = state ? state.timeSpan : undefined;
            resourceInputs["timeUnit"] = state ? state.timeUnit : undefined;
            resourceInputs["tolerableQueueTime"] = state ? state.tolerableQueueTime : undefined;
        } else {
            const args = argsOrState as DataEngineArgs | undefined;
            if ((!args || args.autoResume === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoResume'");
            }
            if ((!args || args.clusterType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterType'");
            }
            if ((!args || args.dataEngineName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataEngineName'");
            }
            if ((!args || args.engineType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'engineType'");
            }
            if ((!args || args.mode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mode'");
            }
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoResume"] = args ? args.autoResume : undefined;
            resourceInputs["autoSuspend"] = args ? args.autoSuspend : undefined;
            resourceInputs["autoSuspendTime"] = args ? args.autoSuspendTime : undefined;
            resourceInputs["cidrBlock"] = args ? args.cidrBlock : undefined;
            resourceInputs["clusterType"] = args ? args.clusterType : undefined;
            resourceInputs["crontabResumeSuspend"] = args ? args.crontabResumeSuspend : undefined;
            resourceInputs["crontabResumeSuspendStrategy"] = args ? args.crontabResumeSuspendStrategy : undefined;
            resourceInputs["dataEngineConfigPairs"] = args ? args.dataEngineConfigPairs : undefined;
            resourceInputs["dataEngineName"] = args ? args.dataEngineName : undefined;
            resourceInputs["defaultDataEngine"] = args ? args.defaultDataEngine : undefined;
            resourceInputs["elasticLimit"] = args ? args.elasticLimit : undefined;
            resourceInputs["elasticSwitch"] = args ? args.elasticSwitch : undefined;
            resourceInputs["engineExecType"] = args ? args.engineExecType : undefined;
            resourceInputs["engineType"] = args ? args.engineType : undefined;
            resourceInputs["imageVersionName"] = args ? args.imageVersionName : undefined;
            resourceInputs["mainClusterName"] = args ? args.mainClusterName : undefined;
            resourceInputs["maxClusters"] = args ? args.maxClusters : undefined;
            resourceInputs["maxConcurrency"] = args ? args.maxConcurrency : undefined;
            resourceInputs["message"] = args ? args.message : undefined;
            resourceInputs["minClusters"] = args ? args.minClusters : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["payMode"] = args ? args.payMode : undefined;
            resourceInputs["resourceType"] = args ? args.resourceType : undefined;
            resourceInputs["sessionResourceTemplate"] = args ? args.sessionResourceTemplate : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
            resourceInputs["timeSpan"] = args ? args.timeSpan : undefined;
            resourceInputs["timeUnit"] = args ? args.timeUnit : undefined;
            resourceInputs["tolerableQueueTime"] = args ? args.tolerableQueueTime : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataEngine.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataEngine resources.
 */
export interface DataEngineState {
    /**
     * Engine auto renew, only support 0: Default, 1: AutoRenewON, 2: AutoRenewOFF.
     */
    autoRenew?: pulumi.Input<number>;
    /**
     * Whether to automatically start the cluster, prepay not support.
     */
    autoResume?: pulumi.Input<boolean>;
    /**
     * Whether to automatically suspend the cluster, prepay not support.
     */
    autoSuspend?: pulumi.Input<boolean>;
    /**
     * Cluster automatic suspension time, default 10 minutes.
     */
    autoSuspendTime?: pulumi.Input<number>;
    /**
     * Engine VPC network segment, just like 192.0.2.1/24.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * Engine cluster type, only support: spark_cu/presto_cu.
     */
    clusterType?: pulumi.Input<string>;
    /**
     * Engine crontab resume or suspend strategy, only support: 0: Wait(default), 1: Kill.
     */
    crontabResumeSuspend?: pulumi.Input<number>;
    /**
     * Engine auto suspend strategy, when AutoSuspend is true, CrontabResumeSuspend must stop.
     */
    crontabResumeSuspendStrategy?: pulumi.Input<inputs.Dlc.DataEngineCrontabResumeSuspendStrategy>;
    /**
     * Cluster advanced configuration.
     */
    dataEngineConfigPairs?: pulumi.Input<pulumi.Input<inputs.Dlc.DataEngineDataEngineConfigPair>[]>;
    /**
     * Engine name.
     */
    dataEngineName?: pulumi.Input<string>;
    /**
     * Whether it is the default virtual cluster.
     */
    defaultDataEngine?: pulumi.Input<boolean>;
    /**
     * For spark Batch ExecType, yearly and monthly cluster elastic limit.
     */
    elasticLimit?: pulumi.Input<number>;
    /**
     * For spark Batch ExecType, yearly and monthly cluster whether to enable elasticity.
     */
    elasticSwitch?: pulumi.Input<boolean>;
    /**
     * Engine exec type, only support SQL(default) or BATCH.
     */
    engineExecType?: pulumi.Input<string>;
    /**
     * Engine type, only support: spark/presto.
     */
    engineType?: pulumi.Input<string>;
    /**
     * Cluster image version name. Such as SuperSQL-P 1.1; SuperSQL-S 3.2, etc., do not upload, and create a cluster with the latest mirror version by default.
     */
    imageVersionName?: pulumi.Input<string>;
    /**
     * Primary cluster name, specified when creating a disaster recovery cluster.
     */
    mainClusterName?: pulumi.Input<string>;
    /**
     * Engine max cluster size, MaxClusters less than or equal to 10 and MaxClusters bigger than MinClusters.
     */
    maxClusters?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent tasks in a single cluster, default 5.
     */
    maxConcurrency?: pulumi.Input<number>;
    /**
     * Engine description information.
     */
    message?: pulumi.Input<string>;
    /**
     * Engine min size, greater than or equal to 1 and MaxClusters bigger than MinClusters.
     */
    minClusters?: pulumi.Input<number>;
    /**
     * Engine mode, only support 1: ByAmount, 2: YearlyAndMonthly.
     */
    mode?: pulumi.Input<number>;
    /**
     * Engine pay mode type, only support 0: postPay, 1: prePay(default).
     */
    payMode?: pulumi.Input<number>;
    /**
     * Engine resource type not match, only support: Standard_CU/Memory_CU(only BATCH ExecType).
     */
    resourceType?: pulumi.Input<string>;
    /**
     * For spark Batch ExecType, cluster session resource configuration template.
     */
    sessionResourceTemplate?: pulumi.Input<inputs.Dlc.DataEngineSessionResourceTemplate>;
    /**
     * Cluster size. Required when updating.
     */
    size?: pulumi.Input<number>;
    /**
     * Engine TimeSpan, prePay: minimum of 1, representing one month of purchasing resources, with a maximum of 120, default 3600, postPay: fixed fee of 3600.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * Engine TimeUnit, prePay: use m(default), postPay: use h.
     */
    timeUnit?: pulumi.Input<string>;
    /**
     * Tolerable queuing time, default 0. scaling may be triggered when tasks are queued for longer than the tolerable time. if this parameter is 0, it means that capacity expansion may be triggered immediately once a task is queued.
     */
    tolerableQueueTime?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DataEngine resource.
 */
export interface DataEngineArgs {
    /**
     * Engine auto renew, only support 0: Default, 1: AutoRenewON, 2: AutoRenewOFF.
     */
    autoRenew?: pulumi.Input<number>;
    /**
     * Whether to automatically start the cluster, prepay not support.
     */
    autoResume: pulumi.Input<boolean>;
    /**
     * Whether to automatically suspend the cluster, prepay not support.
     */
    autoSuspend?: pulumi.Input<boolean>;
    /**
     * Cluster automatic suspension time, default 10 minutes.
     */
    autoSuspendTime?: pulumi.Input<number>;
    /**
     * Engine VPC network segment, just like 192.0.2.1/24.
     */
    cidrBlock?: pulumi.Input<string>;
    /**
     * Engine cluster type, only support: spark_cu/presto_cu.
     */
    clusterType: pulumi.Input<string>;
    /**
     * Engine crontab resume or suspend strategy, only support: 0: Wait(default), 1: Kill.
     */
    crontabResumeSuspend?: pulumi.Input<number>;
    /**
     * Engine auto suspend strategy, when AutoSuspend is true, CrontabResumeSuspend must stop.
     */
    crontabResumeSuspendStrategy?: pulumi.Input<inputs.Dlc.DataEngineCrontabResumeSuspendStrategy>;
    /**
     * Cluster advanced configuration.
     */
    dataEngineConfigPairs?: pulumi.Input<pulumi.Input<inputs.Dlc.DataEngineDataEngineConfigPair>[]>;
    /**
     * Engine name.
     */
    dataEngineName: pulumi.Input<string>;
    /**
     * Whether it is the default virtual cluster.
     */
    defaultDataEngine?: pulumi.Input<boolean>;
    /**
     * For spark Batch ExecType, yearly and monthly cluster elastic limit.
     */
    elasticLimit?: pulumi.Input<number>;
    /**
     * For spark Batch ExecType, yearly and monthly cluster whether to enable elasticity.
     */
    elasticSwitch?: pulumi.Input<boolean>;
    /**
     * Engine exec type, only support SQL(default) or BATCH.
     */
    engineExecType?: pulumi.Input<string>;
    /**
     * Engine type, only support: spark/presto.
     */
    engineType: pulumi.Input<string>;
    /**
     * Cluster image version name. Such as SuperSQL-P 1.1; SuperSQL-S 3.2, etc., do not upload, and create a cluster with the latest mirror version by default.
     */
    imageVersionName?: pulumi.Input<string>;
    /**
     * Primary cluster name, specified when creating a disaster recovery cluster.
     */
    mainClusterName?: pulumi.Input<string>;
    /**
     * Engine max cluster size, MaxClusters less than or equal to 10 and MaxClusters bigger than MinClusters.
     */
    maxClusters?: pulumi.Input<number>;
    /**
     * Maximum number of concurrent tasks in a single cluster, default 5.
     */
    maxConcurrency?: pulumi.Input<number>;
    /**
     * Engine description information.
     */
    message?: pulumi.Input<string>;
    /**
     * Engine min size, greater than or equal to 1 and MaxClusters bigger than MinClusters.
     */
    minClusters?: pulumi.Input<number>;
    /**
     * Engine mode, only support 1: ByAmount, 2: YearlyAndMonthly.
     */
    mode: pulumi.Input<number>;
    /**
     * Engine pay mode type, only support 0: postPay, 1: prePay(default).
     */
    payMode?: pulumi.Input<number>;
    /**
     * Engine resource type not match, only support: Standard_CU/Memory_CU(only BATCH ExecType).
     */
    resourceType?: pulumi.Input<string>;
    /**
     * For spark Batch ExecType, cluster session resource configuration template.
     */
    sessionResourceTemplate?: pulumi.Input<inputs.Dlc.DataEngineSessionResourceTemplate>;
    /**
     * Cluster size. Required when updating.
     */
    size?: pulumi.Input<number>;
    /**
     * Engine TimeSpan, prePay: minimum of 1, representing one month of purchasing resources, with a maximum of 120, default 3600, postPay: fixed fee of 3600.
     */
    timeSpan?: pulumi.Input<number>;
    /**
     * Engine TimeUnit, prePay: use m(default), postPay: use h.
     */
    timeUnit?: pulumi.Input<string>;
    /**
     * Tolerable queuing time, default 0. scaling may be triggered when tasks are queued for longer than the tolerable time. if this parameter is 0, it means that capacity expansion may be triggered immediately once a task is queued.
     */
    tolerableQueueTime?: pulumi.Input<number>;
}