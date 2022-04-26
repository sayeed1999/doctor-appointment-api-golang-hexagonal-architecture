import { NgModule } from '@angular/core';
import { SharedModule } from 'src/app/shared/shared.module';
import { DoctorComponent } from './components/doctor/doctor.component';
import { DoctorsListComponent } from './components/doctors-list/doctors.component';
import { LandingPageComponent } from './components/landing-page/landing-page.component';
import { MakeAppointmentComponent } from './components/make-appointment/make-appointment.component';
import { RoutingModule } from './routing.module';
import { SearchAvailabilityComponent } from './components/search-availability/search-availability.component';



@NgModule({
  declarations: [
    LandingPageComponent,
    SearchAvailabilityComponent,
    MakeAppointmentComponent,
    DoctorComponent,
    DoctorsListComponent,
  ],
  imports: [
    RoutingModule,
    SharedModule,
  ],
})
export class HomeModule { }
