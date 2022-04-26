import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiResponse } from 'src/app/core/models/api-response.model';
import { TimeService } from 'src/app/core/services/time/time.service';
import swal from 'sweetalert';
import { DoctorService } from '../../services/doctor/doctor.service';

@Component({
  selector: 'app-search-availability',
  templateUrl: './search-availability.component.html',
  styleUrls: ['./search-availability.component.css']
})
export class SearchAvailabilityComponent implements OnInit {

  schedules: any[] = [];
  today: string = ""
  fromDateError: string = ""
  toDateError: string = ""

  searchAvailabilityForm = new FormGroup({
    fromDate: new FormControl(null, Validators.required),
    toDate: new FormControl(null, Validators.required),
  })

  constructor(
    private timeService: TimeService,
    private doctorService: DoctorService
  ) { }

  ngOnInit(): void {
    this.today = this.timeService.getCurrentDateAsMinDateInStringForInputProperty();
  }


  dateChange() {
    // {fromDate: null, toDate: '2022-03-23'}
    const formValue = this.searchAvailabilityForm.getRawValue()
    if(formValue.fromDate !== null && formValue.toDate === null) {
      this.fromDateError = ""
      this.toDateError = "Now choose a to_date"
      this.searchAvailabilityForm.controls['toDate'].markAsTouched()
    }
    else if(formValue.toDate !== null && formValue.fromDate === null) {
      this.fromDateError = "Choose a from_date first"
      this.searchAvailabilityForm.controls['fromDate'].markAsTouched()
      this.searchAvailabilityForm.patchValue({
        toDate: new FormControl(null, Validators.required)
      })
      this.toDateError = ""
    }
    else if(formValue.fromDate > formValue.toDate) {
      this.fromDateError = ""
      this.toDateError = "to_date must not be less than from_date"
      this.searchAvailabilityForm.controls['toDate'].setValue(null)
    }
  }

  formSubmit() {
    this.getDoctorsByAvailability()
  }

  getDoctorsByAvailability() {
    let fromDate: Date = this.searchAvailabilityForm.controls['fromDate'].value
    let toDate: Date = this.searchAvailabilityForm.controls['toDate'].value

    this.doctorService.getDoctorsByAvailability(fromDate, toDate).subscribe(
    (res: ApiResponse) => {
        this.schedules = res.body
      },
      err => {

      }
    )
  }

}
