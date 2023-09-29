# auth

UserPool

# Development

## Structure

```mermaid
graph LR;
  subgraph auth
    UserPool[UserPool]
    PreSign["PreSignupFunction\n(Lambda)"]
    PreAuth["PreAuthenticationFunction\n(Lambda)"]
    Product[("ProductTable\n(DynamoDB)")]
    Restriction[("Restriction\n(DynamoDB)")]
  end

  UserPool -- PreAuthentication --> PreAuth
  UserPool -- PreSignup --> PreSign
  PreAuth -- read only --> Product
  PreSign -- read only --> Restriction
  PreAuth -- read only --> Restriction
  PreSign -- read only --> Product

  subgraph otherproduct[Other product]
    Function[LambdaFunction]
  end

  PreAuth -. invoke if registered .-> Function
```

## Table: Product

|AttrName |Type|Schema|Description        |
|---------|----|------|-------------------|
|ClientId |S   |PK    |                   |
|AppCode  |S   |      |                   |
|FuncArn  |S   |      |ClientTriggerLambda|
|Timestamp|N   |      |(Micro second)     |

## Table: Restriction

|AttrName |Type|Schema|Description        |
|---------|----|------|-------------------|
|Email    |S   |PK    |                   |
|AppCodes |SS  |      |                   |
|Timestamp|N   |      |(Micro second)     |
