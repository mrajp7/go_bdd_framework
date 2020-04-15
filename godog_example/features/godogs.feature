# file: $GOPATH/godogs/features/godogs.feature
Feature: eat godogs
    In order to be happy
    As a hungry gopher
    I need to be able to eat godogs

    Scenario: Eat 5 out of 12
        Given there are 12 godogs
        When I eat 5
        Then there should be 7 remaining

    Scenario: Eat 2 out of 12
        Given there are 12 godogs
        When I eat 2
        Then there should be 7 remaining

    Scenario: When nothing is available to Eat
        Given there are 0 godogs
        When I eat 1
        Then there should be an error indicating "No godogs available to Eat!"