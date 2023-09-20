// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { BucketArgs } from "./bucket";
export type Bucket = import("./bucket").Bucket;
export const Bucket: typeof import("./bucket").Bucket = null as any;
utilities.lazyLoad(exports, ["Bucket"], () => require("./bucket"));

export { OrganizationArgs } from "./organization";
export type Organization = import("./organization").Organization;
export const Organization: typeof import("./organization").Organization = null as any;
utilities.lazyLoad(exports, ["Organization"], () => require("./organization"));

export { PgFunctionArgs } from "./pgFunction";
export type PgFunction = import("./pgFunction").PgFunction;
export const PgFunction: typeof import("./pgFunction").PgFunction = null as any;
utilities.lazyLoad(exports, ["PgFunction"], () => require("./pgFunction"));

export { ProjectArgs } from "./project";
export type Project = import("./project").Project;
export const Project: typeof import("./project").Project = null as any;
utilities.lazyLoad(exports, ["Project"], () => require("./project"));

export { SecretArgs } from "./secret";
export type Secret = import("./secret").Secret;
export const Secret: typeof import("./secret").Secret = null as any;
utilities.lazyLoad(exports, ["Secret"], () => require("./secret"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "supabase:supabase:Bucket":
                return new Bucket(name, <any>undefined, { urn })
            case "supabase:supabase:Organization":
                return new Organization(name, <any>undefined, { urn })
            case "supabase:supabase:PgFunction":
                return new PgFunction(name, <any>undefined, { urn })
            case "supabase:supabase:Project":
                return new Project(name, <any>undefined, { urn })
            case "supabase:supabase:Secret":
                return new Secret(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("supabase", "supabase", _module)
