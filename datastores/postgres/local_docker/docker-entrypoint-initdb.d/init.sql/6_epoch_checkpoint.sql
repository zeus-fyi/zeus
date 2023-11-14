CREATE TABLE "public"."validators_epoch_checkpoint"
(
    "validators_balance_epoch"      serial NOT NULL,
    "validators_active"             int4   NOT NULL CHECK (validators_active >= validators_balances_recorded),
    "validators_balances_recorded"  int4   NOT NULL DEFAULT 0 CHECK (validators_balances_recorded <= validators_active),
    "validators_balances_remaining" int4   NOT NULL GENERATED ALWAYS AS (validators_active - validators_balances_recorded) STORED
)
;
ALTER TABLE "public"."validators_epoch_checkpoint"
    ADD CONSTRAINT "validators_balance_epoch_pkey" PRIMARY KEY ("validators_balance_epoch");

CREATE INDEX amount_not_zero_idx ON validators_epoch_checkpoint ((validators_balances_remaining <> 0)) WHERE validators_balances_remaining <> 0;

CREATE OR REPLACE FUNCTION "public"."validators_active_at_epoch"(epoch int4)
    RETURNS "pg_catalog"."int4" AS
$BODY$
DECLARE
BEGIN
    RETURN (SELECT COUNT(*)
            FROM validators
            WHERE validators.activation_epoch <= epoch AND epoch < validators.exit_epoch);
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."validators_balances_recorded_at_epoch"(vb_epoch int4)
    RETURNS "pg_catalog"."int4" AS
$BODY$
DECLARE
BEGIN
    RETURN (SELECT COUNT(*)
            FROM validator_balances_at_epoch vbe
                     JOIN validators v ON v.index = vbe.validator_index
            WHERE vbe.epoch = vb_epoch
              AND v.activation_epoch <= vb_epoch
              AND vb_epoch < v.exit_epoch);
END
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."update_checkpoint_at_epoch"(checkpoint_epoch int4)
    RETURNS VOID AS
$BODY$
DECLARE
BEGIN
    UPDATE validators_epoch_checkpoint vc
    SET validators_balances_recorded = (SELECT validators_balances_recorded_at_epoch(checkpoint_epoch)),
        validators_active            = (SELECT validators_active_at_epoch(checkpoint_epoch))
    WHERE checkpoint_epoch < (SELECT mainnet_head_epoch())
      AND vc.validators_balance_epoch = checkpoint_epoch;
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;
