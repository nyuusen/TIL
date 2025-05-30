# Docker

## 概要

なんとなく触っているDockerを以下を題材にきちんと学ぶ。

[入門 Docker](https://y-ohgi.com/introduction-docker/)

## 基礎編

### Dockerとは

- Dockerは任意のタイミングの状態（ランタイム・ライブラリ・コードのバージョン）を1つのスナップショットとして保存している
- 従来のVM型の仮想化は、物理レイヤの仮想化から行なっている
  - つまりVMには、
- 一方DockerはLinuxカーネルの機能を用いて1プロセスとして隔離された環境を実現している
- なのでDockerの方が軽量でオーバーヘッドが少ない環境を実現できる
- DockerはImmutable InfrastructureをDockerfileとイメージという機能で実現
  - 任意の時点で確実に動作するスナップショット（何か変更する場合は新しく構築する）
  - かつてはサーバーへ変更を加えて実現するMutableなものだった
  - また、イメージにはバージョン情報がつくので、それを指定することでロールバックが可能になり、可用性が向上する
- コンテナ自体は古くからある概念で、Dockerが流行ったのは「配布の容易性」がある（スナップショットを取得し配布を標準の機能として提供している）

### VMとDocker

- どちらも隔離された環境を提供する技術
- **VMはコンピュータ自体の抽象化（仮想化）するのに対し、Dockerはプロセス自体の抽象化（仮想化）を行う**
  - この2つは競合するのではなく、むしろ共存する

#### VM

- ハードウェアから上、ハードウェア・カーネル（OS）、ユーザースペースの低レイヤから仮想化を行う
- 例えばmacOS上でWindowsを動かすことが可能だったりして、非常に自由度が高い
- 仮想化の対象となる領域が広いためオーバーヘッドが大きくなる

#### Docker

- Linuxカーネルの機能を用いた技術で、cgroup・namespace・Capabilityのような機能を組み合わせて実現している
- VMとは異なり、ホストOSとリソースを共有し、効率的にホストOSのリソースを使うことができるので、高速化つ軽量な仮想化を実現している

- Linuxカーネルとは？
  - OSの中核となる部分で、ハードウェアとアプリケーションの間を取り持つ橋渡し役
  - 例えばファイルの読み書き・ネットワーク通信・メモリ管理・CPU割り当てなどの低レイヤ処理を全てカーネルが担っている
  - UbuntuやCentOSなどのディストリビューションは全てLinuxカーネルをベースにしている
  - カーネルが主に行なっているのは以下の4つ
    - プロセス管理：プロセスに対してCPU割り当て（今どのプロセスにCPU時間を割り当てるかをスケジュール・切り替え＝コンテキストスイッチ）
    - メモリ管理：アプリケーションへのメモリ割り当て（他のプロセスのメモリ領域を勝手にアクセスできないようにしたり、メモリ不足時はディスクにスワッピングしたり）
    - ファイルシステム管理：ファイル操作を仲介（実際にディスクを触るのはカーネルが行なっている）
    - ネットワーク管理：カーネルがパケット処理を行なっている

- でもDockerfileでFROM句にubuntu:latestみたいな記述をしているよね...？
  - ホスト側のLinuxカーネルを使うのに、コンテナ側でOS指定が必要な理由は何？という話
  - 結論、これはユーザースペースのOS環境を指定しているだけ
  - ユーザースペースというのは、コマンドやライブラリ（bashとか）
    - Distrolessとか軽量なOSイメージを指定すると、デフォルトのままではcurl等のコマンドがインストールされておらず使用できないのは、コンテナ側のユーザースペースを使っているため

### Docker Image

- イメージは、任意のタイミングのスナップショットとしての役割を持つ
- ファイルシステムのスナップショットである
  - もっとシンプルに表現すると、OSの中身のフォルダ構造一式をZIP圧縮したようなもの
  - 例えばubuntuイメージの中身は、ユーザースペース（`ls /`した時に見るような bin/,etc/,lib/等）のファイル・フォルダ一式が入っている（でもカーネル部分は含まない）
- てっきりイメージ＝プロセスのスナップショットと理解していたけどそれは間違っていて、**イメージ＝実行前のファイル群**が正しい表現
  - イメージを使ってプロセス実行したものが「コンテナ」かな？
- イメージを指定するときの命名は<イメージ名>:<タグ>であり、<タグ>を省略するとlatestタグが自動的に付与される

#### レイヤー構造

- Dockerイメージはレイヤーの積み重ね
- RUNやCOPY毎に新しいレイヤが作られる
- レイヤは読み取り専用で、キャッシュとして使い回せる
- レイヤー構造である理由
  - キャッシュ：同じベースイメージなら、そのレイヤはpull済み
  - ストレージ節約：変更があった部分だけ新しいレイヤに保存
    - もし仮にレイヤ構造ではない場合はフルフルのものを保存しなければいけない
  - 高速化：レイヤー単位でダウンロード・展開
- docker pushすると以下のようなハッシュ値が表示されるが、これがまさにレイヤ構造（レイヤ単位で処理している）
  ```
    The push refers to repository [123456789012.dkr.ecr.us-east-1.amazonaws.com/myapp]
    d4f4f6a4b8c2: Pushed 
    7a0437f04f83: Pushed 
    8c662931926f: Pushed 
    ...
  ```
  - 各レイヤのコンテンツハッシュ(SHA256)らしい
- 整理すると、Dockerイメージはファイルシステムのスナップショットであり、それはレイヤー構造になっている
  - 以下のDockerfileを例に考えると、
    ```
      FROM ubuntu:latest
      RUN apt-get update
      RUN apt-get install -y curl
    ```
    - 各行がレイアとなっていて、そのレイヤで変更があったファイル群を持っているイメージ
    - この辺りは全て読み取り専用なので、変更がなければ
  - この上に、Dockerはコンテナ起動した時に、書き込みできるレイヤを作っている


## Dockerfile

- COPY
  - 2つの引数を設定する
    - 1つ目はホスト側のディレクトリ、2つ目はDocker側のディレクトリ
  - ホスト側のディレクトリは `docker build .` で指定したディレクトリ
    - この場合 `.` を指定しており、カレントディレクトリが参照される
  - Docker側はデフォルトのパス、もしくは `WORKDIR` で定義されたディレクトリを参照する
- EXPOSE
  - このポートを使いますよというドキュメント的な宣言
- CMD
  - Dockerはここで設定したコマンドがフォアグラウンドで実行されている間が生存期間となる

#### RUN vs CMD

- RUN: 
  - イメージビルド時に実行される
  - レイヤとしてファイルシステムに反映されキャッシュが効く
- CMD:
  - コンテナ起動時に実行される
  - イメージには反映せず、ただの実行時オプション

#### CMD vs ENTRYPOINT

- CMDはデフォルト設定で、オーバーライド可能
- 一方、ENTRYPOINTは起動時に必ず実行される
- テクニックとしてENTRYPOINTでコマンドを指定し、CMDで引数を指定するというのがある（引数だけ利用者側で指定可能となる）

```
基本的に CMD を使うのが良いでしょう。
ENTRYPOINT はDocker起動時のコマンドを強制します。
コマンドのラップをするDocker Image の場合は ENTRYPOINT のほうが好ましいですが、一般的なWebアプリケーションの場合は CMD を使用する方がユーザーにとって使いやすいDocker Image になります。
```

#### COPYは最後に実行するとキャッシュが効きやすい

- Dockerfileの前段で、`COPY . .`を実行してしまうと、ローカル側のソースコードが1文字でも変わっていると、そのレイヤのキャッシュが効かなくなる
- レイヤーのキャッシュは、親が変わると問答無用でキャッシュ無効化されるので、`COPY . .`以降は全てキャッシュが利用されなくなる（レイヤは差分を持っているようなものだから、親が変更があったら当然子にも影響あるよねということだと理解した）
- なので、変更が頻繁にあるような`COPY . .`の処理は、Dockerfileの中で後ろの方に持っていくと良い
- Nodeならpackage系だけを最初にコピーする→依存関係をインストール→ソースコードをコピーしてくるみたいな形で工夫できる

## Container

- イメージがスナップショットだとすると、そのスナップショットから起動したプロセスがコンテナ
- コンテナは「1つのコマンド（プロセス）をフォアグラウンドで動かす」ように設計されている
  - コンテナは1つのコマンドを隔離された環境で実行し、そのコマンドの実行がフォアグラウンドで終了するまで生存する
  - ライフサイクル
    - Image -- (docker run <$IMAGE>) --> RUNNING -- STOPPED -- DELETED
    - 正常終了 or 異常終了 or docker killするとSTOPPED
    - docker rmするとDELETED
    - pauseすると停止状態を表すPAUSEDにもなる
- プロセスの隔離
  - コンテナ内のプロセスはホストマシンや他のコンテナと隔離されて実行される
  - CMDもしくはENTRYPOINTで定義されたプロセスはPID 1となる

## Network

- Dockerではネットワークの扱いが重要となる
- 1コンテナでは1プロセスを動かす設計となっている
- nginxとphp-fpmのように複数プロセスを協調して動かす必要があるときはソケットではなく、ネットワークで通信を行うことが推奨されている
- Dockerでのネットワークは特にKubernetes・ECS・docker-composeのような各種オーケストレーションツールを使用する際に意識する必要がある

### Driverの種類

- ネットワークドライバーはネットワークの振る舞いの定義で、デフォルトでは2種類ある
- 複数のコンテナ（プロセス）はネットワークを介して通信を行う

#### bridge

- 基本的にはこれ
- コンテナごとに仮想IPが割り振られ、同じネットワークに属するコンテナ間で通信が可能
- こうすることでコンテナ同士はコンテナ名で互いに名前解決して通信できる
- 少し深掘りすると..
  - 何も指定せずにコンテナ起動すると、docker0という名前のbridgeネットワークに所属する
    - docker0というのははホストOS上に仮想ブリッジという仮想スイッチ
    - hostネットワーク -- docker0（仮想ブリッジ） -- コンテナA, B...という構成
    - この構成では、各コンテナに独立した仮想ネットワーク内のIPが与えられ、同じネットワークにいるコンテナ間はIPやホスト名で通信できる
    - コンテナに対し、外部からアクセスしたい場合はNAT変換によるポートフォワードが必要(ex: `docker run -p 8080:80 nginx`)
  - Linuxカーネルのbridgeネットワークを使用する

#### host

- コンテナがホストのIPアドレスとポート空間（ネットワーク名前空間）をそのまま使う
  - 仮想NICやブリッジを介さずに、直接ホストのIP・ポート空間にアクセスできる
  - 例：コンテナが80番でListenすると、それはホスト側の80番ポートを使うことになる
- オーバーヘッドが少ない分通信が速いが、ポートの競合に注意する必要がある
- ホスト側のlocalhost:80等でそのままコンテナにアクセスできるような仕組み

#### none

- コンテナにネットワークを割り当てない
- セキュリティ上、外部から完全に切り離したい時に使う

#### Docker Composeにおけるネットワークの考え方

- Docker Composeはマルチコンテナを簡単に定義・管理できるツール
- 内部ネットワークは自動で作成される
- 各サービスは、自動で同じカスタムネットワークに所属するので、互いにサービス名で名前解決が可能となる（内部DNSによって解決されている）

## Volume

- データを永続化するための機能
- Dockerコンテナは基本的にはエフェメラル（短命）なもので、ライフサイクルの終了とともにコンテナ上で作成されたファイルは消失する
- ボリュームタイプには以下の2種類がある

### Data Volume

- コンテナのライフサイクルの外で管理されるファイル/ディレクトリの設定
- `-v <CONTAINER PATH>` or `-v <HOST PATH>:<CONTAINER PATH>`
- コンテナの外側＝ホスト側にファイルが保管される

### Data Volume Container

- 他のコンテナで指定されているボリュームを参照するための機能（コンテナ間でボリュームを共有する）
- `--volumes-from`でコンテナ名を指定することで、別のコンテナのボリュームを参照できる
- 今の時代では、Named Volume(`-v mydata:data`)や上記のData VolumeのBind Mount(`-v <HOST PATH>:<CONTAINER PATH>`)を使用するケースが多そう
  - Named Volumeの補足として、Linuxの場合は大抵`/var/lib/docker/volumes/mydata/_data/`に保存される

## プロダクションでの活用Tips

### セキュリティ

- rootユーザーを使わない

## 1コンテナ1プロセスの原則

なぜ1コンテナ1プロセスが良いか？

- 可観測性
  - プロセス単位でメトリクスやログ収集がしやすい
- 責務の明確化
  - まぁ責務は明確にして分離すべきという話
- リソース制御
  - CPUやメモリの利用量をプロセス単位で制御できる
- 障害隔離性
  - 複数のプロセスが同一コンテナにあると、一方のクラッシュがもう一方に影響しやすい