## 型アサーション

```
// interfaceと型T
i.(T)
```

- 1値返しはpanic
- 型アサーションはif文で1つずつ

## 型switch

```
switch v := i.(type) {
case T1:
}
```

- インターフェース型、ポインタ型をチェックできる