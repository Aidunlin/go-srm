package app

type StudentRecordColumn struct {
	Name                string
	Label               string
	title               string
	BasicSearch         bool
	BasicSearchExact    bool
	AdvancedSearchExact bool
}

func (c StudentRecordColumn) Title() string {
	if len(c.title) > 0 {
		return c.title
	}
	return c.Label
}

func GetStudentColumns() []StudentRecordColumn {
	return []StudentRecordColumn{
		{
			Name:                "id",
			Label:               "ID",
			BasicSearch:         true,
			BasicSearchExact:    true,
			AdvancedSearchExact: true,
		}, {
			Name:        "first_name",
			Label:       "First Name",
			BasicSearch: true,
		}, {
			Name:        "last_name",
			Label:       "Last Name",
			BasicSearch: true,
		}, {
			Name:                "gpa",
			Label:               "GPA",
			AdvancedSearchExact: true,
		}, {
			Name:        "degree_program",
			Label:       "Degree Program",
			BasicSearch: true,
		}, {
			Name:                "graduation_date",
			Label:               "Graduation",
			title:               "Graduation Date",
			AdvancedSearchExact: true,
		}, {
			Name:                "financial_aid",
			Label:               "Aid",
			title:               "Financial Aid",
			AdvancedSearchExact: true,
		}, {
			Name:                "email",
			Label:               "Email",
			BasicSearch:         true,
			BasicSearchExact:    true,
			AdvancedSearchExact: true,
		}, {
			Name:                "phone",
			Label:               "Phone",
			BasicSearch:         true,
			BasicSearchExact:    true,
			AdvancedSearchExact: true,
		},
	}
}

func GetStudentColumn(name string) StudentRecordColumn {
	for _, column := range GetStudentColumns() {
		if column.Name == name {
			return column
		}
	}
	return StudentRecordColumn{}
}
