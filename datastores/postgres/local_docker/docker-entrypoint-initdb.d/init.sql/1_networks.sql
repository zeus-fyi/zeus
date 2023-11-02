CREATE TABLE "public"."protocol_networks"
(
    "protocol_network_id" bigint NOT NULL DEFAULT next_id(),
    "protocol_name"       text   NOT NULL,
    "network_name"        text   NOT NULL
);
ALTER TABLE "public"."protocol_networks"
    ADD CONSTRAINT "protocol_networks_pk" PRIMARY KEY ("protocol_network_id");
ALTER TABLE "public"."protocol_networks"
    ADD CONSTRAINT "protocol_networks_unique" UNIQUE ("protocol_name", "network_name");

BEGIN;
INSERT INTO "public"."protocol_networks"
VALUES (1, 'ethereum', 'mainnet');
INSERT INTO "public"."protocol_networks"
VALUES (5, 'ethereum', 'goerli');
INSERT INTO "public"."protocol_networks"
VALUES (11155111, 'ethereum', 'sepolia');
INSERT INTO "public"."protocol_networks"
VALUES (1673748447294772000, 'ethereum', 'ephemery');
COMMIT;
