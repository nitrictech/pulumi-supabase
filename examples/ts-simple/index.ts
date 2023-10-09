import * as supabase from "@nitric/pulumi-supabase";

const org = new supabase.organizations.Organization("my-org", {
    name: "my-org",
});

const project = new supabase.projects.Project("my-project", {
   organization_id: org.organization_id,
   db_pass: "testing",
   kps_enabled: false,
   plan: "tier_free",
   region: "Oceania (Sydney)",
   cloud: "AWS",
   name: "my-project"
});

// The URL at which the REST API will be served.
export const endpoint = project.project_endpoint;
export const host = project.database_host;
