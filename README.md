# awsconfctl

`awsconfctl` is a command line tool to work with a couple of AWS configuration services (Systems Manager Parameters or Secrets Manager). 

## Installation

### DIY build

For go 1.12 and higher:

```
git clone https://github.com/sarmad-abualkaz/awsconfctl.git

go build
```

## Usage

Note: this tool expects to find AWS credentials (Access Key and Secret Key) in 

As it stands this tool can perform the following sub-commands:

### List SSM (System Manager) Parameters

Lists a group of [AWS SSM parameters](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) containing a specific string pattern. 

exmaple:
`awsconfctl ssm-list-params --contains <string>`

Optional flags:

- `--region <aws-region>` to specify AWS Region (can also be passed via environment variable or through `AWS_REGION=<aws-region>`) (defaults to `us-east-1`). 

- `--profile <aws-profile>` to specify AWS Profile (defaults to `dev`).

- `--maxRes <integer>` to specifiy the maximum result if required to limit output.

### Get SSM Parameter Value

Retreives the value of an [AWS SSM parameters](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) by specifying the name of parameter.

`awsconfctl ssm-get-value --name <parameter-name>`

Optional flags:

- `--region <aws-region>` to specify AWS Region (can also be passed via environment variable or through `AWS_REGION=<aws-region>`) (defaults to `us-east-1`). 

- `--profile <aws-profile>` to specify AWS Profile (defaults to `dev`).

### List AWS Secrets (Secrets Manager)

Lists a group of [AWS Secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) containing a specific string pattern. 


`awsconfctl sm-list-secrets --contains <name>`

Optional flags:

- `--region <aws-region>` to specify AWS Region (can also be passed via environment variable or through `AWS_REGION=<aws-region>`) (defaults to `us-east-1`). 

- `--profile <aws-profile>` to specify AWS Profile (defaults to `dev`).

- `--maxRes <integer>` to specifiy the maximum result if required to limit output.

### Get AWS Secret Value

Retreives the value of an [AWS Secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) by specifying the name of a secret.

`awsconfctl sm-get-secret-value --name <secret-name/alias>`

Optional flags:

- `--region <aws-region>` to specify AWS Region (can also be passed via environment variable or through `AWS_REGION=<aws-region>`) (defaults to `us-east-1`). 

- `--profile <aws-profile>` to specify AWS Profile (defaults to `dev`).

### Apply (for updates or creations of configs)

This deals with both [AWS SSM parameters](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) or [AWS Secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) and attempts to either update or create new configurations (parameters or secrets).

#### How to use

`awsconfctl apply -f <file-path>`

Optional flags:

- `--region <aws-region>` to specify AWS Region (can also be passed via environment variable or through `AWS_REGION=<aws-region>`) (defaults to `us-east-1`). 

- `--profile <aws-profile>` to specify AWS Profile (defaults to `dev`).

#### Expected YAML schema

For dealing with [AWS SSM parameters](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) the example below provides a good YAML sample:

```
configSetup:
  configType: systemManager
  params:
  - key: foo-param
    value: foo-vale
    type: string
  - key: bar-param
    value: bar-value
    type: secureString (optional)
    KMSKey: <kms-key-id> (required if secureString is the type)
  - ...
```

For dealing with [AWS Secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) the example below provides a good YAML sample:

```
configSetup:
  configType: secretsManager
  params:
  - key: foo-secret
    value: foo-vale
    KMSKey: <kms-key-id>
  - key: bar-secret
    value: bar-value
    KMSKey: <kms-key-id>
  - ...
```

### Delete (for removal of configs)

This deals with both [AWS SSM parameters](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) or [AWS Secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) and remove configurations as a result (parameters or secrets).

#### How to use

`awsconfctl delete -f <file-path>`

Optional flags:

- `--region <aws-region>` to specify AWS Region (can also be passed via environment variable or through `AWS_REGION=<aws-region>`) (defaults to `us-east-1`). 

- `--profile <aws-profile>` to specify AWS Profile (defaults to `dev`).

- `--deleteForGood` or `-d` (boolean) special flag for dealing with [AWS Secrets]((https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) ) only to remove a secret without a recovery window (defaults to `false`).

- `--recWindow` (int) special flag for dealing with [AWS Secrets]((https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) ) only to setup a recovery window (Note this defaults to `7`and will need to be set to `0` when passing `-d`/ `--deleteForGood` to `true`

#### Expected YAML schema

For dealing with [AWS SSM parameters](https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html) the example below provides a good YAML sample:

```
configSetup:
  configType: systemManager
  params:
  - key: foo-param
    value: foo-vale
    type: string
  - key: bar-param
    value: bar-value
    type: secureString (optional)
    KMSKey: <kms-key-id> (required if secureString is the type)
  - ...
```

For dealing with [AWS Secrets](https://docs.aws.amazon.com/secretsmanager/latest/userguide/intro.html) the example below provides a good YAML sample:

```
configSetup:
  configType: secretsManager
  params:
  - key: foo-secret
    value: foo-vale
    KMSKey: <kms-key-id>
  - key: bar-secret
    value: bar-value
    KMSKey: <kms-key-id>
  - ...
```
