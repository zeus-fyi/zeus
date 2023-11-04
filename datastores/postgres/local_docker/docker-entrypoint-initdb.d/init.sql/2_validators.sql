-- ----------------------------
-- Table structure for validators
-- ----------------------------
CREATE TABLE "public"."validators"
(
    "index"                        int4                                                                            NOT NULL,
    "protocol_network_id"          int8                                                                            NOT NULL REFERENCES protocol_networks (protocol_network_id) DEFAULT 1,
    "balance"                      int8,
    "effective_balance"            int8,
    "activation_eligibility_epoch" int8 CHECK (activation_eligibility_epoch <= activation_epoch)                   NOT NULL                                                    DEFAULT (9223372036854775807),
    "activation_epoch"             int8 CHECK (activation_epoch >= activation_eligibility_epoch)                   NOT NULL                                                    DEFAULT (9223372036854775807),
    "exit_epoch"                   int8                                                                            NOT NULL                                                    DEFAULT (9223372036854775807),
    "withdrawable_epoch"           int8                                                                            NOT NULL                                                    DEFAULT (9223372036854775807),
    "updated_at"                   timestamptz                                                                     NOT NULL                                                    DEFAULT NOW(),
    "slashed"                      bool                                                                            NOT NULL                                                    DEFAULT false,
    "pubkey"                       text                                                                            NOT NULL CHECK (LENGTH(pubkey) = 98),
    "status"                       text CHECK (status IN ('pending', 'active', 'exited', 'withdrawal', 'unknown')) NOT NULL                                                    DEFAULT 'unknown',
    "substatus"                    text CHECK (substatus IN ('pending_initialized', 'pending_queued', 'active_ongoing',
                                                             'active_exiting', 'active_slashed', 'exited_unslashed',
                                                             'exited_slashed', 'withdrawal_possible', 'withdrawal_done',
                                                             'unknown'))                                                                                                       DEFAULT 'unknown',
    "withdrawal_credentials"       text
)
;

-- ----------------------------
-- Indexes structure for table validators
-- ----------------------------
CREATE UNIQUE INDEX "pubkey_index" ON "public"."validators" USING btree (
                                                                         "pubkey" COLLATE "pg_catalog"."default"
                                                                         "pg_catalog"."bpchar_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table validators
-- ----------------------------
ALTER TABLE "public"."validators"
    ADD CONSTRAINT "validators_pkey" PRIMARY KEY ("index");
CREATE INDEX "last_updated_at_index" ON "public"."validators" (updated_at ASC);

ALTER TABLE "public"."validators"
    ADD CONSTRAINT "pubkey_index_uniq" UNIQUE ("pubkey", "index");
ALTER TABLE "public"."validators"
    ADD CONSTRAINT "pubkey_network_uniq" UNIQUE ("pubkey", "protocol_network_id");
