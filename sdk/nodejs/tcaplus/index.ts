// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { ClusterArgs, ClusterState } from "./cluster";
export type Cluster = import("./cluster").Cluster;
export const Cluster: typeof import("./cluster").Cluster = null as any;
utilities.lazyLoad(exports, ["Cluster"], () => require("./cluster"));

export { GetClustersArgs, GetClustersResult, GetClustersOutputArgs } from "./getClusters";
export const getClusters: typeof import("./getClusters").getClusters = null as any;
export const getClustersOutput: typeof import("./getClusters").getClustersOutput = null as any;
utilities.lazyLoad(exports, ["getClusters","getClustersOutput"], () => require("./getClusters"));

export { GetIdlsArgs, GetIdlsResult, GetIdlsOutputArgs } from "./getIdls";
export const getIdls: typeof import("./getIdls").getIdls = null as any;
export const getIdlsOutput: typeof import("./getIdls").getIdlsOutput = null as any;
utilities.lazyLoad(exports, ["getIdls","getIdlsOutput"], () => require("./getIdls"));

export { GetTablegroupsArgs, GetTablegroupsResult, GetTablegroupsOutputArgs } from "./getTablegroups";
export const getTablegroups: typeof import("./getTablegroups").getTablegroups = null as any;
export const getTablegroupsOutput: typeof import("./getTablegroups").getTablegroupsOutput = null as any;
utilities.lazyLoad(exports, ["getTablegroups","getTablegroupsOutput"], () => require("./getTablegroups"));

export { GetTablesArgs, GetTablesResult, GetTablesOutputArgs } from "./getTables";
export const getTables: typeof import("./getTables").getTables = null as any;
export const getTablesOutput: typeof import("./getTables").getTablesOutput = null as any;
utilities.lazyLoad(exports, ["getTables","getTablesOutput"], () => require("./getTables"));

export { IdlArgs, IdlState } from "./idl";
export type Idl = import("./idl").Idl;
export const Idl: typeof import("./idl").Idl = null as any;
utilities.lazyLoad(exports, ["Idl"], () => require("./idl"));

export { TableArgs, TableState } from "./table";
export type Table = import("./table").Table;
export const Table: typeof import("./table").Table = null as any;
utilities.lazyLoad(exports, ["Table"], () => require("./table"));

export { TablegroupArgs, TablegroupState } from "./tablegroup";
export type Tablegroup = import("./tablegroup").Tablegroup;
export const Tablegroup: typeof import("./tablegroup").Tablegroup = null as any;
utilities.lazyLoad(exports, ["Tablegroup"], () => require("./tablegroup"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tencentcloud:Tcaplus/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "tencentcloud:Tcaplus/idl:Idl":
                return new Idl(name, <any>undefined, { urn })
            case "tencentcloud:Tcaplus/table:Table":
                return new Table(name, <any>undefined, { urn })
            case "tencentcloud:Tcaplus/tablegroup:Tablegroup":
                return new Tablegroup(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tencentcloud", "Tcaplus/cluster", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tcaplus/idl", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tcaplus/table", _module)
pulumi.runtime.registerResourceModule("tencentcloud", "Tcaplus/tablegroup", _module)
