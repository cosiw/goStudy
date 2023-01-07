package main

type Schedule struct {
	ScheduleId   int    `db:"SCHEDULEID`
	OrgId        int    `db:"ORGID"`
	ScheduleDate string `db:"SCHEDULEDATE`
	ScheduleTime string `db:"SCHEDULETIME`
	CustomerId   int    `db:"CUSTOMERID`
	Status       int    `db:SCHEDULESTATUS`
}

type Statement struct {
	StatementId  int     `db:"STATEMENTID"`
	OrgId        int     `db:"ORGID"`
	ScheduleDate string  `db:"SCHEDULEDATE`
	ScheduleId   int     `db:"SCHEDULEID"`
	CustomerId   int     `db:"CUSTOMERID"`
	LiabilityAmt int     `db:"LIABILITYAMT"`
	NoninsAmt    int     `db:"NONINSAMT"`
	TaxableAmt   float64 `db:"TAXABLEAMT"`
}

type Payment struct {
	OrgId           int    `db:"ORGID"`
	CustomerId      int    `db:"CUSTOMERID"`
	ScheduleId      int    `db:"SCHEDULEID"`
	PaymentAmt      int    `db:"PAYMENTAMT"`
	PayLiabilityAmt int    `db: "PAYLIABILITYAMT"`
	PaymentCode     string `db:"PAYMENTCODE"`
	PayDate         string `db:"PAYDATE"`
	Tax             int    `db:"TAX"`
	CrTime          string `db:"CRTIME`
}
