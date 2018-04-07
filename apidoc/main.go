package main

import (
	"os"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/vault/apidoc/apidoc"
	"github.com/hashicorp/vault/builtin/logical/aws"
	"github.com/hashicorp/vault/helper/logging"
	"github.com/hashicorp/vault/vault"
)

func main() {
	c := vault.Core{}
	b := vault.NewSystemBackend(&c, logging.NewVaultLogger(hclog.Trace))
	aws_be := aws.Backend()

	doc := apidoc.NewDoc()
	doc.LoadBackend("sys", b.Backend)
	doc.LoadBackend("aws", aws_be.Backend)

	doc.Add(apidoc.SysPaths...)

	r := apidoc.OAPIRenderer{
		Output:   os.Stdout,
		Template: apidoc.Tmpl,
		Version:  2,
	}
	r.Render(doc)
	_ = r
}
