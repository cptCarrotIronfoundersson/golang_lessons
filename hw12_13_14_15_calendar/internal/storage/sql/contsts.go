package sqlstorage

const (
	// language=SQL .
	CreateEvent = `insert into events values (?,?,?,?,?,?,?,?);`

	// language=SQL .
	UpdateEvent = `
					insert into events values (?,?,?,?,?,?,?,?) 
					on conflict on constraint events_pkey
					do update set 
					title=excluded.title,
					datetime=excluded.datetime,
					start_datetime=excluded.start_datetime,
					end_datetime=excluded.end_datetime,
					description=excluded.description,
					userid=excluded.userid,
					remind_time_before=excluded.remind_time_before;
`
	// language=SQL .
	GetEventsByTimeRange = `SELECT * from events where start_datetime > ? and end_datetime < ?;`

	// language=SQL .
	GetAllEvents = `SELECT * from events;`

	// language=SQL .
	DeleteEvent = `delete from events where uuid=?;`
)
