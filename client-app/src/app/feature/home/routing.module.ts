import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { DoctorComponent } from "./components/doctor/doctor.component";
import { DoctorsListComponent } from "./components/doctors-list/doctors.component";
import { LandingPageComponent } from "./components/landing-page/landing-page.component";
import { MakeAppointmentComponent } from "./components/make-appointment/make-appointment.component";
import { SearchAvailabilityComponent } from "./components/search-availability/search-availability.component";

const routes: Routes = [
    { path: 'doctors/:id', component: DoctorComponent },
    { path: 'doctors-list', component: DoctorsListComponent },
    { path: 'doctors/:id/appointment', component: MakeAppointmentComponent },
    { path: 'search-availability', component: SearchAvailabilityComponent },
    { path: '', component: LandingPageComponent },
] 

@NgModule({
    imports: [
        RouterModule.forChild(routes)
    ],
    exports: [
        RouterModule
    ]
})
export class RoutingModule{}