import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { RepositoryService } from 'src/app/core/services/repository/repository.service';
import { AppointmentVm } from '../../models/appointmentVm.model';

@Injectable({
  providedIn: 'root'
})
export class DoctorService extends RepositoryService {

  constructor(http: HttpClient) {
    super(http)
    this.postUrl = 'doctors'
  }

  getDoctorsByAvailability(fromDate: Date, toDate: Date): Observable<any> {
    return this.http.get(
      `${this.url}/${this.postUrl}/fromDate/${fromDate}/toDate/${toDate}`
    )
  }

  postAppointmentToDoctor(doctorId: number, appointmentVm: AppointmentVm): Observable<any> {
    return this.http.post(
      this.url + '/' + this.postUrl + '/' + doctorId + '/make-appointment',
      appointmentVm
    )
  }
}
