package orm

type DbProps struct {
	Datasource DataSourceProperties
}

func NewDbProps(datasource *DataSourceProperties) *DbProps {
	return &DbProps{
		Datasource: *datasource,
	}
}
