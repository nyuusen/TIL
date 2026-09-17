# AWS CLI　補完

[Configuring command completion in the AWS CLI](https://docs.aws.amazon.com/cli/v1/userguide/cli-configure-completion.html#cli-command-completion-linux)

```
❯ which aws_completer
/opt/homebrew/bin/aws_completer

❯ autoload bashcompinit && bashcompinit
❯ autoload -Uz compinit && compinit
❯ complete -C '/opt/homebrew/bin/aws_completer' aws

# 上記を~/.zshrcなどに入れればシェル起動のたびに有効化される
```