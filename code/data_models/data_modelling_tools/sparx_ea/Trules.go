package sparx_ea

import (
	"database/sql"
)

type TRules struct {
	RuleID     string         `db:"ruleid"`
	Rulename   string         `db:"rulename"`
	Ruletype   string         `db:"ruletype"`
	Ruleactive sql.NullInt64  `db:"ruleactive"`
	Errormsg   sql.NullString `db:"errormsg"`
	Flags      sql.NullString `db:"flags"`
	Ruleocl    sql.NullString `db:"ruleocl"`
	RuleXML    sql.NullString `db:"rulexml"`
	Notes      sql.NullString `db:"notes"`
}
