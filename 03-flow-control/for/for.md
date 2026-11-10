## forループ

```
for 初期化; 条件; 後処理 {
    // loop
}
```

```
for i := 0; i < 10; i++ {
    sum += i
}
```

- 読みやすさ優先で省略記法を扱っていく

## while風

- Goにはwhileがない
- forを使っていく

```
sum := 1
for sum < 1000; {
    sum += sum
}
```