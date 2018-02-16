{{- $short := (shortname .Type.Name) -}}
// {{ .Name }} returns the {{ .RefType.Name }} associated with the {{ .Type.Name }}'s {{ .Field.Name }} ({{ .Field.Col.ColumnName }}).
//
// Generated from foreign key '{{ .ForeignKey.ForeignKeyName }}'.
func ({{ $short }} *{{ .Type.Name }}) Get{{ .Name }}s(db XODB) (*{{ .RefType.Name }}, error) {
	service := {{ .Name }}ServiceImpl{
		DB: db,
	}
	return service.{{ .RefType.Name }}By{{ .RefField.Name }}({{ convext $short .Field .RefField }}, {{ convext $short .Field .RefField }})
}
