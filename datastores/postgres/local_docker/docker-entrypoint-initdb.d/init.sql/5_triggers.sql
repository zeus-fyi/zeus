CREATE OR REPLACE FUNCTION update_or_insert_validator_status()
    RETURNS trigger AS
$$
DECLARE
    current_epoch    int8 := (SELECT mainnet_head_epoch());
    FAR_FUTURE_EPOCH int8 := 9223372036854775807;
BEGIN
    IF NEW.activation_eligibility_epoch > current_epoch THEN
        NEW.status := 'pending';
        CASE
            WHEN (NEW.activation_eligibility_epoch >= FAR_FUTURE_EPOCH) THEN NEW.substatus := 'pending_initialized';
            WHEN (NEW.activation_eligibility_epoch < FAR_FUTURE_EPOCH) AND (NEW.activation_epoch > current_epoch)
                THEN NEW.substatus := 'pending_queued';
            ELSE NEW.substatus := 'unknown';
            END CASE;
    ELSIF NEW.activation_eligibility_epoch <= current_epoch AND current_epoch < NEW.exit_epoch THEN
        NEW.status := 'active';
        CASE
            WHEN (NEW.activation_epoch <= current_epoch) AND (NEW.exit_epoch >= FAR_FUTURE_EPOCH)
                THEN NEW.substatus := 'active_ongoing';
            WHEN (NEW.activation_epoch <= current_epoch) AND
                 (current_epoch < NEW.exit_epoch AND NEW.exit_epoch < FAR_FUTURE_EPOCH) AND (NOT NEW.slashed)
                THEN NEW.substatus := 'active_exiting';
            WHEN (NEW.activation_epoch <= current_epoch) AND
                 (current_epoch < NEW.exit_epoch AND NEW.exit_epoch < FAR_FUTURE_EPOCH) AND NEW.slashed
                THEN NEW.substatus := 'active_slashed';
            ELSE NEW.substatus := 'unknown';
            END CASE;
    ELSIF NEW.exit_epoch <= current_epoch AND current_epoch < NEW.withdrawable_epoch THEN
        NEW.status := 'exited';
        CASE
            WHEN (NEW.exit_epoch <= current_epoch AND current_epoch < NEW.withdrawable_epoch) AND (NOT NEW.slashed)
                THEN NEW.substatus := 'exited_unslashed';
            WHEN (NEW.exit_epoch <= current_epoch AND current_epoch < NEW.withdrawable_epoch) AND NEW.slashed
                THEN NEW.substatus := 'exited_slashed';
            ELSE NEW.substatus := 'unknown';
            END CASE;
    ELSIF NEW.withdrawable_epoch <= current_epoch THEN
        NEW.status := 'withdrawal';
        CASE
            WHEN (NEW.withdrawable_epoch <= current_epoch) AND (NEW.balance != 0)
                THEN NEW.substatus := 'withdrawal_possible';
            WHEN (NEW.withdrawable_epoch <= current_epoch) AND (NEW.balance <= 0)
                THEN NEW.substatus := 'withdrawal_done';
            ELSE NEW.substatus := 'unknown';
            END CASE;
    ELSE
        NEW.status := 'unknown';
    END IF;

    IF (TG_OP = 'INSERT') AND (NEW.activation_epoch IS NOT NULL) THEN
        CASE
            WHEN NEW.activation_epoch < FAR_FUTURE_EPOCH
                THEN INSERT INTO validator_balances_at_epoch (epoch, validator_index, total_balance_gwei, current_epoch_yield_gwei)
                     VALUES (NEW.activation_epoch, NEW.index, 32000000000, 0);
            ELSE
            END CASE;
    END IF;

    IF (TG_OP = 'UPDATE') THEN
        CASE
            WHEN (NEW.activation_epoch < FAR_FUTURE_EPOCH) AND (OLD.activation_epoch != NEW.activation_epoch)
                THEN INSERT INTO validator_balances_at_epoch (epoch, validator_index, total_balance_gwei, current_epoch_yield_gwei)
                     VALUES (NEW.activation_epoch, NEW.index, 32000000000, 0);
            ELSE
            END CASE;
    END IF;

    UPDATE validators
    SET status    = NEW.status,
        substatus = NEW.substatus
    WHERE index = NEW.index;

    RETURN NEW;
END;
$$
    LANGUAGE 'plpgsql';

CREATE TRIGGER trigger_update_validator_status
    AFTER INSERT OR UPDATE OF balance, effective_balance, activation_eligibility_epoch, activation_epoch, exit_epoch, withdrawable_epoch, slashed
    ON validators
    FOR EACH ROW
EXECUTE PROCEDURE update_or_insert_validator_status();

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
    RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
    BEFORE UPDATE
    ON validators
    FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();