// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BucketArgs } from "./bucket";
export type Bucket = import("./bucket").Bucket;
export const Bucket: typeof import("./bucket").Bucket = null as any;
utilities.lazyLoad(exports, ["Bucket"], () => require("./bucket"));

export { MigrationArgs } from "./migration";
export type Migration = import("./migration").Migration;
export const Migration: typeof import("./migration").Migration = null as any;
utilities.lazyLoad(exports, ["Migration"], () => require("./migration"));

export { OrganizationArgs } from "./organization";
export type Organization = import("./organization").Organization;
export const Organization: typeof import("./organization").Organization = null as any;
utilities.lazyLoad(exports, ["Organization"], () => require("./organization"));

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "supabase:v0:Bucket":
                return new Bucket(name, <any>undefined, { urn })
            case "supabase:v0:Migration":
                return new Migration(name, <any>undefined, { urn })
            case "supabase:v0:Organization":
                return new Organization(name, <any>undefined, { urn })
            case "supabase:v0:Project":
                return new Project(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("supabase", "v0", _module)
