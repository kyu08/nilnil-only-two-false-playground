# nilnil-only-two-false-playground
[`nilnil@v1.1.0`](https://github.com/Antonboom/nilnil/releases/tag/v1.1.0)でリリースされた3つ以上の連続するnilを検出する`only-two`オプションのサンプルコードです。

サンプルコードには`return nil, nil, nil`のようなコードが含まれているため`only-two=false`を指定するとエラーになることが以下の手順で確認できます。

1. `make lint`
2. following error will be shown
    ```sh
    ❯ make lint
    GOTOOLCHAIN=go1.24.0 go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.0.2 run  ./... --config ./.golangci.yml
    main.go:38:3: return both a `nil` error and an invalid value: use a sentinel error instead (nilnil)
                    return nil, nil, nil
                    ^
    1 issues:
    * nilnil: 1
    exit status 1
    make: *** [lint] Error 1
    ```
