package sparx_ea

import (
	"database/sql"
)

type TConnector struct {
	ConnectorID        int            `db:"connector_id"`
	Name               sql.NullString `db:"name"`
	Direction          sql.NullString `db:"direction"`
	Notes              sql.NullString `db:"notes"`
	ConnectorType      sql.NullString `db:"connector_type"`
	Subtype            sql.NullString `db:"subtype"`
	Sourcecard         sql.NullString `db:"sourcecard"`
	Sourceaccess       sql.NullString `db:"sourceaccess"`
	Sourceelement      sql.NullString `db:"sourceelement"`
	Destcard           sql.NullString `db:"destcard"`
	Destaccess         sql.NullString `db:"destaccess"`
	Destelement        sql.NullString `db:"destelement"`
	Sourcerole         sql.NullString `db:"sourcerole"`
	Sourceroletype     sql.NullString `db:"sourceroletype"`
	Sourcerolenote     sql.NullString `db:"sourcerolenote"`
	Sourcecontainment  sql.NullString `db:"sourcecontainment"`
	Sourceisaggregate  sql.NullInt64  `db:"sourceisaggregate"`
	Sourceisordered    sql.NullInt64  `db:"sourceisordered"`
	Sourcequalifier    sql.NullString `db:"sourcequalifier"`
	Destrole           sql.NullString `db:"destrole"`
	Destroletype       sql.NullString `db:"destroletype"`
	Destrolenote       sql.NullString `db:"destrolenote"`
	Destcontainment    sql.NullString `db:"destcontainment"`
	Destisaggregate    sql.NullInt64  `db:"destisaggregate"`
	Destisordered      sql.NullInt64  `db:"destisordered"`
	Destqualifier      sql.NullString `db:"destqualifier"`
	StartObjectID      sql.NullInt64  `db:"start_object_id"`
	EndObjectID        sql.NullInt64  `db:"end_object_id"`
	TopStartLabel      sql.NullString `db:"top_start_label"`
	TopMIDLabel        sql.NullString `db:"top_mid_label"`
	TopEndLabel        sql.NullString `db:"top_end_label"`
	BtmStartLabel      sql.NullString `db:"btm_start_label"`
	BtmMIDLabel        sql.NullString `db:"btm_mid_label"`
	BtmEndLabel        sql.NullString `db:"btm_end_label"`
	StartEdge          sql.NullInt64  `db:"start_edge"`
	EndEdge            sql.NullInt64  `db:"end_edge"`
	Ptstartx           sql.NullInt64  `db:"ptstartx"`
	Ptstarty           sql.NullInt64  `db:"ptstarty"`
	Ptendx             sql.NullInt64  `db:"ptendx"`
	Ptendy             sql.NullInt64  `db:"ptendy"`
	Seqno              sql.NullInt64  `db:"seqno"`
	Headstyle          sql.NullInt64  `db:"headstyle"`
	Linestyle          sql.NullInt64  `db:"linestyle"`
	Routestyle         sql.NullInt64  `db:"routestyle"`
	Isbold             sql.NullInt64  `db:"isbold"`
	Linecolor          sql.NullInt64  `db:"linecolor"`
	Stereotype         sql.NullString `db:"stereotype"`
	Virtualinheritance sql.NullString `db:"virtualinheritance"`
	Linkaccess         sql.NullString `db:"linkaccess"`
	Pdata1             sql.NullString `db:"pdata1"`
	Pdata2             sql.NullString `db:"pdata2"`
	Pdata3             sql.NullString `db:"pdata3"`
	Pdata4             sql.NullString `db:"pdata4"`
	Pdata5             sql.NullString `db:"pdata5"`
	DiagramID          sql.NullInt64  `db:"diagramid"`
	EaGuID             sql.NullString `db:"ea_guid"`
	Sourceconstraint   sql.NullString `db:"sourceconstraint"`
	Destconstraint     sql.NullString `db:"destconstraint"`
	Sourceisnavigable  sql.NullInt64  `db:"sourceisnavigable"`
	Destisnavigable    sql.NullInt64  `db:"destisnavigable"`
	Isroot             sql.NullInt64  `db:"isroot"`
	Isleaf             sql.NullInt64  `db:"isleaf"`
	Isspec             sql.NullInt64  `db:"isspec"`
	Sourcechangeable   sql.NullString `db:"sourcechangeable"`
	Destchangeable     sql.NullString `db:"destchangeable"`
	Sourcets           sql.NullString `db:"sourcets"`
	Destts             sql.NullString `db:"destts"`
	Stateflags         sql.NullString `db:"stateflags"`
	Actionflags        sql.NullString `db:"actionflags"`
	Issignal           sql.NullInt64  `db:"issignal"`
	Isstimulus         sql.NullInt64  `db:"isstimulus"`
	Dispatchaction     sql.NullString `db:"dispatchaction"`
	Target2            sql.NullInt64  `db:"target2"`
	Styleex            sql.NullString `db:"styleex"`
	Sourcestereotype   sql.NullString `db:"sourcestereotype"`
	Deststereotype     sql.NullString `db:"deststereotype"`
	Sourcestyle        sql.NullString `db:"sourcestyle"`
	Deststyle          sql.NullString `db:"deststyle"`
	Eventflags         sql.NullString `db:"eventflags"`
}
