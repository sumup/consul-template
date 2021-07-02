# Consul Template

**This is a fork modified to support AWS SSM Parameter & AWS Secrets Manager seemlessly**.

For the official README of hashicorp/consul-template check https://github.com/hashicorp/consul-template.

## Modifications

The following env variables affect functionality:
* `CONSUL_TEMPLATE_KEY_STRATEGY` - changes which backend service consul-template uses to fetch keys from
    * If set to `aws_ssm_parameter_store`, will get keys from AWS SSM Parameter store. 
      AWS settings are loaded from `AWS_REGION`, `~/.aws/config`, `~/.aws/credentials`, etc.
    * If set to anything else or blank, will get keys from HashiCorp Consul, just as before.
* `CONSUL_TEMPLATE_SECRET_STRATEGY` - changes which backend service consul-template uses to fetch secrets from
    * If set to `aws_secrets_manager`, will get keys from AWS Secrets Manager.
      AWS settings are loaded from `AWS_REGION`, `~/.aws/config`, `~/.aws/credentials`, etc.
    * If set to anything else or blank, will get keys from HashiCorp Vault, just as before.
* `CONSUL_TEMPLATE_KEY_PREFIX` - changes every `key '<path>'` path found in `*.ctmpl` dependencies to retrieve from `<prefix><path>`. **Only works with AWS SSM Parameter Store**.
* `CONSUL_TEMPLATE_SECRET_PREFIX` - changes every `with secret '<path>'` path found in `*.ctmpl` dependencies to retrieve from `<prefix><path>`. **Only works with AWS Secrets Manager**.
