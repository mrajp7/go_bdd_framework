#file: features/account.feature
Feature: bank account
    A user's bank account must be able to withdraw, deposit and transfer cash

    Background:
        Given a user bank account with 0$
        And a beneficiary bank account with 0$

    # Test ID:  001
    # Author :  mohanraj.palanisamy
    @Deposit
    Scenario Outline: Deposit
        Given a user bank account with <start>$
        When he deposits <deposit>$
        Then the user bank account should have a balance of <end>$

        Examples:
            | start | deposit | end |
            | 10    | 0       | 10  |
            | 10    | 10      | 20  |
            | 100   | 50      | 150 |

    # Test ID:  002
    # Author :  mohanraj.palanisamy
    @Deposit
    Scenario: Deposit a non positive number
        Given a user bank account with 10$
        When he deposits 0$
        Then the system should throw an error
            """
            Not a valid amount to deposit: 0$
            """

    # Test ID:  003
    # Author :  mohanraj.palanisamy
    @Withdrawal
    Scenario: Withdrawal when account has insufficient funds
        Given a user bank account with 10$
        When he withdraws 50$
        Then the system should throw an error
            """
            Insufficient Funds! Available Balance: 10$
            """
        And the user bank account should have a balance of 10$

    # Test ID:  004
    # Author :  mohanraj.palanisamy
    @Withdrawal
    Scenario Outline: Withdrawal
        Given a user bank account with <start>$
        When he withdraws <withdrawal>$
        Then the user bank account should have a balance of <end>$
        And no error should be found

        Examples:
            | start | withdrawal | end |
            | 10    | 0          | 10  |
            | 20    | 10         | 10  |
            | 100   | 50         | 50  |

    # Test ID:  005
    # Author :  mohanraj.palanisamy
    @Transfer
    Scenario Outline: Transfer Fund
        Given a user bank account with <start>$
        And a beneficiary bank account with <b_start>$
        When  he transfers <transfer>$
        Then  the user bank account should have a balance of <end>$
        And the beneficiary bank account should have a balance of <b_end>$

        Examples:
            | start | b_start | transfer | end | b_end |
            | 50    | 100     | 20       | 30  | 120   |
            | 20    | 0       | 10       | 10  | 10    |
            | 100   | 21      | 100      | 0   | 121   |

    # Test ID:  006
    # Author :  mohanraj.palanisamy
    @Transfer @Negative
    Scenario: Transfer Fund when Account has insufficient funds
        Given a user bank account with 10$
        And a beneficiary bank account with 0$
        When  he transfers 50$
        Then  the system should throw an error
            """
            Insufficient funds! Requested Amount: 50$, Available Balance: 10$
            """
        And the user bank account should have a balance of 10$
        And the beneficiary bank account should have a balance of 0$
