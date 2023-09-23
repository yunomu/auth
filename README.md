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
    WhiteList[("WhiteList\n(S3)")]
  end

  UserPool -- PreAuthentication --> PreAuth
  UserPool -- PreSignup --> PreSign
  PreAuth -- read only --> Product
  PreSign -- read only --> WhiteList

  subgraph otherproduct[Other product]
    Function[LambdaFunction]
  end

  PreAuth -. invoke if registered .-> Function
```

## Table: Product

|AttrName|Type|Schema|Description        |
|--------|----|------|-------------------|
|ClientId|S   |PK    |                   |
|AppCode |S   |      |                   |
|Created |N   |      |                   |
|FuncArn |S   |      |ClientTriggerLambda|
