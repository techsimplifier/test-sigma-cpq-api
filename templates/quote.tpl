# Quote

Id: {{ .Id }}
QuoteNumber: {{ .QuoteNumber }}
Created: {{ .Created }}
Updated: {{ .Updated }}

## Customer

CustomerRef: {{ .CustomerRef }}
Name: {{ .Name }}

## Items
{{range $item := .Items}}
Id: {{$item.Id}}
Name: {{$item.Name}}
Description: {{$item.Description}}
ProductId: {{$item.ProductId}}

### Meta per item?
{{range $key, $value := $item.MetaTypeLookup}}
   {{$key}}: {{index $value "name"}}
{{end}}

{{end}}