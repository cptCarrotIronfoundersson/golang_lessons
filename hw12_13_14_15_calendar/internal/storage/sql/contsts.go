package sqlstorage

const (
	// language=SQL .
	CreateEvent = `insert into events values ($1,$2,$3,$4,$5,$6,$7,$8);`

	// language=SQL .
	UpdateEvent = `
					insert into events values ($1,$2,$3,$4,$5,$6,$7,$8) 
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
	GetEventsByTimeRange = `SELECT * from events 
							where 
								start_datetime::date >= $1::date 
							  	and end_datetime::date <= $2::date;`

	// language=SQL .
	GetAllEvents = `SELECT * from events;`

	// language=SQL .
	DeleteEvent = `delete from events where uuid=$1;`

	DeleteOldEvent = `delete from events where end_datetime < now()-interval '1 year';`
)
