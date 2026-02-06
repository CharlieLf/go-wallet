INSERT INTO dbuserbalance (userid, balance ) VALUES (1, 0)
INSERT INTO dbuserbalance (userid, balance ) VALUES (2, 0)

CALL AddUserBalance('1', 350000.00, 'topup');
CALL AddUserBalance('2', 500000.00, 'topup');