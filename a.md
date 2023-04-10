Given the following go function:

```go
package schema

import (
"bytes"
"context"
"encoding/json"
"fmt"
"os"
"path/filepath"
"strings"
"text/template"

"github.com/ahmetb/go-linq/v3"
"github.com/hashicorp/go-version"
"github.com/hashicorp/hc-install/product"
"github.com/hashicorp/hc-install/releases"
"github.com/hashicorp/terraform-exec/tfexec"
tfjson "github.com/hashicorp/terraform-json"
"github.com/iancoleman/strcase"
)

func RefreshAzureRMSchema() (version *version.Version, err error) {
s, err := ExtractAzureRMProviderSchema()
if err != nil {
return nil, err
}
return s.Version, SaveProviderSchema(s.ProviderSchema)
}
```

I'd like to have a main function which could:

1. Remove folder `generated`
2. run `schema.RefreshAzureRMSchema()`
3. use `go-github` package to check whether `github.com/lonegunmanb/azurerm-provider-schema` repo has a tag with the same version as the one returned by `schema.RefreshAzureRMSchema()`
4. if the tag doesn't exist, execute `git commit -am "update schema to version X.Y.Z"` and `git push -u origin $tag`