import { HttpErrorResponse, HttpResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Params, Router } from '@angular/router';
import { ApiResponse } from 'src/app/core/models/api-response.model';
import { alertType, SweetalertService } from 'src/app/core/services/sweetalert/sweetalert.service';
import { TimeService } from 'src/app/core/services/time/time.service';
import { AppointmentVm } from '../../models/appointmentVm.model';
import { DoctorService } from '../../services/doctor/doctor.service';

@Component({
  selector: 'app-make-appointment',
  templateUrl: './make-appointment.component.html',
  styleUrls: ['./make-appointment.component.css']
})
export class MakeAppointmentComponent implements OnInit {

  today: string = "";
  doctor: any;

  appointmentForm = new FormGroup({
    patientName: new FormControl(null, [Validators.required, Validators.minLength(3)]),
    patientEmail: new FormControl(null, [Validators.required, Validators.email]),
    patientPhone: new FormControl(null, [Validators.required, Validators.minLength(5), Validators.maxLength(14)]),
    date: new FormControl(null, Validators.required),
  })

  constructor(
    private router: Router,
    private activatedRoute: ActivatedRoute,
    private doctorService: DoctorService,
    private sweetalert: SweetalertService,
    private timeService: TimeService,
  ) { }

  ngOnInit(): void {
   this.today = this.timeService.getCurrentDateAsMinDateInStringForInputProperty();
   this.activatedRoute.params.subscribe(
     (params: Params) => {
      const ID = params['id']
      this.getDoctorById(ID);
     }
   ) 
  }

  getDoctorById(ID: number) {
    this.doctorService.getById(ID).subscribe(
      (res: ApiResponse) => {
        this.doctor = res.body;
      },
      err => {

      }
    );
  }

  onSubmit() {
    let formValue = this.appointmentForm.getRawValue()
    let appointmentVm = new AppointmentVm(
      this.doctor.ID,
      new Date(formValue.date),
      formValue.patientName,
      formValue.patientEmail,
      formValue.patientPhone
    )

    this.doctorService.postAppointmentToDoctor(this.doctor.ID, appointmentVm).subscribe(
      (res: ApiResponse) => {
        this.router.navigate(['/'])
        this.sweetalert.alert("", res.message, alertType.success)
      },
      (err: HttpErrorResponse) => {
        this.sweetalert.alert("", err.error.message, alertType.info)
      }
    )
  }
}
