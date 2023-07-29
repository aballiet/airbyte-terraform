# Testing SpeakEasy for terraform provider generation

-> Generated a terraform provider for Airbyte Open Source with [SpeakEasy](https://speakeasyapi.dev/)

With this tool you can either directly generate a terraform provider using their Website or using their dedicated CLI.

Bot solutions require to have a SpeakEasy account

# Steps for CLI genration
- Download Airbyte OSS OpenAPI defintion file
- Install SpeakEasy CLI using `brew install speakeasy-api/homebrew-tap/speakeasy`
- Run command to generate Terraform provider :

```bash
speakeasy generate sdk -l terraform -o ./generated -s ./data/airbyte_open_api_config.yaml`
```
- Generated provider is located in the `generated` folder
- publish terraform provider on Terraform

