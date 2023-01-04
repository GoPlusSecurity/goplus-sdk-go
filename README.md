# GoPlus API SDK for Go

### Installation
> go get github.com/GoPlusSecurity/goplus-sdk-go

### Get AccessToken
```go

key := ""
secret := ""
app := auth.NewApp(key, secret, nil)
data, err := app.GetAccessToken()

if err != nil {
    panic(err.Error())
}

if data.Code != 1 {
    panic(data.Message)
}

accessToken := data.Result.AccessToken
```

### Get Supported Blockchains

```go
accessToken := ""
supportedChains := NewSupportedChains(accessToken, nil)
data, err := supportedChains.Run("")

if err != nil {
    panic(err)
}
if data.Code != 1 {
    panic(data.Message)
}

```

### Token Security

```go

accessToken := ""
tokenSecurity := token.NewTokenSecurity(accessToken, nil)
chainId := "1"
contractAddresses := []string{"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"}
data, err := tokenSecurity.Run(chainId, contractAddresses)
if err != nil {
    panic(err)
}
if data.Code != 1 {
    panic(data.Message)
}
value, ok := data.Result["0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"]

if ok {
    fmt.Printf("%v", value["token_symbol"])
}

```

### Address Security

```go

accessToken := ""
addressSecurity := address.NewAddressSecurity(accessToken, nil)
data, err := addressSecurity.Run("", "0xc8b759860149542a98a3eb57c14aadf59d6d89b9")
if err != nil {
    panic(err)
}
if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result["data_source"])
```


### Approval Security API v1
```go

accessToken := ""
approvalSecurityV1 := approval.NewApprovalSecurity(accessToken, nil)
data, err := approvalSecurityV1.Run("1", "0x4639cd8cd52ec1cf2e496a606ce28d8afb1c792f")

if err != nil {
    panic(err)
}

if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result["contract_name"])

```


### Approval Security API v2

#### Token Approval Security

```go

accessToken := ""
approvalSecurityV2 := approval.NewApprovalSecurityV2(accessToken, nil)
data, err := approvalSecurityV2.Token("56", "0xd018e2b543a2669410537f96293590138cacedf3")
if err != nil {
    panic(err)
}
if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result[0]["token_address"])

```

#### ERC721 NFT Approval Security

```go
accessToken := ""
approvalSecurityV2 := approval.NewApprovalSecurityV2(accessToken, nil)
data, err := approvalSecurityV2.ERC721NFT("1", "0xd95dbdab08a9fed2d71ac9c3028aac40905d8cf3")
if err != nil {
    panic(err)
}
if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result[0]["nft_address"])
```

#### ERC1155 NFT Approval Security

```go

accessToken := ""
approvalSecurityV2 := approval.NewApprovalSecurityV2(accessToken, nil)
data, err := approvalSecurityV2.ERC1155NFT("56", "0xb0dccbb9c4a65a94a41a0165aaea79c8b2fc54ce")
if err != nil {
    panic(err)
}
if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result[0]["nft_address"])
```

### Signature Data Decode
```go
accessToken := ""
signatureDataDecode := decode.NewSignatureDataDecode(accessToken, nil)

chainId := "1"
contractAddress := "0x4cc8aa0c6ffbe18534584da9b592aa438733ee66"
inputData := "0xa0712d680000000000000000000000000000000000000000000000000000000062fee481"
data, err := signatureDataDecode.Run(chainId, contractAddress, inputData)
if err != nil {
    panic(err)
}
if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result[""])
```
### NFT Security
```go

accessToken := ""
nftSecurity := nft.NewNFTSecurity(accessToken, nil)

chainId := "1"
contractAddress := "0x82f5ef9ddc3d231962ba57a9c2ebb307dc8d26c2"
data, err := nftSecurity.Run(chainId, contractAddress)

if err != nil {
    panic(err)
}

if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result["nft_address"])
```

### dApp Security Info API
```go
accessToken := ""
dAppSecurity := dapp.NewDAppSecurity(accessToken, nil)

url := "https://for.tube"
data, err := dAppSecurity.Run(url)

if err != nil {
    panic(err)
}

if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result["project_name"])
```

### Phishing Site Detection API
```go
accessToken := ""
phishingSiteSecurity := phishing_site.NewPhishingSiteDetection(accessToken, nil)

url := "https://xn--cm-68s.cc/"
data, err := phishingSiteSecurity.Run(url)

if err != nil {
    panic(err)
}

if data.Code != 1 {
    panic(data.Message)
}

fmt.Printf("%v", data.Result["phishing_site"])
	
```
