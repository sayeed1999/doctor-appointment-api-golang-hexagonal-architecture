import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class TimeService {

  constructor() { }

  getCurrentDateAsMinDateInStringForInputProperty(): string {
    let today = new Date() // new Date() returns current date
    
    let dd = '0' + today.getDate()
    if(dd.length === 3) dd = dd.substring(1)
    
    let mm = '0' + (today.getMonth() + 1) // January is 0!
    if(mm.length === 3) mm = mm.substring(1)

    let yyyy = today.getFullYear()
    
    return yyyy + '-' + mm + '-' + dd 
  }

}
