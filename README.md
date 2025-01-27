# Kamoney API2 SDK


```bash
go get github.com/kamoney/sdk_golang@latest
```

How to import ?

```text
import(
    private "github.com/kamoney/sdk_golang/kamoney_sdk/private_kamoney"
    public "github.com/kamoney/sdk_golang/kamoney_sdk/public_kamoney"
)



func main() {
 public_ep := public.NewPublicRequests("public")
 response, err := public_ep.ServicesOrder()
 if err != nil {
  panic(err)
 }
 fmt.Println(response)

 private_ep := private.NewPrivateRequests("email", "pass", "public", "secret")
 response2, err2 := private_ep.GetAccountInfo(kamoney_sdk_dtos.ChangeAccountInfoRequestParams{})
 if err2 != nil {
  panic(err2)
 }
 fmt.Println(response2)
}

```
