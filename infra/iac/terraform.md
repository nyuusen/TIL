# Terraform

## はじめに

業務で雰囲気で触っているTerraformをきちんと学んでみる。

## コマンド群

- init
  - 実行するディレクトリにおける準備を行う
  - 以下が生成される
    - .terraform/ : providerの実体が置かれる
      - これはコマンドを実行するディレクトリごとに作成されるので、離任したプロジェクトのもの等を放置しているとそれなりに容量を食ってしまう点に注意
    - .terraform.lock.hcl : providerの同じバージョンを使うためのファイル
- plan
  - 構築するインフラの計画を行う
- apply
  - 実際にインフラの構築を行う
- destroy
  - インフラの削除を行う

## Tips

### 変数定義

- Terraformでは`variables.tf`を用いて変数を定義し、その変数を`var.`の形で参照することができる
  ```variables.tf
  variable "allow_ssh" {
    type = bool
    
    default = false
  }
  ```
- 値を何もセットしなかった場合はterraform planやapplyコマンド実行時にCLIで聞かれる(対話形式でコマンド実行が進むイメージ)
- 値をセットする方法としては、[Assigning Values to Root Module Variables](https://developer.hashicorp.com/terraform/language/values/variables#assigning-values-to-root-module-variables)に記載されている通りいくつかある
- (恐らく業務では扱う変数の数は一定あると思うので)`.tfvars`に変数の値を代入していく形が多くなりそう
  - 毎回、`terraform plan --var-file="ecs/dev/ecs.tfvars"`みたいに指定するのはとても面倒なので、各リソースのディレクトリ構成を統一させて、Makeコマンドなりで各リソース名や環境名を引数としてもらい、Makefileの中でterraformコマンドを組み立てるのが良さそう(実際に自分が今担当しているプロジェクトではこの作りになっている)

### 状態管理(tfstate)

- これはapplyした後の状態を管理するためのファイル
  - このファイルがあること、一度applyした後にもう一度applyを実行すると前回との差分だけが適用される
- 複数人で運用する場合は、共有ストレージ(S3等)に管理すべきである
  - そうでないと、人によってインフラの実態とtfstateの中身とで乖離ができてしまい、正しく運用できなくなってしまう
- このtfstateをどのように管理するかを指定することをbackendという
- 

### ディレクトリ構成

- 大きく分けて2つのディレクトリに分けるのが良い
  - 1.リソースの共通情報定義用(ex: `modules/`)
  - 2.環境毎の情報定義用(ex: `env/`)
    - 2の`main.tf`を以下のようにすることでmodules配下の共通情報を読み取ることができる
      ```main.tf
      module "ec2" {
          source = "../../modules/ec2"
          allow_ssh = false
      }
      ```

### .terraform/や.terraform.lock.hclの管理方法

### 複数人で運用する際に気をつけること

