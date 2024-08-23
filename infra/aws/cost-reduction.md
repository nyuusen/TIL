# AWSにおけるコスト削減対応

## はじめに

業務や参加した勉強会でコスト削減についての学びがあったのでメモ

勉強会URL: https://www.youtube.com/live/WOsD-Rk3SVg?si=El_AK4863iGt8L42

## 大前提としての心構え

- コスト削減対応で浮いた分は単純な利益増につながる
- つまり後回しにせず、早くやろうぜという気持ちで

## やり方

- モブでやってあーだこーだ言ってみる
    - 職能を絞らないでやるのも視野が広がって良いアイディアが生まれたりする

## 具体的なコストカットポイント

- NATゲートウェイ→VPC Endpoint
    - やはりNATゲートウェイは高い！
- ECRからのイメージプルをECR pull through cacheでキャッシュ化
- コンテナイメージのサイズダウン
- 各インスタンスのサイジングやスペック見直し
