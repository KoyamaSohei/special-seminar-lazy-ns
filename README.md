# special seminar lazy ns

全てのDNSクエリ(`Aレコード`)で`TARGET_IP`を返す
`Aレコード`以外は`DNS_SERVER`にクエリを流す

## Environment Variables

- TARGET_IP(required)
- DNS_SERVER(required)
