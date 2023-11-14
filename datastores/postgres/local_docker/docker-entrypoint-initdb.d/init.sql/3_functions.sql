CREATE OR REPLACE FUNCTION "public"."slot_to_epoch"("slot" int8)
    RETURNS "pg_catalog"."int8" AS
$BODY$
BEGIN
    IF slot < 32 THEN
        RETURN 0;
    ELSE
        RETURN int8(slot / 32);
    END IF;
END;
$BODY$
    LANGUAGE plpgsql IMMUTABLE
                     COST 100;


CREATE OR REPLACE FUNCTION "public"."slot_is_epoch"("slot" int8)
    RETURNS "pg_catalog"."bool" AS
$BODY$
BEGIN
    IF mod(slot, 32) <= 0 THEN
        RETURN true;
    ELSE
        RETURN false;
    END IF;
END;
$BODY$
    LANGUAGE plpgsql IMMUTABLE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."normalize_balance"("balance" int8)
    RETURNS "pg_catalog"."bool" AS
$BODY$
BEGIN
    RETURN balance - 32000000000;
END;
$BODY$
    LANGUAGE plpgsql IMMUTABLE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."mainnet_head_slot"()
    RETURNS "pg_catalog"."numeric" AS
$BODY$
DECLARE
    unix_utc_now           numeric := (SELECT EXTRACT(epoch FROM NOW() at TIME ZONE ('UTC')));
    unix_time_from_genesis numeric := (unix_utc_now - 1606824023);
    seconds_per_slot       int8    := 12;
BEGIN
    RETURN unix_time_from_genesis / seconds_per_slot;
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."mainnet_finalized_slot"()
    RETURNS "pg_catalog"."int8" AS
$BODY$
DECLARE
    unix_utc_now                       numeric := (SELECT EXTRACT(epoch FROM NOW() at TIME ZONE ('UTC')));
    unix_time_from_genesis             numeric := (unix_utc_now - 1606824023);
    four_epoch_delay_seconds           numeric := 12 * 32 * 4;
    unix_time_from_genesis_behind_slot numeric := unix_time_from_genesis - four_epoch_delay_seconds;
BEGIN
    RETURN unix_time_from_genesis_behind_slot / 12;
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."mainnet_finalized_epoch"()
    RETURNS "pg_catalog"."int8" AS
$BODY$
DECLARE
BEGIN
    RETURN (SELECT mainnet_finalized_slot() / 32);
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."is_mainnet_finalized_epoch"(epoch int8)
    RETURNS "pg_catalog"."bool" AS
$BODY$
DECLARE
BEGIN
    RETURN epoch < (SELECT mainnet_finalized_epoch());
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;

CREATE OR REPLACE FUNCTION "public"."mainnet_head_epoch"()
    RETURNS "pg_catalog"."int8" AS
$BODY$
DECLARE
BEGIN
    RETURN (SELECT mainnet_head_slot() / 32);
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;
