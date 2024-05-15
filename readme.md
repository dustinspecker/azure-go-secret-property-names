# azure-go-secret-property-names

> List of Azure Go SDK property names that are secret

## Install

Install by running:

```sh
go get github.com/dustinspecker/azure-go-secret-property-names
```

## Usage

### GetAll

`GetAll` returns every single secret property name.

```go
import (
 "fmt"

 "github.com/dustinspecker/azure-go-secret-property-names/secretnames"
)

func main() {
 // GetAll returns all Azure Go SDK secret property names
 allSecretNames := secretnames.GetAll()

 fmt.Println(len(allSecretNames) > 0)

 // Output:
 // true
}
```

### GetSubset

While using `GetAll`, you may decide there are a number of them that you don't believe are secret for your use cases.

`GetSubset` allows you to filter out the secret property names you don't believe are secret.

```go
import (
 "fmt"

 "github.com/dustinspecker/azure-go-secret-property-names/secretnames"
)

func main() {
 // You may decide that GetAll returns too many secret property names for your use case.
 // GetSubset allows you to filter out secret property names you don't care about.
 // ignoreList is matched case insensitively
 ignoreList := []string{
   "id",
 }
 subset := secretnames.GetSubset(ignoreList)

 fmt.Println(len(subset) > 0)
 fmt.Println(!slices.Contains(subset, "id"))

 // Output:
 // true
 // true
}
```

## How are secret property names determined?

[generate-secret-property-names.sh](./generate-secret-property-names.sh) and [find-secret-property-names.js](./find-secret-property-names.js) are used to generate the list of secret property names.

This is done by running against a cloned [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs) repository. These scripts look for property names in the specifications that are annotated with `x-ms-secret`.

## Regenerating the list

1. Clone this repository
1. Navigate to the repository directory
1. Clone [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs)
1. Run the following command to generate the list of secret property names:

    ```sh
    ./generate-secret-property-names.sh
    ```

## License

MIT
