# 🧪 BDD with Gherkin & Godog in Go

This project demonstrates how to use **Behavior-Driven Development (BDD)** in Go using the **Gherkin language** and
the [cucumber/godog](https://github.com/cucumber/godog) library.

BDD is an effective approach for testing business logic in a way that's **clear and understandable for stakeholders**.

---

## 🛠️ Steps to Run

### 1. Install `cucumber/godog`

Make sure Go is installed. Then run:

```bash
go install github.com/cucumber/godog/cmd/godog@latest
```

Ensure `~/go/bin` is in your `PATH`.

### 2. Create a Gherkin Feature

Define your use cases in `.feature` files using the Gherkin syntax:

```gherkin
Feature: User Registration and Welcome Email

  Scenario: Successful user registration
    Given a new user with valid details
    When the user registers
    Then the user account should be created
    And a welcome email should be sent
```

Save the file as, for example: `features/register.feature`.

### 3. Create Step Definitions

Generate step stubs from your feature file:

To run all features:

```bash
godog .
```

To run a specific feature file:

```bash
godog tests/features/auth/register.feature
```

To run godog with the `-t` flag to match the tag:

```bash
godog . -t=@wip
godog tests/features/auth/register.feature -t=@wip
```

This outputs step definition templates in the terminal.

Copy and implement the steps inside a Go test file like `register_steps_test.go`:

```go
func iHaveAValidUser() error {
// implementation here
return nil
}
```

You can organize your step definitions inside the `/tests` folder.

### 4. Run Tests

Use the standard Go test command to run the BDD tests:

```bash
go test -v ./tests/...
```

---

## ✅ Benefits of BDD

- Aligns developers and stakeholders with shared language (Gherkin)
- Clear, testable business logic
- Seamless test automation using Go