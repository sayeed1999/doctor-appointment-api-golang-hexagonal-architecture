use [doctor-appointment-db]
go

create procedure CreateAppointmentToDoctor (
	@doctorId int
	,@dateOfAppointment datetime
	,@patientName nvarchar(max)
	,@patientEmail nvarchar(max)
	,@patientPhone nvarchar(max)
) as
begin
	set nocount on
	declare @appointmentId int
	declare @capacity int = 3 -- capacity a doctor can hold

	IF NOT EXISTS( select * from doctors where @doctorId = id )
		throw 90001, 'no doctor found with the given primary key ID of doctor', 1
	
	declare @appointmentsCount int
	set @appointmentsCount = (
		select
			count(id) 
		from appointments 
		where @doctorId = doctor_id 
			and @dateOfAppointment = date_of_appointment
		group by
			doctor_id,
			date_of_appointment
	)

	IF @appointmentsCount >= @capacity
		throw 90001, 'doctor is full on this date. please make appointment on another date. also check Search Availability', 2

	BEGIN TRANSACTION

		insert into appointments
		(doctor_id, date_of_appointment)
		values
		(@doctorId, @dateOfAppointment)

		-- set error returned from inserting appointments table
		declare @error int = @@error

		set @appointmentId = SCOPE_IDENTITY()
	
		insert into appointment_details
		(appointment_id, patient_name, patient_email, patient_phone)
		values
		(@appointmentId, @patientName, @patientEmail, @patientPhone)
		
		-- set error returned from inserting appointment_details table
		set @error = IIF(@@error <> 0, @@error, @error)

		if @error <> 0
		begin
			ROLLBACK
			print 'Transaction rolled back'
		end
		else
		begin
			COMMIT
			print 'Transaction committed'
			select 'Doctor appointment on your preferred date has been processed successfully! Thank you. :)'
		end

end
go

-- exec CreateAppointmentToDoctor 3, '2022-12-22', 'asd', 'asd', 'asd'
-- exec CreateAppointmentToDoctor 2, '2022-03-15', 'Sayeed Rahman', 'hahaha@ruet.com', '+88017654321'
-- select * from dbo.appointments
-- select * from dbo.appointment_details