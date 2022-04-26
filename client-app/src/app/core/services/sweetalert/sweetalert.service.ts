import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import swal from 'sweetalert'

@Injectable({
  providedIn: 'root'
})
export class SweetalertService {

  constructor(
    private http: HttpClient
  ) { }

  alert(title: string = "Good job!", text: string = "", icon: alertType = alertType.success) {
    swal({
      title: title,
      text: text,
      icon: icon,
    });
  }

}

export enum alertType {
  warning = "warning",
  error = "error",
  success = "success",
  info = "info"
}