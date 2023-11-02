-- this should be replaced with a better function to allow parallelism
CREATE OR REPLACE FUNCTION "public"."next_id"()
    RETURNS "pg_catalog"."int8" AS
$BODY$
DECLARE
    unix_utc_now bigint := (SELECT (EXTRACT('epoch' from NOW() at TIME ZONE ('UTC')) * 1000000000));
BEGIN
    RETURN unix_utc_now;
END;
$BODY$
    LANGUAGE plpgsql VOLATILE
                     COST 100;

CREATE EXTENSION hstore;
