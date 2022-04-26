use [doctor-appointment-db]
go


CREATE PROCEDURE SearchAvailability (
	@fromDate datetime
	,@toDate datetime
) AS
BEGIN
	set nocount on

	declare @capacity int = 3 -- the capacity a doctor can hold for each day
	
	declare @doctor_appointments_count as table (
		doctor_id int
		,date_of_appointment datetime
		,no_of_appointments int
	)

	insert into @doctor_appointments_count
	select
		doctor_id
		,date_of_appointment
		,count(id)
	from appointments
	where date_of_appointment between @fromDate and @toDate
	group by 
		date_of_appointment
		,doctor_id


	declare @calender as table (
		date datetime
	)

	while @fromDate <= @toDate
	begin
		insert into @calender
		select @fromDate	
		set @fromDate = DATEADD(day, 1, @fromDate)
	end


	select
		c.date
		,d.id doctor_id
		,d.name doctor_name
		,CAST(ISNULL(dac.no_of_appointments, 0) as nvarchar(max)) + '/' + CAST(3 as nvarchar(max)) appointments_to_capacity_ratio
		,IIF(ISNULL(dac.no_of_appointments, 0) < @capacity, 1, 0) is_available
	from doctors d
	cross join @calender c
	left join @doctor_appointments_count dac
		on dac.doctor_id = d.id
		and dac.date_of_appointment = c.date
	order by
		c.date
		,d.name
END

-- exec SearchAvailability '2022-03-12', '2022-03-13'
