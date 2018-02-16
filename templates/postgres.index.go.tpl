{{- $short := (shortname .Type.Name "err" "sqlstr" "db" "q" "res" "XOLog" .Fields) -}}
{{- $table := (schema .Schema .Type.Table.TableName) -}}


// {{ .FuncName }} retrieves a row from '{{ $table }}' as a {{ .Type.Name }}.
//
// Generated from index '{{ .Index.IndexName }}'.
func (serviceImpl *{{ .Type.Name }}ServiceImpl){{ .FuncName }}(_{{ goparamlist .Fields true true }}) ({{ if not .Index.IsUnique }}[]{{ end }}*{{ .Type.Name }}, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`{{ colnames .Type.Fields }} ` +
		`FROM {{ $table }} ` +
		`WHERE {{ colnamesquery .Fields " AND " }}`

	// run query
	XOLog(sqlstr{{ goparamlist .Fields true false }})
{{- if .Index.IsUnique }}

	{{ $short }} := {{ .Type.Name }}{
	{{- if .Type.PrimaryKey }}

	{{ end -}}
}

	err = serviceImpl.DB.QueryRow(sqlstr{{ goparamlist .Fields true false }}).Scan({{ fieldnames .Type.Fields (print "&" $short) }})
	if err != nil {
		return nil, err
	}

	return &{{ $short }}, nil
{{- else }}
	q, err := db.Query(sqlstr{{ goparamlist .Fields true false }})
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*{{ .Type.Name }}{}
	for q.Next() {
		{{ $short }} := {{ .Type.Name }}{
		{{- if .Type.PrimaryKey }}
			
		{{ end -}}
		}

		// scan
		err = q.Scan({{ fieldnames .Type.Fields (print "&" $short) }})
		if err != nil {
			return nil, err
		}

		res = append(res, &{{ $short }})
	}

	return res, nil
{{- end }}
}
