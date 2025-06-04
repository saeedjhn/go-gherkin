Feature: User Registration and Welcome Email

  Scenario: Successful user registration
    Given a new user with valid details
    When the user registers
    Then the user account should be created
    And a welcome email should be sent

  @wip
  Scenario: User tries to register with an invalid email
    Given a new user with invalid email "invalid-email"
    When the user tries to register
    Then the registration should fail
    And an error message "invalid email format" should be shown

  Scenario: User tries to register with an existing email
    Given an existing user with email "existing@example.com"
    When a new user tries to register with email "existing@example.com"
    Then the registration should fail
    And an error message "email already exists" should be shown

#  Scenario: Successful user registration and sending welcome email
#    Given a new user with valid details (name: "John Doe", email: "john@example.com", password: "Password123")
#    When the user registers
#    Then the user account should be created
#    And a welcome email should be sent to "john@example.com"

#  Scenario: User tries to register without providing an email
#    Given a new user with no email
#    When the user tries to register
#    Then the registration should fail
#    And an error message "Email is required" should be shown

#  Scenario: User tries to register with a short password
#    Given a new user with email "user@example.com" and password "short"
#    When the user tries to register
#    Then the registration should fail
#    And an error message "Password must be at least 8 characters long" should be shown
