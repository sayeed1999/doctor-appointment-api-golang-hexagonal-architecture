import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Params } from '@angular/router';
import { ApiResponse } from 'src/app/core/models/api-response.model';
import { DoctorService } from '../../services/doctor/doctor.service';

@Component({
  selector: 'app-doctor',
  templateUrl: './doctor.component.html',
  styleUrls: ['./doctor.component.css']
})
export class DoctorComponent implements OnInit {
  doctor: any = null

  constructor(
    private activatedRoute: ActivatedRoute,
    private doctorService: DoctorService,
  ) { }

  ngOnInit(): void {
    this.getDoctorIdFromRouteParams()
  }

  getDoctorIdFromRouteParams() {
    this.activatedRoute.params.subscribe(
      (params: Params) => {
        const id = params['id']
        this.getDoctorById(id)
      }
    )
  }

  getDoctorById(id: number) {
    this.doctorService.getById(id).subscribe(
      (res: ApiResponse) => {
        this.doctor = res.body
      },
      err => {

      }
    )
  }

}
