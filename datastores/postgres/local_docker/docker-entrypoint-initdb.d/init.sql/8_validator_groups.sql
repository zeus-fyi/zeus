CREATE TABLE "public"."validators_service_org_groups"
(
    "group_name"          text NOT NULL,
    "org_id"              int8 NOT NULL REFERENCES orgs (org_id),
    "pubkey"              text NOT NULL CHECK (LENGTH(pubkey) = 98),
    "protocol_network_id" int8 NOT NULL REFERENCES protocol_networks (protocol_network_id) DEFAULT 1,
    "fee_recipient"       text NOT NULL,
    "enabled"             bool NOT NULL                                                    DEFAULT false,
    "mev_enabled"         bool NOT NULL                                                    DEFAULT true,
    "service_url"         text NOT NULL
);

ALTER TABLE "public"."validators_service_org_groups"
    ADD CONSTRAINT "validators_org_group_pubkey_org_uniq" PRIMARY KEY ("group_name", "pubkey");
ALTER TABLE "public"."validators_service_org_groups"
    ADD CONSTRAINT "validators_org_group_validator_pubkey_uniq" UNIQUE ("pubkey");
ALTER TABLE "public"."validators_service_org_groups"
    ADD CONSTRAINT "validators_org_group_validator_pubkey_network_uniq" UNIQUE ("pubkey", "protocol_network_id");
CREATE INDEX "org_group_index" ON "public"."validators_service_org_groups" ("group_name", "org_id", "protocol_network_id");

CREATE TABLE "public"."validators_service_org_groups_cloud_ctx_ns"
(
    "cloud_ctx_ns_id"         int8 NOT NULL REFERENCES topologies_org_cloud_ctx_ns (cloud_ctx_ns_id),
    "pubkey"                  text NOT NULL CHECK (LENGTH(pubkey) = 98) NOT NULL REFERENCES validators_service_org_groups (pubkey),
    "validator_client_number" int  NOT NULL DEFAULT 0
);

ALTER TABLE "public"."validators_service_org_groups_cloud_ctx_ns"
    ADD CONSTRAINT "validators_service_org_groups_cloud_ctx_ns_pk" PRIMARY KEY ("pubkey", "cloud_ctx_ns_id");
ALTER TABLE "public"."validators_service_org_groups_cloud_ctx_ns"
    ADD CONSTRAINT "validators_service_org_groups_cloud_ctx_ns_pubkey_uniq" UNIQUE ("pubkey");

