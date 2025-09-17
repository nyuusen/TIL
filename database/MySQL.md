# MySQL

## コマンド

### 確認・調査系

- 該当テーブルのカラムの詳細情報

```
show columns from [table_name];
```

- 該当テーブルのカラムのより詳細な情報(コメントとか\
  \Gをつけることで行表示になり見やすくなる

```
select * from INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME='[table_name]'\G
```

- テーブルの情報からCREATE文を生成する

```
show create table [table_name];
```

### DDL
