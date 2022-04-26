import { NgModule } from "@angular/core";
import { RouterModule, Routes } from "@angular/router";
import { AccountLoginComponent } from "./account-login/account-login.component";
import { AccountRegisterComponent } from "./account-register/account-register.component";

const routes: Routes = [
    { path: 'login', component: AccountLoginComponent },
    { path: 'register', component: AccountRegisterComponent },
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