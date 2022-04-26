export class AppointmentVm {
    constructor(
        public doctorId: number,
        public dateOfAppointment: Date,
        public patientName: string,
        public patientEmail: string,
        public patientPhone: string,
    ) {}
}