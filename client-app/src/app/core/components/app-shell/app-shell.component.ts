import { Component, OnInit } from '@angular/core';
import { DoctorService } from 'src/app/feature/home/services/doctor/doctor.service';
import { ApiResponse } from '../../models/api-response.model';

@Component({
  selector: 'app-shell',
  templateUrl: './app-shell.component.html',
  styleUrls: ['./app-shell.component.css']
})
export class AppShellComponent implements OnInit {

  doctors: any[] = []

  constructor(private doctorService: DoctorService) { }

  ngOnInit(): void {
    this.getDoctors()
  }

  getDoctors() {
    this.doctorService.getAll().subscribe(
      (res: ApiResponse) => {
        this.doctors = res.body
      },
      err => {

      }
    );
  }

}
