-- 1. Procedure get user balance
DELIMITER //
CREATE PROCEDURE GetUserBalance(
    IN p_userid VARCHAR(20)
)
BEGIN
    SELECT 
        b.balance as balance,
        b.updatedat,
        COALESCE(b.last_transaction_id, 0) as last_transaction_id
    FROM dbuserbalance b
    WHERE b.userid = p_userid;
END //
DELIMITER ;

-- 2. Procedure add balance
DELIMITER //
CREATE PROCEDURE AddUserBalance(
    IN p_userid VARCHAR(20),
    IN p_amount DECIMAL(15,2),
    IN p_source VARCHAR(50)
)
BEGIN
    DECLARE current_balance DECIMAL(15,2) DEFAULT 0.00;
    DECLARE new_balance DECIMAL(15,2);
    DECLARE last_insert_id BIGINT;
    
    START TRANSACTION;
    
    -- Validate amount
    IF p_amount <= 0 THEN
        SIGNAL SQLSTATE '45000' 
        SET MESSAGE_TEXT = 'Amount must be positive';
    END IF;
    
    -- Check if user exists in dbuserbalance, create if not
    IF NOT EXISTS (SELECT 1 FROM dbuserbalance WHERE userid = p_userid) THEN
        INSERT INTO dbuserbalance (userid, balance) VALUES (p_userid, 0.00);
    END IF;
    
    -- Get current balance with FOR UPDATE lock
    SELECT balance INTO current_balance
    FROM dbuserbalance WHERE userid = p_userid
    FOR UPDATE;
    
    -- Calculate new balance
    SET new_balance = current_balance + p_amount;
    
    -- Insert transaction record
    INSERT INTO balancehistory (
        userid, 
        changeamount, 
        balance_after, 
        type, 
        source
    ) VALUES (
        p_userid,
        p_amount,
        new_balance,
        'credit',
        p_source
    );
    
    -- Get the inserted transaction ID
    SET last_insert_id = LAST_INSERT_ID();
    
    -- Update user balance
    UPDATE dbuserbalance 
    SET balance = new_balance,
        updatedat = NOW(),
        last_transaction_id = last_insert_id
    WHERE userid = p_userid;
    
    -- Return the new balance
    SELECT new_balance as balance, last_insert_id as transaction_id;
    
    COMMIT;
END //
DELIMITER ;


-- 2. Procedure withdraw balance
DELIMITER //

CREATE PROCEDURE WithdrawUserBalance(
    IN p_userid VARCHAR(20),
    IN p_amount DECIMAL(15,2),
    IN p_source VARCHAR(50)
)
BEGIN
    DECLARE current_balance DECIMAL(15,2) DEFAULT 0.00;
    DECLARE new_balance DECIMAL(15,2);
    DECLARE last_insert_id BIGINT;
    DECLARE err_msg VARCHAR(255);

    -- Rollback on any error
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;

    START TRANSACTION;

    -- Validate amount
    IF p_amount <= 0 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Amount must be positive';
    END IF;

    -- Lock user row
    SELECT balance
    INTO current_balance
    FROM dbuserbalance
    WHERE userid = p_userid
    FOR UPDATE;

    -- Check User
    IF current_balance IS NULL THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'User account does not exist';
    END IF;

    -- Check sufficient balance
    IF current_balance < p_amount THEN
        SET err_msg = CONCAT(
            'Insufficient balance. Available: ',
            current_balance
        );

        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = err_msg;
    END IF;

    -- Calculate new balance
    SET new_balance = current_balance - p_amount;

    -- Insert history record
    INSERT INTO balancehistory (
        userid,
        changeamount,
        balance_after,
        type,
        source
    ) VALUES (
        p_userid,
        -p_amount,
        new_balance,
        'debit',
        p_source
    );

    SET last_insert_id = LAST_INSERT_ID();

    -- Update balance snapshot
    UPDATE dbuserbalance
    SET balance = new_balance,
        updatedat = NOW(),
        last_transaction_id = last_insert_id
    WHERE userid = p_userid;

    COMMIT;

    -- Return result
    SELECT
        new_balance AS balance,
        last_insert_id AS transaction_id;

END //

DELIMITER ;